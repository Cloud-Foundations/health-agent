package main

import (
	"github.com/Cloud-Foundations/health-agent/lib/proberlist"
	"github.com/Cloud-Foundations/health-agent/probers/aws"
	"github.com/Cloud-Foundations/health-agent/probers/dmi"
	"github.com/Cloud-Foundations/health-agent/probers/filesystems"
	"github.com/Cloud-Foundations/health-agent/probers/kernel"
	"github.com/Cloud-Foundations/health-agent/probers/memory"
	"github.com/Cloud-Foundations/health-agent/probers/netif"
	"github.com/Cloud-Foundations/health-agent/probers/network"
	"github.com/Cloud-Foundations/health-agent/probers/packages"
	"github.com/Cloud-Foundations/health-agent/probers/scheduler"
	"github.com/Cloud-Foundations/health-agent/probers/storage"
	"github.com/Cloud-Foundations/health-agent/probers/systime"
	"github.com/Cloud-Foundations/health-agent/probers/virsh"
)

func setupProbers() (*proberlist.ProberList, error) {
	pl := proberlist.New("/probers")
	pl.CreateAndAdd(filesystems.Register, "/sys/fs", 0)
	pl.CreateAndAdd(scheduler.Register, "/sys/sched", 0)
	pl.CreateAndAdd(memory.Register, "/sys/memory", 0)
	pl.CreateAndAdd(netif.Register, "/sys/netif", 0)
	pl.CreateAndAdd(dmi.Register, "/sys/dmi", 0)
	pl.CreateAndAdd(network.Register, "/sys/network", 0)
	pl.CreateAndAdd(storage.Register, "/sys/storage", 0)
	pl.CreateAndAdd(systime.Register, "/sys/systime", 0)
	pl.CreateAndAdd(kernel.Register, "/sys/kernel", 0)
	pl.CreateAndAdd(packages.Register, "/sys/packages", 0)
	pl.Add(virsh.New(), "/sys/hypervisor/virsh", 0)
	go func() { pl.Add(aws.New(), "/sys/cloud/aws", 0) }()
	return pl, nil
}
