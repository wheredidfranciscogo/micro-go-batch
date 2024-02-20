package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot represents an assembled robot
type Robot struct {
	Components   []RobotComponent
	SerialNumber string
}

// RobotComponent represents a component of a robot
type RobotComponent struct {
	Name string
}

// creates a new Robot instance with the given components
func NewRobot(components []RobotComponent) *Robot {
	serialNumber := GenerateSpecialNumber()
	fmt.Println("Generated serial number:", serialNumber)
	return &Robot{
		Components:   components,
		SerialNumber: serialNumber,
	}
}

// init initializes the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateSpecialNumber generates a special serial number for the robot
func GenerateSpecialNumber() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialNumber := make([]byte, 6)
	for i := range specialNumber {
		specialNumber[i] = letters[rand.Intn(len(letters))]
	}
	return string(specialNumber)
}
