package netif

import (
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

func register(dir *tricorder.DirectorySpec) *prober {
	p := new(prober)
	p.dir = dir
	p.netInterfaces = make(map[string]*netInterface)
	return p
}
