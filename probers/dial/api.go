package dial

import (
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type dialConfig struct {
	network string
	address string
	healthy bool
}

func MakeDialProber(network, address string) *dialConfig {
	return &dialConfig{
		address: address,
		network: network,
	}
}

func (p *dialConfig) HealthCheck() bool {
	return p.healthy
}

func (p *dialConfig) Probe() error {
	return p.probe()
}

func (p *dialConfig) Register(dir *tricorder.DirectorySpec) error {
	return p.register(dir)
}
