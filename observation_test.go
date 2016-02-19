package labassistant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func make_ob() *Observation {
	return &Observation{}
}

func good_func(i int) int {
	return i
}

func panic_func(i int) int {
	panic("panic")
}

func bad_func(i int) int {
	return i + 1
}

func TestObservationRunIsOK(t *testing.T) {
	a := assert.New(t)
	out := make_ob().run(good_func, 1)
	a.Equal([]interface{}{1}, out)

	out = make_ob().run(good_func, 0)
	a.NotEqual([]interface{}{1}, out)

	// Missing inputs
	a.Panics(func() {
		make_ob().run(good_func)
	})

	// Too many inputs
	a.Panics(func() {
		make_ob().run(good_func, 1, 1)
	})

	// Invalid input type
	a.Panics(func() {
		make_ob().run(good_func, "string")
	})
}

func TestObservationRunSetsValues(t *testing.T) {
	a := assert.New(t)
	good := make_ob()
	good.run(good_func, 1)
	a.Equal("github.com/lmas/labassistant.good_func", good.Name)
	a.Empty(good.Panic)
	a.Equal([]interface{}{1}, good.Outputs)
	a.NotEmpty(good.Start)
	a.NotZero(good.Duration)

	bad := make_ob()
	bad.can_panic = true
	a.NotPanics(func() {
		bad.run(panic_func, 1)
	})
	a.Equal("panic", bad.Panic)
	a.Empty(bad.Outputs)
	a.NotEmpty(bad.Start)
	a.Zero(bad.Duration)
}
