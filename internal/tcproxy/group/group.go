package group

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/tymbaca/tcproxy/internal/config"
	"github.com/tymbaca/tcproxy/internal/strategy"
	"golang.org/x/sync/errgroup"
)

func New(cfg config.Group) (*Group, error) {
	strategy, err := newStrategy(cfg)
	if err != nil {
		return nil, fmt.Errorf("can't init strategy: %w", err)
	}

	return &Group{
		cfg:      cfg,
		strategy: strategy,
	}, nil
}

type Group struct {
	cfg      config.Group
	strategy strategy.Strategy
}

func (g *Group) Run(ctx context.Context) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", g.cfg.Port))
	if err != nil {
		return fmt.Errorf("can't listen: %w", err)
	}

	go g.listen(ctx, l)

	return nil
}

func (g *Group) listen(ctx context.Context, l net.Listener) {
	log.Printf("group listening on port %d...\n", g.cfg.Port)
	defer l.Close()

	for {
		if err := ctx.Err(); err != nil {
			log.Println("context canceled, closing group")
			return
		}

		conn, err := l.Accept()
		if err != nil {
			log.Printf("can't accept connection on port %d: %s", g.cfg.Port, err)
			continue
		}

		go g.handleConn(ctx, conn)
	}
}

func (g *Group) handleConn(ctx context.Context, clientConn net.Conn) {
	defer clientConn.Close()

	serverAddr := g.strategy.GetTarget()

	serverConn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Printf("can't dial the target: %s", err)
		return
	}
	defer serverConn.Close()

	log.Printf("established connection (client: %v, server: %v)", clientConn.RemoteAddr(), serverConn.RemoteAddr())

	var wg errgroup.Group
	wg.Go(func() error {
		// TODO: wg context mutual cancelation
		_, err := io.Copy(serverConn, clientConn)
		return err
	})
	wg.Go(func() error {
		_, err := io.Copy(clientConn, serverConn)
		return err
	})

	if err := wg.Wait(); err != nil {
		log.Printf("can't tranfer data: %s", err)
		return
	}
}

func newStrategy(cfg config.Group) (strategy.Strategy, error) {
	switch cfg.Strategy {
	case config.RandomStrategy:
		return strategy.NewRandom(cfg.Targets), nil
	case config.RoundRobinStrategy:
		return strategy.NewRoundRobin(cfg.Targets), nil
	}

	return nil, fmt.Errorf("unsupported strategy: %s", cfg.Strategy)
}
