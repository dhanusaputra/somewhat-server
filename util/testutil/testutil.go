package testutil

import "reflect"

type ptr struct {
	ptr         interface{}
	oriFunc     interface{}
	oriFuncType reflect.Type
}
type ptrs struct {
	ptrs []*ptr
}

// TearDownInterface ...
type TearDownInterface interface {
	Restore()
}

// NewPtrs ...
func NewPtrs(ins []interface{}) TearDownInterface {
	outs := make([]*ptr, 0, len(ins))
	for _, in := range ins {
		outs = append(outs, &ptr{
			ptr:         in,
			oriFunc:     reflect.ValueOf(in).Elem().Interface(),
			oriFuncType: reflect.ValueOf(in).Elem().Type(),
		})
	}
	return &ptrs{ptrs: outs}
}

// Restore ...
func (p *ptrs) Restore() {
	for _, pair := range p.ptrs {
		if pair.oriFunc == nil {
			reflect.ValueOf(pair.ptr).Elem().Set(reflect.Zero(pair.oriFuncType))
		} else {
			reflect.ValueOf(pair.ptr).Elem().Set(reflect.ValueOf(pair.oriFunc))
		}
	}
}
