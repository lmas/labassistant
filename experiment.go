package labassistant

type Experiment struct {
	Name       string
	Control    *Observation
	Candidates []*Observation
	RunOrder   []string

	args []interface{}
}

func NewExperiment(name string) *Experiment {
	ex := &Experiment{Name: name}
	return ex
}

func (ex *Experiment) AddControl(f interface{}, args ...interface{}) {
	ex.Control = &Observation{Name: "Control", can_panic: false, fun: f}
	ex.args = args
}

func (ex *Experiment) AddCandidate(f interface{}) {
	ob := &Observation{can_panic: true, fun: f}
	ex.Candidates = append(ex.Candidates, ob)
}

func (ex *Experiment) Run() []interface{} {
	if ex.Control == nil {
		panic("Experiment must have a control case.")
	}

	// Make a copy of the candidate list, or else we will modify it (pointer)
	all := append([]*Observation(nil), ex.Candidates...)
	all = append(all, ex.Control)
	shuffle(all)

	for _, ob := range all {
		ob.run(ob.fun, ex.args...)
		ex.RunOrder = append(ex.RunOrder, ob.Name)
	}

	return ex.Control.Outputs
}
