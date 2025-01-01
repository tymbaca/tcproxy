package group

import (
	"net"

	conn_mw "github.com/tymbaca/tcproxy/pkg/middleware/conn"
)

type dialer struct {
	readMws  []conn_mw.Middleware
	writeMws []conn_mw.Middleware
}

func newDialer(readMiddlewares, writeMiddlewares []conn_mw.Middleware) *dialer {
	return &dialer{
		readMws:  readMiddlewares,
		writeMws: writeMiddlewares,
	}
}

func (b *dialer) Dial(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return conn_mw.Wrap(conn, b.readMws, b.writeMws), nil
}
