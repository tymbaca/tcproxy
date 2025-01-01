package config

import "github.com/hashicorp/hcl/v2/hclsimple"

type Config struct {
	Groups []Group `hcl:"group,block"`
}

type Group struct {
	Port     int      `hcl:"port"`
	Protocol Protocol `hcl:"protocol"`
	Strategy Strategy `hcl:"strategy"`
	Targets  []Target `hcl:"target,block"`
}

type Strategy = string

const (
	StrategyRandom     Strategy = "random"
	StrategyRoundRobin Strategy = "round_robin"
	StrategyLeastConn  Strategy = "least_conn"
)

type Protocol = string

const (
	ProtocolTCP Protocol = "tcp"
	ProtocolUDP Protocol = "udp"
)

type Target struct {
	Addr string `hcl:"addr"`
}

func Parse(path string) (Config, error) {
	var config Config
	err := hclsimple.DecodeFile(path, nil, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
