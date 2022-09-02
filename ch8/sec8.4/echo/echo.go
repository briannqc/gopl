package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		if conn, ok := conn.(*net.TCPConn); ok {
			go handleConn(conn)
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		text := input.Text()

		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, text, 1*time.Second)
		}()
	}
	_ = c.CloseRead()

	wg.Wait()
	_ = c.CloseWrite()
}
