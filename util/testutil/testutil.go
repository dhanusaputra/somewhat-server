package testutil

import "reflect"

type ptrStruct struct {
	ptr         interface{}
	oriFunc     interface{}
	oriFuncType reflect.Type
}
type ptrs struct {
	ptrs []ptrStruct
}

// TearDownInterface ...
type TearDownInterface interface {
	Restore()
}

// NewPtrs ...
func NewPtrs(ins []interface{}) TearDownInterface {
	outs := make([]ptrStruct, 0, len(ins))
	for _, in := range ins {
		outs = append(outs, ptrStruct{
			ptr:         in,
			oriFunc:     reflect.ValueOf(in).Elem().Interface(),
			oriFuncType: reflect.ValueOf(in).Elem().Type(),
		})
	}
	return &ptrs{ptrs: outs}
}

// Restore ...
func (p *ptrs) Restore() {
	for _, ptrStruct := range p.ptrs {
		if ptrStruct.oriFunc == nil {
			reflect.ValueOf(ptrStruct.ptr).Elem().Set(reflect.Zero(ptrStruct.oriFuncType))
		} else {
			reflect.ValueOf(ptrStruct.ptr).Elem().Set(reflect.ValueOf(ptrStruct.oriFunc))
		}
	}
}
