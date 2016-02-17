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

	fun       interface{}
	can_panic bool
}

func (ob *Observation) run(f interface{}, args ...interface{}) []interface{} {
	fv := reflect.ValueOf(f)
	fvtype := fv.Type()
	if fvtype.Kind() != reflect.Func {
		panic(fmt.Errorf("First argument is not a func"))
	}

	inputs := []reflect.Value{}
	for i, a := range args {
		tmp := reflect.ValueOf(a)
		if tmp.Type() != fvtype.In(i) {
			panic(fmt.Errorf("Invalid argument to function"))
		}
		inputs = append(inputs, tmp)
	}

	if len(ob.Name) == 0 {
		if rf := runtime.FuncForPC(fv.Pointer()); rf != nil {
			ob.Name = rf.Name()
		}
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

// Do the actual call to the function
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
