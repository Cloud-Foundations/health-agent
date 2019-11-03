package netif

import (
	"io"
	"time"

	libprober "github.com/Cloud-Foundations/health-agent/lib/prober"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type prober struct {
	dir           *tricorder.DirectorySpec
	netInterfaces map[string]*netInterface // map[name]*netInterface
}

type netInterface struct {
	dir                 *tricorder.DirectorySpec
	name                string
	carrier             bool
	multicastFrames     uint64
	rxCompressedPackets uint64
	rxData              wrappingInt
	rxDropped           uint64
	rxErrors            uint64
	rxFrameErrors       uint64
	rxOverruns          uint64
	rxPackets           wrappingInt
	speed               uint64
	txCarrierLosses     uint64
	txCollisionErrors   uint64
	txCompressedPackets uint64
	txData              wrappingInt
	txDropped           uint64
	txErrors            uint64
	txOverruns          uint64
	txPackets           wrappingInt
	probed              bool
	lastProbeTime       time.Time
	lastRxData          uint64
	rxDataRate          uint64
	lastTxData          uint64
	txDataRate          uint64
}

type wrappingInt struct {
	value  uint64
	offset uint64
	tmp    uint64
}

func Register(dir *tricorder.DirectorySpec) libprober.Prober {
	return register(dir)
}

func (p *prober) Probe() error {
	return p.probe()
}

func (p *prober) WriteHtml(writer io.Writer) {
	p.writeHtml(writer)
}
