package test

import (
    "fmt"
    "testing"
    "time"
    "micro-go-batch/pkg"
)

// MockBatchProcessor implements the BatchProcessor interface for testing purposes
type MockBatchProcessor struct {
    Calls            int
    expectedRobot    *pkg.Robot
    expectedResults  []pkg.JobResult
}

// Process is the mock implementation of the Process method in BatchProcessor interface
func (m *MockBatchProcessor) Process(robot *pkg.Robot, results []pkg.JobResult) {
    // Your mock implementation here, for example, you can keep track of the number of calls
    m.Calls++
}

// ExpectHandleBatchCompletion sets the expected robot and results for the next call to Process method.
func (m *MockBatchProcessor) ExpectHandleBatchCompletion(robot *pkg.Robot, results []pkg.JobResult) {
    m.expectedRobot = robot
    m.expectedResults = results
}


func TestRobotBuilder_AddComponent(t *testing.T) {
    // Define a function to handle batch completion
    handleBatchCompletion := func(robot *pkg.Robot, results []pkg.JobResult) {
        fmt.Println("Batch completed")
        fmt.Println("Robot assembled:", robot.SerialNumber)
        for _, result := range results {
            fmt.Println("Result:", result.Data)
        }
    }

    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor
    builder := &pkg.RobotBuilder{
        BatchSize:             1,
        BuildTime:             100 * time.Millisecond,
        HandleBatchCompletion: handleBatchCompletion,
        ComponentsQueue:       make(chan *pkg.RobotComponent, 2),
        BatchProcessor:        mockProcessor, // Pass the mock processor
    }

    // Add a component
    component1 := &pkg.RobotComponent{Name: "Component1"}
    builder.AddComponent(component1)

    // Check if the StartBuilding method is called
    if !builder.Building {
        t.Errorf("Expected StartBuilding method to be called, but it wasn't")
    }

    // Wait for a short duration to ensure StartBuilding has completed
    time.Sleep(200 * time.Millisecond)

    // Check if the mock processor's Process method was called
    if mockProcessor.Calls != 1 {
        t.Errorf("Expected Process method to be called once, got %d calls", mockProcessor.Calls)
    }
}

func TestRobotBuilder_AddComponentWithShutdownFlag(t *testing.T) {
    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor and ShutdownFlag set to true
    builder := &pkg.RobotBuilder{
        BatchSize:       2,
        BuildTime:       100 * time.Millisecond,
        ShutdownFlag:    true, // Set ShutdownFlag to true
        Building:        false,
        ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
        BatchProcessor:  mockProcessor,                      // Pass the mock processor
    }

    // Attempt to add a component when ShutdownFlag is true
    component := &pkg.RobotComponent{Name: "Component1"}
    builder.AddComponent(component)

    // Ensure that the HandleBatchCompletion function of the BatchProcessor is not called
    if mockProcessor.Calls != 0 {
        t.Error("HandleBatchCompletion should not be called when ShutdownFlag is true")
    }
}

func TestRobotBuilder_StartBuildingWithInsufficientComponents(t *testing.T) {
    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor
    builder := &pkg.RobotBuilder{
        BatchSize:       3,
        BuildTime:       100 * time.Millisecond,
        ShutdownFlag:    false,
        Building:        false,
        ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
        BatchProcessor:  mockProcessor,                      // Pass the mock processor
    }

    // Add only one component
    component := &pkg.RobotComponent{Name: "Component1"}
    builder.AddComponent(component)

    // Attempt to start building with insufficient components
    builder.StartBuilding()

    // Ensure that the HandleBatchCompletion function of the BatchProcessor is not called
    if mockProcessor.Calls != 0 {
        t.Error("HandleBatchCompletion should not be called with insufficient components")
    }
}

func TestRobotBuilder_StartBuildingWithZeroBuildTime(t *testing.T) {
    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor and zero BuildTime
    builder := &pkg.RobotBuilder{
        BatchSize:       2,
        BuildTime:       0, // Set BuildTime to zero
        ShutdownFlag:    false,
        Building:        false,
        ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
        BatchProcessor:  mockProcessor,                      // Pass the mock processor
    }

    // Add components to start building
    component1 := &pkg.RobotComponent{Name: "Component1"}
    component2 := &pkg.RobotComponent{Name: "Component2"}
    builder.AddComponent(component1)
    builder.AddComponent(component2)

    // Start building
    builder.StartBuilding()

    // Ensure that the HandleBatchCompletion function of the BatchProcessor is called immediately
    expectedResults := []pkg.JobResult{
        {Data: "Component1 assembled"},
        {Data: "Component2 assembled"},
    }
    expectedRobot := &pkg.Robot{Components: []pkg.RobotComponent{*component1, *component2}}
    mockProcessor.ExpectHandleBatchCompletion(expectedRobot, expectedResults)

    // Ensure that the Building flag is set to false after zero build time
    time.Sleep(100 * time.Millisecond) // Wait for a short period
    if builder.Building {
        t.Error("Building flag should be false after zero build time")
    }
}

func TestRobotBuilder_StartBuildingWithNoComponents(t *testing.T) {
    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor
    builder := &pkg.RobotBuilder{
        BatchSize:       2,
        BuildTime:       100 * time.Millisecond,
        ShutdownFlag:    false,
        Building:        false,
        ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
        BatchProcessor:  mockProcessor,                      // Pass the mock processor
    }

    // Attempt to start building with no components
    builder.StartBuilding()

    // Ensure that the HandleBatchCompletion function of the BatchProcessor is not called
    if mockProcessor.Calls != 0 {
        t.Error("HandleBatchCompletion should not be called with no components")
    }
}

func TestRobotBuilder_Shutdown(t *testing.T) {
    // Initialize a mock BatchProcessor
    mockProcessor := &MockBatchProcessor{}

    // Initialize a RobotBuilder for testing with the BatchProcessor
    builder := &pkg.RobotBuilder{
        BatchSize:       2,
        BuildTime:       100 * time.Millisecond,
        ShutdownFlag:    false,
        Building:        false,
        ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
        BatchProcessor:  mockProcessor,                      // Pass the mock processor
    }

    // Shutdown the builder
    builder.Shutdown()

    // Test if shutdown flag is set
    if !builder.ShutdownFlag {
        t.Error("Shutdown flag should be true after calling Shutdown()")
    }
}
