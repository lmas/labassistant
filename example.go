package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/lmas/labassistant/src"
)

func main() {
	ex := labassistant.NewExperiment("test")
	ex.AddControl(testcontrol, 1, 2)
	ex.AddCandidate(testcan1)
	ex.AddCandidate(testcan2)
	ex.AddCandidate(testcan3)
	ex.AddCandidate(testcan4)
	ex.AddCandidate(testcan5)
	ex.AddCandidate(testcan6)
	ex.AddCandidate(testcan7)
	ex.Run()
	Publish(ex)
}

func Publish(ex *labassistant.Experiment) {
	fmt.Printf("%v duration: %v\t output: %v\n", ex.Control.Name, ex.Control.Duration, ex.Control.Outputs)

	for _, ob := range ex.Candidates {
		output := ob.Outputs
		if ob.Panic != nil {
			output = []interface{}{ob.Panic}
		}

		fmt.Printf("Candidate %v duration: %v \t output: %v\n", ob.Name, ob.Duration, output)
	}

	fmt.Println("Run order: ", strings.Join(ex.RunOrder, ", "))
}

func testcontrol(i1, i2 int) (int, int, int) {
	return i1, i2, i1 + i2
}

func testcan1(i1, i2 int) (int, int, int) {
	return i1, i2, i1 + i2
}

// The rest of the candidates are bad

func testcan2(i1, i2 int) (int, int, int) {
	return i1, i2, -1
}

func testcan3(i1, i2 int) error {
	return fmt.Errorf("bad output")
}

func testcan4(i1, i2 int) (int, int, int) {
	i := 1 - 1
	return i1, i2, 1 / i
}

func testcan5(i1, i2 int) (int, int) {
	return i1, i2
}

func testcan6(i1, i2 int) (int, int, int, int) {
	return i1, i2, i1 + i2, 0
}

func testcan7(i1, i2 int) (int, int, int) {
	time.Sleep(10000 * time.Nanosecond)
	return i1, i2, i1 + i2
}
