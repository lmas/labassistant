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
	var ex *Experiment
	a.NotPanics(func() {
		ex = NewExperiment("", good_func)
	})
	a.NotEmpty(ex.Control)
}

func TestSetBadControl(t *testing.T) {
	a := assert.New(t)
	var ex *Experiment
	a.Panics(func() {
		ex = NewExperiment("", "not_a_func")
	})
	a.Empty(ex)
}

func TestAddCandidateIsOK(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	a.Empty(ex.Candidates)
	a.NotPanics(func() {
		ex.AddCandidate(good_func)
	})

	a.Panics(func() {
		ex.AddCandidate("string")
	})

	a.NotEmpty(ex.Candidates)
	a.NotPanics(func() {
		ex.AddCandidate(panic_func)
	})

	a.Equal(2, len(ex.Candidates))
	a.NotEqual(ex.Candidates[0], ex.Candidates[1])
}

func TestRunIsOK(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	a.Panics(func() {
		ex.Run()
	})

	ex.AddCandidate(good_func)
	a.NotPanics(func() {
		ex.Run(1)
	})
	a.Equal([]interface{}{1}, ex.Inputs)
	a.Equal([]interface{}{1}, ex.Control.Outputs)
	a.NotEmpty(ex.RunOrder)

	a.Empty(ex.Candidates[0].Panic)
	a.False(ex.Candidates[0].Mismatch)
	a.Equal(ex.Control.Outputs, ex.Candidates[0].Outputs)
}

func TestRunWithPanicCandidates(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	ex.AddCandidate(panic_func)
	a.NotPanics(func() {
		ex.Run(1)
	})
	a.NotEmpty(ex.Candidates[0].Panic)
	a.True(ex.Candidates[0].Mismatch)

}

func TestRunWithBadCandidates(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	ex.AddCandidate(bad_func)
	a.NotPanics(func() {
		ex.Run(1)
	})
	a.Empty(ex.Candidates[0].Panic)
	a.True(ex.Candidates[0].Mismatch)
	a.NotEqual(ex.Control.Outputs, ex.Candidates[0].Outputs)

}

func TestRunWithCustomCompare(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	ex.AddCandidate(good_func)
	ex.SetCompare(fail_compare)
	ex.Run(1)
	a.Equal([]interface{}{1}, ex.Candidates[0].Outputs)
	a.True(ex.Candidates[0].Mismatch)
}

func TestRunWithCustomIgnore(t *testing.T) {
	a := assert.New(t)
	ex := NewExperiment("", good_func)
	ex.AddCandidate(bad_func)
	ex.SetIgnore(ignore_all)
	ex.Run(1)
	a.Equal([]interface{}{2}, ex.Candidates[0].Outputs)
	a.Empty(ex.Candidates[0].Panic)
	a.False(ex.Candidates[0].Mismatch)
}
