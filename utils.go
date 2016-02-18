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
