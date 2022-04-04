package executor

import (
	"os/exec"
	"time"

	"github.com/jahid90/just/core/logger"
)

type Executor struct {
	ctx   *executionContext
	units []*executionUnit
}

func NewExecutor() *Executor {
	return &Executor{
		ctx: &executionContext{
			skipUnitFailures: false,
		},
		units: make([]*executionUnit, 0),
	}
}

func NewFailuresSkippingExecutor() *Executor {
	return &Executor{
		ctx: &executionContext{
			skipUnitFailures: true,
		},
		units: make([]*executionUnit, 0),
	}
}

func (e *Executor) AddExecutionUnit(cmd *exec.Cmd, description string) {
	newUnit := &executionUnit{
		cmd:         cmd,
		description: description,
	}

	e.units = append(e.units, newUnit)
}

// Execute Executes a slice of execution units
func (e *Executor) Execute() error {

	if len(e.units) == 0 {
		logger.Warn("no units added; nothing to do")
		return nil
	}

	start := time.Now()

	for _, unit := range e.units {
		if err := unit.execute(); err != nil {
			if e.ctx.skipUnitFailures {
				logger.Warn("step failed:", err.Error())
			} else {
				return err
			}
		}
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}
