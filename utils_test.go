package labassistant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func observation_order(slice []*Observation) string {
	var tmp string
	for _, ob := range slice {
		tmp += ob.Name
	}
	return tmp
}

// TODO: there's a chance that the shuffled order would match the original order...
func TestShuffleIsOK(t *testing.T) {
	a := assert.New(t)
	slice := []*Observation{
		&Observation{Name: "0"},
		&Observation{Name: "1"},
		&Observation{Name: "2"},
		&Observation{Name: "3"},
		&Observation{Name: "4"},
		&Observation{Name: "5"},
		&Observation{Name: "6"},
		&Observation{Name: "7"},
		&Observation{Name: "8"},
		&Observation{Name: "9"},
	}
	order1 := observation_order(slice)

	shuffle(slice)
	order2 := observation_order(slice)

	a.Len(slice, 10, "missing items from slice")
	a.NotEqual(order1, order2, "slice was not shuffled")
}

func TestIsFuncIsOK(t *testing.T) {
	a := assert.New(t)
	f := func() {}
	a.True(is_func(f))
	a.False(is_func(1))
}

func TestDefaultMisMatchCompareIsOK(t *testing.T) {
	a := assert.New(t)
	// Yep. I lack imagination.
	control := []interface{}{1, 2, 3, "string"}
	good := []interface{}{1, 2, 3, "string"}
	bad_value := []interface{}{1, 2, 3, ""}
	missing := []interface{}{1, 2, 3}
	empty := []interface{}{}
	too_many := []interface{}{1, 2, 3, 4, 5, 6, 7, "string", 9}

	a.False(DefaultMismatchCompare(control, good))
	a.True(DefaultMismatchCompare(control, bad_value))
	a.True(DefaultMismatchCompare(control, missing))
	a.True(DefaultMismatchCompare(control, empty))
	a.True(DefaultMismatchCompare(control, too_many))
}
