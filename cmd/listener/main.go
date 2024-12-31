package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"slices"
	"strings"
)

var portFlag = flag.String("p", "8080", "port to listen to")

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", ":"+*portFlag)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Printf("listening port %s...", *portFlag)
	for {
		conn, err := l.Accept()
		log.Printf("accepted connection (%v)\n", conn.RemoteAddr())
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Printf("closed connection (%v)\n", conn.RemoteAddr())
	}()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			return
		}

		msg := buf[:n:n]

		fmt.Printf("got: %s", msg)

		slices.Reverse(msg)

		_, err = io.Copy(conn, strings.NewReader(string(msg)+"\n"))
		if err != nil {
			panic(err)
		}

	}

	// warn: promlem here
	// data, err := io.ReadAll(conn)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(string(data))
	//
	// slices.Reverse(data)
	//
	// _, err = conn.Write(data)
	// if err != nil {
	// 	panic(err)
	// }
}
