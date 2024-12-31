package strategy

import "sync"

type RoundRobin struct {
	mu        sync.Mutex
	nextIndex int
	targets   []string
}

func NewRoundRobin(targets []string) *RoundRobin {
	return &RoundRobin{
		targets: targets,
	}
}

func (s *RoundRobin) GetTarget() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	target := s.targets[s.nextIndex]

	// move next index
	s.nextIndex++
	if s.nextIndex >= len(s.targets) {
		s.nextIndex = 0
	}

	return target
}
