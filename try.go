package try

import (
	"fmt"
	"reflect"
)

type Try struct {
	F interface{}
}

func (p Try) Catch(exhandle interface{}) Catch{
	return catch{t:p, c:exhandle}
}

type Catch interface {
	Go(args ...interface{}) []reflect.Value
}

type catch struct {
	t Try
	c interface{}
}

func (p catch) Go(args ...interface{}) []reflect.Value {
	f := p.t.F
	rv := reflect.ValueOf(f)
	rt := reflect.TypeOf(f)
	defer func() {
		if err := recover(); err != nil {
			reflect.ValueOf(p.c).Call([]reflect.Value{reflect.ValueOf(fmt.Errorf("%s", err))})
		}
	}()
	switch rt.Kind() {
	case reflect.Func:
		in := make([]reflect.Value, rv.Type().NumIn())
		for i, a := range args {
			in[i] = reflect.ValueOf(a).Convert(rv.Type().In(i))
		}
		outputs :=  rv.Call(in)
		for _, out := range outputs {
			if _, ok := out.Interface().(error); ok {
				reflect.ValueOf(p.c).Call([]reflect.Value{out})
			}
		}
		return outputs
	}
	return nil
}
