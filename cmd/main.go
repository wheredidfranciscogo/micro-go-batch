package main

import (
	"fmt"
	"time"

	"micro-go-batch/pkg"
)

func main() {
	// Create a RobotBuilder instance
	builder := &pkg.RobotBuilder{
		BatchSize: 5,
		BuildTime: 4 * time.Second,
		HandleBatchCompletion: func(robot *pkg.Robot, results []pkg.JobResult) {
			// This function will be called when a batch of robots is assembled
			fmt.Println("Robot assembled:", robot.SerialNumber)
			for _, result := range results {
				fmt.Println("Result:", result.Data)
			}
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
