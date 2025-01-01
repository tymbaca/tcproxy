package conn

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type Middleware func(conn net.Conn, next Handler) Handler

func wrapHandler(conn net.Conn, handler Handler, middlewares ...Middleware) Handler {
	for _, mw := range middlewares {
		handler = mw(conn, handler)
	}

	return handler
}

func WithWriteLogging(conn net.Conn, next Handler) Handler {
	return func(b []byte) (n int, err error) {
		return commonLogging("write", conn, next, b)
	}
}

func WithReadLogging(conn net.Conn, next Handler) Handler {
	return func(b []byte) (n int, err error) {
		return commonLogging("read", conn, next, b)
	}
}

func commonLogging(op string, conn net.Conn, next Handler, b []byte) (n int, err error) {
	prefix := fmt.Sprintf("[conn: %v] %s", conn.RemoteAddr(), op)

	n, err = next(b)
	if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
		return n, err
	}
	if err != nil {
		log.Printf("%s: %s\n", prefix, err)
		return n, err
	}

	log.Printf("%s %d bytes: %s\n", prefix, n, b[:n])

	return n, nil
}
