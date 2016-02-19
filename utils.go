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

// By default an experiment will do a simple `control output[x] == candidate output[x]`
// comparison for each output of each candidate.
func DefaultMismatchCompare(control, candidate interface{}) bool {
	return control == candidate
}
