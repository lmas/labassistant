package labassistant

import "fmt"

type Experiment struct {
	Name       string
	Control    *Observation
	Candidates []*Observation
	RunOrder   []string
	Inputs     []interface{}
}

// Create a new experiment to run and set a name for it.
func NewExperiment(name string) *Experiment {
	ex := &Experiment{Name: name}
	return ex
}

// Add a function f as the control for the experiemnt.
// There can only be one control for each experiment and it's required to be
// set before the experiment can be run.
func (ex *Experiment) AddControl(f interface{}) {
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
		if len(ob.Outputs) != len(ex.Control.Outputs) {
			ob.Mismatch = true
			continue
		}
		for i := range ob.Outputs {
			if ob.Outputs[i] != ex.Control.Outputs[i] {
				ob.Mismatch = true
				continue
			}
		}
	}

	return ex.Control.Outputs
}
