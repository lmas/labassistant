package labassistant

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type Observation struct {
	Name     string
	Panic    interface{}
	Outputs  []interface{}
	Start    time.Time
	Duration time.Duration
	Mismatch bool

	fun       interface{}
	can_panic bool
}

// Try to run the function for an observation, and if successful, set some stats.
func (ob *Observation) run(f interface{}, args ...interface{}) []interface{} {
	fv := reflect.ValueOf(f)
	if len(ob.Name) == 0 {
		if rf := runtime.FuncForPC(fv.Pointer()); rf != nil {
			ob.Name = rf.Name()
		}
	}

	fvtype := fv.Type()
	if len(args) != fvtype.NumIn() {
		panic(fmt.Errorf("Incorrect number of inputs to %v", ob.Name))
	}

	inputs := []reflect.Value{}
	for i, a := range args {
		tmp := reflect.ValueOf(a)
		tmptype := tmp.Type()
		in := fvtype.In(i)
		if tmptype != in {
			panic(fmt.Errorf("Invalid input (%v) to function (expected %v)",
				tmptype.Kind(),
				in.Kind(),
			))
		}
		inputs = append(inputs, tmp)
	}

	ret := ob.make_call(fv, inputs)
	if ob.Panic != nil {
		return nil
	}

	for _, r := range ret {
		ob.Outputs = append(ob.Outputs, r.Interface())
	}
	return ob.Outputs
}

// Do the actual call to the function and make some measurements.
func (ob *Observation) make_call(fv reflect.Value, inputs []reflect.Value) []reflect.Value {
	if ob.can_panic {
		defer func() {
			ob.Panic = recover()
		}()
	}

	ob.Start = time.Now()
	ret := fv.Call(inputs)
	ob.Duration = time.Since(ob.Start)
	return ret
}
