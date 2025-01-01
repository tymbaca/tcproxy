package conn

import "net"

type Handler func(b []byte) (n int, err error)

type Conn struct {
	net.Conn
	readHandler  Handler
	writeHandler Handler
}

func Wrap(conn net.Conn, readMiddlewares, writeMiddlewares []Middleware) *Conn {
	readHandler := wrapHandler(conn, conn.Read, readMiddlewares...)
	writeHandler := wrapHandler(conn, conn.Write, writeMiddlewares...)

	return &Conn{
		Conn:         conn,
		readHandler:  readHandler,
		writeHandler: writeHandler,
	}
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.readHandler(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return c.writeHandler(b)
}
