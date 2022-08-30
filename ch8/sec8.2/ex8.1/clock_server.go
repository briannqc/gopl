package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 8000, "Listening port, default is 8000")
	flag.Parse()

	loc, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Println("Connected to", conn.RemoteAddr())
		go handleConn(conn, loc)
	}
}

func handleConn(conn net.Conn, loc *time.Location) {
	defer func() {
		_ = conn.Close()
	}()

	for {
		_, err := io.WriteString(conn, time.Now().In(loc).Format(time.RFC850+"\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
