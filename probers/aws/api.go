package aws

import (
	"github.com/Cloud-Foundations/health-agent/lib/proberlist"
)

func New() proberlist.RegisterProber {
	return _new()
}
