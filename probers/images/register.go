package images

import (
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
	"github.com/Cloud-Foundations/tricorder/go/tricorder/units"
)

func register(dir *tricorder.DirectorySpec) *prober {
	p := &prober{}
	if err := dir.RegisterMetric("initial-image", &p.initialImage, units.None,
		"image the system was created with"); err != nil {
		panic(err)
	}
	if err := dir.RegisterMetric("patched-image", &p.patchedImage, units.None,
		"image the system was last patched with"); err != nil {
		panic(err)
	}
	return p
}
