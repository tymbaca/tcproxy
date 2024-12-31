package strategy

type Strategy interface {
	GetTarget() string
	// maybe second method for leastConn
}
