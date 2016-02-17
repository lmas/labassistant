package labassistant

import (
	"math/rand"
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
