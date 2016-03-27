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
	return reflect.ValueOf(f).Type().Kind() == reflect.Func
}

// The default output mismatch comparison.
// If there's a mismatch it will return true.
//
// It loops over all outputs for a candidate and do a simple == comparison
// against the control's output. It also ensures that the number of outputs of
// the candidate is the same as the control's.
func DefaultMismatchCompare(control, candidate []interface{}) bool {
	// We will likely run into a "index out of range" panic if we don't make
	// sure the len of outputs are equal.
	if len(candidate) != len(control) {
		return true
	}

	for i := range candidate {
		if candidate[i] != control[i] {
			return true
		}
	}
	return false
}

// The default ignore mismatch check.
// It always returns false, so it will never ignore any output mismatches.
func DefaultIgnoreMismatch(control, candidate []interface{}) bool {
	return false
}
