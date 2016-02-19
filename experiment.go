package labassistant

import "fmt"

type Experiment struct {
	Name       string
	Control    *Observation
	Candidates []*Observation
	RunOrder   []string
	Inputs     []interface{}

	mismatch_compare func([]interface{}, []interface{}) bool
}

// Create a new experiment to run and set a name for it.
func NewExperiment(name string) *Experiment {
	ex := &Experiment{Name: name, mismatch_compare: DefaultMismatchCompare}
	return ex
}

// Set a function f as the control for the experiemnt.
// There can only be one control for each experiment and it's required to be
// set before the experiment can be run.
func (ex *Experiment) SetControl(f interface{}) {
	if !is_func(f) {
		panic(fmt.Errorf("Control is not a function"))
	}
	ex.Control = &Observation{Name: "Control", can_panic: false, fun: f}
}

// Add a candidate function f. You can add as many as you like.
func (ex *Experiment) AddCandidate(f interface{}) {
	if !is_func(f) {
		panic(fmt.Errorf("Candidate is not a function"))
	}
	ob := &Observation{can_panic: true, fun: f}
	ex.Candidates = append(ex.Candidates, ob)
}

// Set a custom comparison function.
// It is run once for each candidate.
//
// Format of the custom comparison function:
// func name_of_func(control []interface{}, candidate []interface{}) bool{}
// and it should return wether the control outputs "control" is equal to the
// candidate outputs "candidate".
func (ex *Experiment) SetCompare(f func([]interface{}, []interface{}) bool) {
	ex.mismatch_compare = f
}

// Run the experiment, calling the control and candidate functions, one at a
// time, in a random order.
// You can optionally provide input arguments to give to the control/candidates.
func (ex *Experiment) Run(inputs ...interface{}) []interface{} {
	if ex.Control == nil {
		panic("Experiment must have a control case.")
	}
	ex.Inputs = inputs

	// Make a copy of the candidate list or else we will accidently modify it
	// because of pointer magic.
	all := append([]*Observation(nil), ex.Candidates...)
	all = append(all, ex.Control)
	shuffle(all)

	for _, ob := range all {
		ob.run(ob.fun, ex.Inputs...)
		ex.RunOrder = append(ex.RunOrder, ob.Name)
	}

	// Check for output mismatches now that we've run everything
	for _, ob := range ex.Candidates {
		if ex.mismatch_compare(ex.Control.Outputs, ob.Outputs) == false {
			ob.Mismatch = true
		}
	}

	return ex.Control.Outputs
}
