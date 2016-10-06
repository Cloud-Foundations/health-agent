package script

import (
	"github.com/Symantec/tricorder/go/tricorder"
	"github.com/Symantec/tricorder/go/tricorder/units"
)

func (p *scriptconfig) register(dir *tricorder.DirectorySpec) error {
	if err := dir.RegisterMetric("run-successful", &p.runSuccessful,
		units.None, "Run successful?"); err != nil {
		return err
	}
	if err := dir.RegisterMetric("exit-error", &p.exitError, units.None,
		"Error on exit if any"); err != nil {
		return err
	}
	return nil
}
