package memory

import (
	"github.com/Cloud-Foundations/Dominator/lib/meminfo"
)

func (p *prober) probe() error {
	if info, err := meminfo.GetMemInfo(); err != nil {
		return err
	} else {
		p.available = info.Available
		p.free = info.Free
		p.haveAvailable = info.HaveAvailable
		p.total = info.Total
	}
	return nil
}
