package pkg

import (
	"fmt"
	"time"
)

type RobotBuilder struct {
	BatchSize             int
	BuildTime             time.Duration
	BatchProcessor        BatchProcessor
	HandleBatchCompletion func(*Robot, []JobResult)
	ComponentsQueue       chan *RobotComponent
	Batch                 []RobotComponent
	ShutdownFlag          bool
	Building              bool
}

// AddComponent adds a component to the builder's queue and starts building if necessary.
func (rb *RobotBuilder) AddComponent(component *RobotComponent) {
	if !rb.ShutdownFlag {
		select {
		case rb.ComponentsQueue <- component: // pushing component to the queue
			if !rb.Building && len(rb.ComponentsQueue) >= rb.BatchSize {
				rb.StartBuilding()
			}
		default:
		}
	} else {
		fmt.Println("No more jobs. RobotBuilder is down.")
	}
}

// StartBuilding starts the process of building a batch of robots.
func (rb *RobotBuilder) StartBuilding() {
	rb.Building = true
	var batch []RobotComponent
	for i := 0; i < rb.BatchSize; i++ {
		select {
		case component, ok := <-rb.ComponentsQueue:
			if !ok {
				// Channel closed, stop building
				rb.Building = false
				return
			}
			batch = append(batch, *component)
		default:
			// No more components in the queue, stop building
			rb.Building = false
			return
		}
	}
	// Check if batch is empty
	if len(batch) < rb.BatchSize {
		rb.Building = false
		return
	}
	// If build time is zero, set building flag to false immediately
	if rb.BuildTime == 0 {
		rb.Building = false
		return
	}
	// Launch goroutine to simulate the construction time.
	go func() {
		time.Sleep(rb.BuildTime)
		robot := &Robot{Components: batch, SerialNumber: GenerateSpecialNumber()}
		var results []JobResult
		for _, component := range batch {
			results = append(results, JobResult{Data: component.Name + " assembled"})
		}

		// Call the Process method of the BatchProcessor interface
		rb.BatchProcessor.Process(robot, results)

		rb.HandleBatchCompletion(robot, results)

		if len(rb.ComponentsQueue) >= rb.BatchSize {
			rb.StartBuilding()
		} else {
			rb.Building = false // No more components to build, set flag to false.
		}
	}()
}

// Shutdown method which returns after all previously accepted Jobs are processed
func (rb *RobotBuilder) Shutdown() {
	rb.ShutdownFlag = true
}
