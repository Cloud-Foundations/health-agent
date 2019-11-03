package virsh

import (
	"github.com/Cloud-Foundations/health-agent/lib/proberlist"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

func newProber() proberlist.RegisterProber {
	p := &prober{domains: make(map[string]*domainInfo)}
	if p.listDomains() != nil {
		return nil
	}
	return p
}

func (p *prober) register(dir *tricorder.DirectorySpec) error {
	domainsDir, err := dir.RegisterDirectory("domains")
	if err != nil {
		return err
	}
	p.domainsDir = domainsDir
	return nil
}
