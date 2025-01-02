package conn

import "net"

func WithByteCounter(addFn func(delta int)) Middleware {
	return func(conn net.Conn, next Handler) Handler {
		return func(b []byte) (n int, err error) {
			n, err = next(b)

			if n != 0 {
				addFn(n)
			}

			return n, err
		}
	}
}
