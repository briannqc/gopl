package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		pair := strings.Split(arg, "=")
		_, addr := pair[0], pair[1]

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Print(err)
			continue
		}
		go printClock(conn)
	}

	forever := make(chan bool)
	<-forever
}

func printClock(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
