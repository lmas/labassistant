package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/lmas/labassistant"
)

func main() {
	// Make a new experiment and call it "test".
	ex := labassistant.NewExperiment("test")

	// Add the control, the original function you want to refactor and
	// measure against with your new code.
	ex.AddControl(testcontrol, 1, 2)

	// Add a bunch of new candidates for replacing the original example function.
	ex.AddCandidate(testcan1)
	ex.AddCandidate(testcan2)
	ex.AddCandidate(testcan3)
	ex.AddCandidate(testcan4)
	ex.AddCandidate(testcan5)
	ex.AddCandidate(testcan6)
	ex.AddCandidate(testcan7)

	// Fire up the engines and start running the control and the candidates
	// in a random order. The library will take some measurements while this
	// is running.
	ex.Run()

	// Example function for showing the final results of the run.
	Publish(ex)
}

func Publish(ex *labassistant.Experiment) {
	fmt.Printf("%v duration: %v\t output: %v\n", ex.Control.Name, ex.Control.Duration, ex.Control.Outputs)

	for _, ob := range ex.Candidates {
		// By default we show the output of the candidate functions.
		output := ob.Outputs

		// If any of the candidates would panic during the run, show that
		// instead.
		if ob.Panic != nil {
			output = []interface{}{ob.Panic}
		}

		fmt.Printf("Candidate %v duration: %v \t output: %v\n", ob.Name, ob.Duration, output)
	}

	// And finally show the execution order.
	fmt.Println("Run order: ", strings.Join(ex.RunOrder, ", "))
}

// This is the original func we would like to refactor.
func testcontrol(i1, i2 int) (int, int, int) {
	return i1, i2, i1 + i2
}

// This is a good refactor of previous func (I lack imagination so it's actually
// not a refactor yet, just a straight copy for simplicity's sake D: ).
func testcan1(i1, i2 int) (int, int, int) {
	return i1, i2, i1 + i2
}

// The rest of the candidates are bad.

// This func will probably return a wrong 3rd value.
func testcan2(i1, i2 int) (int, int, int) {
	return i1, i2, -1
}

// This one is guaranteed to return a wrong value.
func testcan3(i1, i2 int) error {
	return fmt.Errorf("bad output")
}

// Sneaky way to raise a runtime panic.
func testcan4(i1, i2 int) (int, int, int) {
	i := 1 - 1
	return i1, i2, 1 / i
}

// This one simply lacks all required outputs.
func testcan5(i1, i2 int) (int, int) {
	return i1, i2
}

// And this one have too many of them.
func testcan6(i1, i2 int) (int, int, int, int) {
	return i1, i2, i1 + i2, 0
}

// Finally this is just a slower and less efficient "refacted" func.
func testcan7(i1, i2 int) (int, int, int) {
	time.Sleep(10000 * time.Nanosecond)
	return i1, i2, i1 + i2
}
