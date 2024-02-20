package pkg

import (
  "fmt"
  "time"
)

// BatchProcessor interface defines the method for processing batches of results
type BatchProcessor interface {
	Process(*Robot, []JobResult)
}

// DefaultBatchProcessor is a default implementation of the BatchProcessor interface
type DefaultBatchProcessor struct{
  BatchSize int           // Batch size for processing
	BuildTime time.Duration // Time interval for building a batch
}

// Process is the implementation of the Process method of the BatchProcessor interface
func (dbp *DefaultBatchProcessor) Process(robot *Robot, results []JobResult) {
	fmt.Println("Default batch processing...")

  // Simulate batch processing time
	time.Sleep(dbp.BuildTime)

  // This function will be called when a batch of robots is assembled
	fmt.Println("Robot assembled:", robot.SerialNumber)
	for _, result := range results {
		fmt.Println("Result:", result.Data)
	}

	// Log each job result
	//  for _, result := range results {
	//     fmt.Println("Processed job result:", result.Data)
	// }

	// // Calculate total and count
	// var total int
	// var count int
	// for _, result := range results {
	//     if num, ok := result.Data.(int); ok {
	//         total += num
	//         count++
	//     }
	// }

	// // Calculate average
	// average := float64(total) / float64(count)

	// // Log the aggregated results
	// fmt.Printf("Total: %d, Average: %.2f\n", total, average)

	// // Print out the assembled robot's information
	// fmt.Println("Assembled Robot:")
	// fmt.Printf("Serial Number: %s\n", robot.SerialNumber)
	// fmt.Println("Components:")
	// for _, component := range robot.Components {
	// 	fmt.Printf("- %s\n", component.Name)
	// }

	// This function will be called when a batch of robots is assembled
	// fmt.Println("Robot assembled:", robot.SerialNumber)
	// for _, result := range results {
	// 	fmt.Println("Result:", result.Data)
	// }
}
