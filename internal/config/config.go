package config

type Config struct {
	Groups []Group
}

type Group struct {
	Port     int
	Strategy Strategy
	Targets  []string
}

type Strategy string

const (
	RandomStrategy     Strategy = "random"
	RoundRobinStrategy Strategy = "roundRobin"
	LeastConnStrategy  Strategy = "leastConn"
)

func Parse(path string) (Config, error) {
	// TODO
	return Config{
		Groups: []Group{
			{
				Port:     8080,
				Strategy: RoundRobinStrategy,
				Targets: []string{
					"localhost:8090",
					"localhost:8091",
					"localhost:8092",
				},
			},
		},
	}, nil
}
