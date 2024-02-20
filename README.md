# micro-go-batch

In this project I'm going to simulate the assembly of robots using a series of classes and functions. The idea is to implement a micro-batching system where individual tasks (represented as components needed to assemble a robot) are grouped together and processed in batches. This package aim for handling batch processing of data. Upon completion of each default batch, a special number is assigned to the assembled robot.

## Overview

Micro Go Batch provides a framework for efficiently processing data in batches. It includes a default implementation of the `BatchProcessor` interface, which users can leverage for basic batch processing needs. Additionally, users can create custom `BatchProcessor` implementations to tailor the batch processing behavior to their specific requirements.

## Usage

### Running the Program

To run the program, follow these steps:
```
git clone <repo>
cd micro-go-batch
go run cmd/main.go
```

### Running the Tests

To run the tests for the package, execute the following command from the root of the project directory:
```sh
go test ./test
```

### Adding a Custom BatchProcessor

Users can add their own custom BatchProcessor implementations to extend the functionality of Micro Go Batch. Here's how to do it:

Implement the BatchProcessor interface in a new Go file.
```go
// MyBatchProcessor implements the BatchProcessor interface
type MyBatchProcessor struct{}

// Process is the implementation of the Process method of the BatchProcessor interface
func (mbp *MyBatchProcessor) Process(robot *Robot, results []JobResult) {
    // Custom batch processing logic here
}

```

Use the new BatchProcessor in the RobotBuilder configuration.
```go
// Create a MyBatchProcessor instance
myProcessor := &MyBatchProcessor{}

// Create a RobotBuilder instance with the MyBatchProcessor
builder := &pkg.RobotBuilder{
    BatchSize:      5,
    BuildTime:      4 * time.Second,
    BatchProcessor: myProcessor,
    HandleBatchCompletion: func(robot *pkg.Robot, results []pkg.JobResult) {
        // Custom completion handling logic here
    },
    ComponentsQueue: make(chan *pkg.RobotComponent, 10),
}

```

### Language Used
The project is attemped to be written in Go (Golang).

### Author
Francisco Arrieta