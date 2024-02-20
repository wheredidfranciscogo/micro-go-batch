package test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
	"micro-go-batch/pkg"
)

func TestDefaultBatchProcessor_Process(t *testing.T) {
	// Create a mock robot
	robot := &pkg.Robot{SerialNumber: "TEST123", Components: []pkg.RobotComponent{{Name: "component1"}, {Name: "component2"}}}

	// Create a mock job result slice
	results := []pkg.JobResult{{Data: "result1"}, {Data: "result2"}}

	// Create a DefaultBatchProcessor instance
	defaultProcessor := pkg.DefaultBatchProcessor{
		BatchSize: 2,
		BuildTime: 100 * time.Millisecond,
	}

	// Redirect stdout to capture printed messages
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the Process method
	defaultProcessor.Process(robot, results)

	// Capture the output
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	// Convert output bytes to string for easier assertion
	output := string(out)

	// Assertions
	// Check if the "Default batch processing..." message is printed
	expectedMsg := "Default batch processing..."
	if !strings.Contains(output, expectedMsg) {
		t.Errorf("Expected to print '%s', but got: %s", expectedMsg, output)
	}

	// Check if the "Robot assembled:" message is printed with the correct serial number
	expectedRobotMsg := "Robot assembled: TEST123"
	if !strings.Contains(output, expectedRobotMsg) {
		t.Errorf("Expected to print '%s', but got: %s", expectedRobotMsg, output)
	}

	// Check if each job result is printed
	expectedResult1 := "Result: result1"
	if !strings.Contains(output, expectedResult1) {
		t.Errorf("Expected to print '%s', but got: %s", expectedResult1, output)
	}
	expectedResult2 := "Result: result2"
	if !strings.Contains(output, expectedResult2) {
		t.Errorf("Expected to print '%s', but got: %s", expectedResult2, output)
	}
}
