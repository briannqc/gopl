package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:2121")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	_, _ = fmt.Fprintf(conn, "Pwd: %s\n", wd)

	space := regexp.MustCompile(`\s+`)

	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		line := space.ReplaceAllString(buf.Text(), " ")
		if line == "" {
			continue
		}

		cmd := strings.Split(line, " ")
		ops := cmd[0]
		var args []string
		if len(cmd) > 1 {
			args = cmd[1:]
		}

		if ops == "cd" {
			to, err := cd(wd, args[0])
			if err != nil {
				_, _ = fmt.Fprintf(conn, "cd failed, err: %v\n", err)
			} else {
				wd = to
				_, _ = fmt.Fprintf(conn, "Changed to: %s\n", wd)
			}
		} else if ops == "ls" {
			if err := ls(conn, wd); err != nil {
				_, _ = fmt.Fprintf(conn, "ls failed, err: %v\n", err)
			}
		} else if ops == "get" {
			file := args[0]
			if err := get(conn, wd, file); err != nil {
				_, _ = fmt.Fprintf(conn, "get file content failed, file: %s err: %v\n", file, err)
			}
		} else if ops == "close" {
			_, _ = fmt.Fprint(conn, "Bye bye!\n")
			return
		} else {
			_, _ = fmt.Fprintf(conn, "%s: unsupported ops\n", ops)
		}
	}
}

func cd(wd string, to string) (string, error) {
	to = filepath.Join(wd, to)
	dir, err := os.Open(to)
	if err != nil {
		return "", err
	}
	stat, err := dir.Stat()
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return "", fmt.Errorf("%s: is not a dir", to)
	}
	return to, nil
}

func ls(w io.Writer, wd string) error {
	fileInfo, err := ioutil.ReadDir(wd)
	if err != nil {
		return err
	}
	for _, fi := range fileInfo {
		_, _ = fmt.Fprintf(w, "%s\n", fi.Name())
	}
	return nil
}

func get(w io.Writer, wd, filename string) error {
	filename = filepath.Join(wd, filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = io.Copy(w, file)
	return err
}
