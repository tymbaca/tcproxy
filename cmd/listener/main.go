package main

import (
	"flag"
	"io"
	"log"
	"net"
	"slices"
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

	data, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	slices.Reverse(data)

	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}
}