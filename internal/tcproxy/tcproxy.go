package tcproxy

import (
	"context"
	"fmt"
	"log"

	"github.com/tymbaca/tcproxy/internal/config"
	"github.com/tymbaca/tcproxy/internal/tcproxy/group"
)

type TCProxy struct {
	Groups []*group.Group
}

func New(cfg config.Config) (*TCProxy, error) {
	var groups []*group.Group
	for _, groupCfg := range cfg.Groups {
		group, err := group.New(groupCfg)
		if err != nil {
			return nil, fmt.Errorf("can't create group (for port %d): %w", groupCfg.Port, err)
		}

		groups = append(groups, group)
	}

	return &TCProxy{
		Groups: groups,
	}, nil
}

func (p *TCProxy) Run(ctx context.Context) error {
	log.Printf("staring tcproxy...\n")
	for _, g := range p.Groups {
		err := g.Run(ctx)
		if err != nil {
			return fmt.Errorf("can't run group: %w", err)
		}
	}

	log.Printf("tcproxy started!\n")

	<-ctx.Done()
	return ctx.Err()
}
