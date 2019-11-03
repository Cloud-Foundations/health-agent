package scheduler

import (
	"io"
	"time"

	libprober "github.com/Cloud-Foundations/health-agent/lib/prober"
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type loadAverage struct {
	oneMinute      float32
	fiveMinutes    float32
	fifteenMinutes float32
}

type timeMetric struct {
	value    time.Duration
	fraction float64
}

type cpuStatistics struct {
	lastProbeTime time.Time
	idleTime      timeMetric
	iOWaitTime    timeMetric
	irqTime       timeMetric
	userNiceTime  timeMetric
	softIrqTime   timeMetric
	systemTime    timeMetric
	userTime      timeMetric
}

type prober struct {
	loadavg      loadAverage
	numCpus      uint64
	cpuStats     cpuStatistics
	lastCpuStats cpuStatistics
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
