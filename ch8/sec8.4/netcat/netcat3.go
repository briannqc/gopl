package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	in := strings.NewReader("Hello!!!")
	mustCopy(conn, in)

	if conn, ok := conn.(*net.TCPConn); ok {
		if err := conn.CloseWrite(); err != nil {
			log.Fatal(err)
		}
	}
	<-done
}

func mustCopy(w io.Writer, r io.Reader) {
	_, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
}
