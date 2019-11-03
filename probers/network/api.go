package network

import (
	libprober "github.com/Cloud-Foundations/health-agent/lib/prober"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type prober struct {
	gatewayAddress              string
	gatewayInterfaceName        string
	gatewayPingTimeDistribution *tricorder.CumulativeDistribution
	gatewayRttDistribution      *tricorder.CumulativeDistribution
	resolverDomain              string
	resolverNameservers         *tricorder.List
	resolverSearchDomains       *tricorder.List
}

func Register(dir *tricorder.DirectorySpec) libprober.Prober {
	return register(dir)
}

func (p *prober) Probe() error {
	return p.probe()
}
