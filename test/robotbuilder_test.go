package test

import (
    "testing"
    "time"
    "micro-go-batch/pkg"
)

func TestRobotBuilder_AddComponent(t *testing.T) {
	// Initialize a RobotBuilder for testing
	builder := &pkg.RobotBuilder{ // Use pkg.RobotBuilder
			BatchSize:       2,
			BuildTime:       100 * time.Millisecond,
			ShutdownFlag:    false,
			Building:        false,
			ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock, keep in mind to use pkg.RobotComponent
	}

	// Mock RobotComponent
	component1 := &pkg.RobotComponent{Name: "Component1"}
	component2 := &pkg.RobotComponent{Name: "Component2"}
	component3 := &pkg.RobotComponent{Name: "Component3"}

	// Add a component
	builder.AddComponent(component1)
	builder.AddComponent(component2)
	builder.AddComponent(component3)

	// Test if component is added to the queue
	select {
	case <-builder.ComponentsQueue:
	default:
			t.Error("Failed to add component to the queue")
	}

	// Test adding more components to start building
	component4 := &pkg.RobotComponent{Name: "Component4"}
	builder.AddComponent(component4)

	// Test if building started
	if !builder.Building {
			t.Error("Building should have started after adding enough components")
	}
}

func TestRobotBuilder_AddComponentWithShutdownFlag(t *testing.T) {
	// Initialize a RobotBuilder for testing with ShutdownFlag set to true
	builder := &pkg.RobotBuilder{
			BatchSize:       2,
			BuildTime:       100 * time.Millisecond,
			ShutdownFlag:    true, // Set ShutdownFlag to true
			Building:        false,
			ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
	}

	// Attempt to add a component when ShutdownFlag is true
	component := &pkg.RobotComponent{Name: "Component1"}
	builder.AddComponent(component)

	// Test if component is not added to the queue
	select {
	case <-builder.ComponentsQueue:
			t.Error("Component should not be added to the queue when ShutdownFlag is true")
	default:
			// Component not added, as expected
	}
}

func TestRobotBuilder_StartBuildingWithInsufficientComponents(t *testing.T) {
	// Initialize a RobotBuilder for testing
	builder := &pkg.RobotBuilder{
			BatchSize:       3, // Set BatchSize larger than available components
			BuildTime:       100 * time.Millisecond,
			ShutdownFlag:    false,
			Building:        false,
			ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
	}

	// Add only one component
	component := &pkg.RobotComponent{Name: "Component1"}
	builder.AddComponent(component)

	// Attempt to start building with insufficient components
	builder.StartBuilding()

	// Ensure building doesn't start
	if builder.Building {
			t.Error("Building should not start with insufficient components")
	}
}

func TestRobotBuilder_StartBuildingWithZeroBuildTime(t *testing.T) {
	// Initialize a RobotBuilder for testing
	builder := &pkg.RobotBuilder{
		BatchSize:       2,
		BuildTime:       0, // Set BuildTime to zero
		ShutdownFlag:    false,
		Building:        false,
		ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
	}

	// Add components to start building
	component1 := &pkg.RobotComponent{Name: "Component1"}
	component2 := &pkg.RobotComponent{Name: "Component2"}
	builder.AddComponent(component1)
	builder.AddComponent(component2)

	// Start building
	builder.StartBuilding()

	// Ensure building starts immediately
	if !builder.Building {
		t.Error("Building should start immediately with zero build time")
	}

	// Ensure building flag is set to false after zero build time
	time.Sleep(100 * time.Millisecond) // Wait for a short period
	if builder.Building {
		t.Error("Building flag should be false after zero build time")
	}
}

func TestRobotBuilder_StartBuildingWithNoComponents(t *testing.T) {
	// Initialize a RobotBuilder for testing
	builder := &pkg.RobotBuilder{
			BatchSize:       2,
			BuildTime:       100 * time.Millisecond,
			ShutdownFlag:    false,
			Building:        false,
			ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
	}

	// Attempt to start building with no components
	builder.StartBuilding()

	// Ensure building doesn't start
	if builder.Building {
			t.Error("Building should not start with no components")
	}
}

func TestRobotBuilder_Shutdown(t *testing.T) {
	// Initialize a RobotBuilder for testing
	builder := &pkg.RobotBuilder{
			BatchSize:       2,
			BuildTime:       100 * time.Millisecond,
			ShutdownFlag:    false,
			Building:        false,
			ComponentsQueue: make(chan *pkg.RobotComponent, 2), // Buffered channel to prevent deadlock
	}

	// Shutdown the builder
	builder.Shutdown()

	// Test if shutdown flag is set
	if !builder.ShutdownFlag {
			t.Error("Shutdown flag should be true after calling Shutdown()")
	}
}