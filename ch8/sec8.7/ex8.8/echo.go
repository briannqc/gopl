package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer func() {
		_ = c.Close()
	}()
	_, _ = fmt.Fprintln(c, "Connected to", c.LocalAddr())

	chText := make(chan string, 10)
	go scanText(chText, c)

	timeout := 10 * time.Second
	ticker := time.NewTicker(timeout)
	for {
		select {
		case <-ticker.C:
			_, _ = fmt.Fprintf(c, "I got no message from you for %v, are you alive?\n", timeout)
			return
		case text := <-chText:
			ticker.Reset(timeout)
			echo(c, text, time.Second)
		}
	}
}

func scanText(c chan<- string, r io.Reader) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		c <- input.Text()
	}
	close(c)
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept a conn failed", err)
			continue
		}
		go handleConn(conn)
	}
}
