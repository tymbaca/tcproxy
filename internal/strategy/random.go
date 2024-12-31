package strategy

import "math/rand"

func NewRandom(targets []string) *Random {
	return &Random{
		targets: targets,
	}
}

type Random struct {
	targets []string
}

func (s *Random) GetTarget() string {
	return s.targets[rand.Intn(len(s.targets))]
}
