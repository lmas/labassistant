package labassistant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fail_compare(co, ca []interface{}) bool {
	// This "comparison" will always fail.
	return true
}

func ignore_all(co, ca []interface{}) bool {
	// This "ignore" will always ignore.
	return true
}

func TestSetControlIsOK(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	a.Zero(e.Control)
	a.NotPanics(func() {
		e.SetControl(good_func)
	})
	a.NotEmpty(e.Control)

	a.Panics(func() {
		e.SetControl("string")
	})
}

func TestAddCandidateIsOK(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	a.Empty(e.Candidates)
	a.NotPanics(func() {
		e.AddCandidate(good_func)
	})

	a.Panics(func() {
		e.AddCandidate("string")
	})

	a.NotEmpty(e.Candidates)
	a.NotPanics(func() {
		e.AddCandidate(panic_func)
	})

	a.Equal(2, len(e.Candidates))
	a.NotEqual(e.Candidates[0], e.Candidates[1])
}

func TestRunIsOK(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	a.Panics(func() {
		e.Run()
	})

	e.SetControl(good_func)
	e.AddCandidate(good_func)
	a.NotPanics(func() {
		e.Run(1)
	})
	a.Equal([]interface{}{1}, e.Inputs)
	a.Equal([]interface{}{1}, e.Control.Outputs)
	a.NotEmpty(e.RunOrder)

	a.Empty(e.Candidates[0].Panic)
	a.False(e.Candidates[0].Mismatch)
	a.Equal(e.Control.Outputs, e.Candidates[0].Outputs)
}

func TestRunWithPanicCandidates(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	e.SetControl(good_func)
	e.AddCandidate(panic_func)
	a.NotPanics(func() {
		e.Run(1)
	})
	a.NotEmpty(e.Candidates[0].Panic)
	a.True(e.Candidates[0].Mismatch)

}

func TestRunWithBadCandidates(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	e.SetControl(good_func)
	e.AddCandidate(bad_func)
	a.NotPanics(func() {
		e.Run(1)
	})
	a.Empty(e.Candidates[0].Panic)
	a.True(e.Candidates[0].Mismatch)
	a.NotEqual(e.Control.Outputs, e.Candidates[0].Outputs)

}

func TestRunWithCustomCompare(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	e.SetControl(good_func)
	e.AddCandidate(good_func)
	e.SetCompare(fail_compare)
	e.Run(1)
	a.Equal([]interface{}{1}, e.Candidates[0].Outputs)
	a.True(e.Candidates[0].Mismatch)
}

func TestRunWithCustomIgnore(t *testing.T) {
	a := assert.New(t)
	e := NewExperiment("")
	e.SetControl(good_func)
	e.AddCandidate(bad_func)
	e.SetIgnore(ignore_all)
	e.Run(1)
	a.Equal([]interface{}{2}, e.Candidates[0].Outputs)
	a.Empty(e.Candidates[0].Panic)
	a.False(e.Candidates[0].Mismatch)
}
