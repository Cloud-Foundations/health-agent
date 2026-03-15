package testprog

import (
	"github.com/Cloud-Foundations/tricorder/go/tricorder"
)

type testprogconfig struct {
	testname  string
	command   string
	args      []string
	healthy   bool
	exitCode  int
	exitError string
	stdout    string
	stderr    string
}

func Maketestprogprober(testname string, command string, args ...string) *testprogconfig {
	p := new(testprogconfig)
	p.testname = testname
	p.command = command
	p.args = args
	return p
}

func (p *testprogconfig) HealthCheck() bool {
	return p.healthy
}

func (p *testprogconfig) Probe() error {
	return p.probe()
}

func (p *testprogconfig) Register(dir *tricorder.DirectorySpec) error {
	return p.register(dir)
}
