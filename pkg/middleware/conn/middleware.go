package conn

import (
	"errors"
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
		n, err = next(b)
		if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
			return n, err
		}
		if err != nil {
			log.Printf("write: %s\n", err)
			return n, err
		}

		log.Printf("write %d bytes: %s\n", n, b[:n])
		return n, nil
	}
}

func WithReadLogging(conn net.Conn, next Handler) Handler {
	return func(b []byte) (n int, err error) {
		n, err = next(b)
		if errors.Is(err, io.EOF) || errors.Is(err, net.ErrClosed) {
			return n, err
		}
		if err != nil {
			log.Printf("read: %s\n", err)
			return n, err
		}

		log.Printf("read %d bytes: %s\n", n, b[:n])
		return n, nil
	}
}
