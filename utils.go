package labassistant

import (
	"math/rand"
	"reflect"
	"time"
)

// Shuffle a slice
func shuffle(slice []*Observation) {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

}

// Check if a variable is a function.
func is_func(f interface{}) bool {
	ftype := reflect.ValueOf(f).Type().Kind()
	if ftype != reflect.Func {
		return false
	}
	return true
}

// The default output mismatch comparison.
// It loops over all outputs for a candidate and do a simple == comparison
// against the control's output. It also ensures that the number of outputs of
// the candidate is the same as the control's.
func DefaultMismatchCompare(control, candidate []interface{}) bool {
	if len(candidate) != len(control) {
		return false
	}
	for i := range candidate {
		if candidate[i] != control[i] {
			return false
		}
	}
	return true
}
