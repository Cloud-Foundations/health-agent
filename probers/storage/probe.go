package storage

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Cloud-Foundations/tricorder/go/tricorder/units"
)

var filename string = "/proc/partitions"

func (p *prober) probe() error {
	for _, device := range p.storageDevices {
		device.probed = false
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := p.processPartitionLine(scanner.Text()); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	for name, device := range p.storageDevices {
		if !device.probed {
			delete(p.storageDevices, name)
			device.dir.UnregisterDirectory()
		}
	}
	return nil
}

func (p *prober) processPartitionLine(line string) error {
	var major, minor, size uint64
	var name string
	nScanned, err := fmt.Sscanf(line, " %d %d %d %s",
		&major, &minor, &size, &name)
	if err != nil {
		return nil
	}
	if nScanned < 4 {
		return nil
	}
	if name[len(name)-1] >= '0' && name[len(name)-1] <= '9' {
		return nil
	}
	device := p.storageDevices[name]
	if device == nil {
		device = new(storageDevice)
		p.storageDevices[name] = device
		metricsDir, err := p.dir.RegisterDirectory(name)
		if err != nil {
			return err
		}
		if err := metricsDir.RegisterMetric("size", &device.size,
			units.Byte, "size of storage device"); err != nil {
			return err
		}
	}
	device.size = size << 10
	device.probed = true
	return nil
}
