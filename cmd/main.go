package main

import (
	"micro-go-batch/pkg"
	"time"
)

func main() {
	// Create a DefaultBatchProcessor instance
	defaultProcessor := pkg.DefaultBatchProcessor{}

	// Create a RobotBuilder instance with the DefaultBatchProcessor
	builder := &pkg.RobotBuilder{
		BatchSize:      5,
		BuildTime:      4 * time.Second,
		BatchProcessor: &defaultProcessor, // Pass the address of the instance
		HandleBatchCompletion: func(robot *pkg.Robot, results []pkg.JobResult) {
		},
		ComponentsQueue: make(chan *pkg.RobotComponent, 10), // Add buffer to the channel
	}

	// Add components to the builder
	components := []*pkg.RobotComponent{
		{Name: "chassis"},
		{Name: "motor"},
		{Name: "sensor"},
		{Name: "arm"},
		{Name: "processor"},
		{Name: "lights"},
		{Name: "bulbs"},
	}

	// Start a goroutine to add components to the builder
	go func() {
		for _, component := range components {
			builder.AddComponent(component)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Wait for components to be processed
	time.Sleep(10 * time.Second)
}
