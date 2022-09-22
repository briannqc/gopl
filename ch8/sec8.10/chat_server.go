package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

type namedClient struct {
	name   string
	client client
}

var (
	entering = make(chan namedClient)
	leaving  = make(chan client)
	messages = make(chan string, 100)
)

func broadcaster() {
	clients := make(map[client]string)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients'
			// outgoing message channels. Wait at most 1sec,
			// or skip if client is unable to read the message.
			for cli := range clients {
				timer := time.NewTimer(1 * time.Second)
				select {
				case cli <- msg:
					timer.Stop()
					continue
				case <-timer.C:
					continue
				}
			}
		case cli := <-entering:
			clients[cli.client] = cli.name

			var participants []string
			for _, who := range clients {
				participants = append(participants, who)
			}
			messages <- "Participants: " + strings.Join(participants, ", ")
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 1000)
	_, _ = fmt.Fprintln(conn, "What is your name?")

	input := bufio.NewScanner(conn)
	var who = conn.RemoteAddr().String()
	if input.Scan() {
		who = input.Text()
	}

	go clientWriter(conn, ch)

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- namedClient{
		name:   who,
		client: ch,
	}

	done := make(chan bool)
	maxIdle := 5 * 60 * time.Second
	ticker := time.NewTicker(maxIdle)

	go func() {
		select {
		case <-done:
			leaving <- ch
			messages <- who + " has left"
			_ = conn.Close()
		case <-ticker.C:
			leaving <- ch
			messages <- who + " stayed idle -> disconnect it"
			_ = conn.Close()
		}
	}()

	for input.Scan() {
		ticker.Reset(maxIdle)
		messages <- who + ": " + input.Text()
	}
	close(done)
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg)
	}
}
