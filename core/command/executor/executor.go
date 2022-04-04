package executor

import (
	"time"

	"github.com/jahid90/just/core/logger"
)

type Executor struct {
	ctx   *executionContext
	units []*ExecutionUnit
}

func NewExecutor(units []*ExecutionUnit) *Executor {
	return &Executor{
		ctx: &executionContext{
			skipUnitFailures: false,
		},
		units: units,
	}
}

func NewSkipFailuresExecutor(units []*ExecutionUnit) *Executor {
	return &Executor{
		ctx: &executionContext{
			skipUnitFailures: true,
		},
		units: units,
	}
}

// Execute Executes a slice of execution units
func (e *Executor) Execute() error {

	if len(e.units) == 0 {
		logger.Warn("no units passed")
		return nil
	}

	start := time.Now()

	for _, unit := range e.units {
		if err := unit.execute(); err != nil {
			if e.ctx.skipUnitFailures {
				logger.Error(err.Error())
			} else {
				return err
			}
		}
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}
