package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			// outgoing message channels.
			for cli := range clients {
				cli <- msg
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
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- namedClient{
		name:   who,
		client: ch,
	}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	_ = conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg)
	}
}
