package pkg

// BatchProcessor interface defines the method for processing batches of results
type BatchProcessor interface {
    Process(*Robot, []JobResult)
}