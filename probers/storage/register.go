package storage

import (
	"os/exec"

	"github.com/Cloud-Foundations/tricorder/go/tricorder"
	"github.com/Cloud-Foundations/tricorder/go/tricorder/units"
)

func register(dir *tricorder.DirectorySpec) *prober {
	p := new(prober)
	p.dir = dir
	if megaCliPath, err := exec.LookPath("megacli"); err == nil {
		err := dir.RegisterMetric("health", &p.health, units.None,
			"health of all storage devices")
		if err == nil {
			go p.loopMegaCliProbe(megaCliPath)
		}
	}
	p.storageDevices = make(map[string]*storageDevice)
	return p
}
