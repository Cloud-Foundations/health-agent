package packages

import (
	"time"

	libprober "github.com/Cloud-Foundations/health-agent/lib/prober"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type packageEntry struct {
	dir     *tricorder.DirectorySpec
	name    string
	version string
	size    uint64
}

type packageList struct {
	dir                      *tricorder.DirectorySpec
	packages                 map[string]*packageEntry
	packagesAddedDuringProbe map[string]struct{}
}

type prober struct {
	dir                *tricorder.DirectorySpec
	lastProbeStartTime time.Time
	packagers          map[string]*packageList
}

func Register(dir *tricorder.DirectorySpec) libprober.Prober {
	return register(dir)
}

func (p *prober) Probe() error {
	return p.probe()
}
