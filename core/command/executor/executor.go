package executor

import (
	"time"

	"github.com/jahid90/just/core/logger"
)

// ExecuteManyUnits Executes a slice of execution units; stops when any one fails
func ExecuteMany(units []*ExecutionUnit) error {

	if len(units) == 0 {
		logger.Warn("no units passed")
		return nil
	}

	start := time.Now()

	for _, unit := range units {
		err := unit.execute()
		if err != nil {
			return err
		}
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}

// Execute Runs an execution unit
func Execute(unit *ExecutionUnit) error {

	start := time.Now()

	err := unit.execute()
	if err != nil {
		return err
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}
