package strategy

import "math/rand"

type Random struct {
	targets []string
}

func NewRandom(targets []string) *Random {
	return &Random{
		targets: targets,
	}
}

func (s *Random) GetTarget() string {
	return s.targets[rand.Intn(len(s.targets))]
}
