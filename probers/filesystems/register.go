package filesystems

import (
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

func register(dir *tricorder.DirectorySpec) *prober {
	p := new(prober)
	p.dir = dir
	p.fileSystems = make(map[string]*fileSystem)
	return p
}
