// Code generated by thriftrw

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thrifttest

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type TestMapMapArgs struct {
	Hello *int32 `json:"hello,omitempty"`
}

func (v *TestMapMapArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Hello != nil {
		w, err = wire.NewValueI32(*(v.Hello)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TestMapMapArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.Hello = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestMapMapArgs) String() string {
	var fields [1]string
	i := 0
	if v.Hello != nil {
		fields[i] = fmt.Sprintf("Hello: %v", *(v.Hello))
		i++
	}
	return fmt.Sprintf("TestMapMapArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestMapMapArgs) MethodName() string {
	return "testMapMap"
}

func (v *TestMapMapArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestMapMapResult struct {
	Success map[int32]map[int32]int32 `json:"success"`
}

type _Map_I32_Map_I32_I32_MapItemList map[int32]map[int32]int32

func (m _Map_I32_Map_I32_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueI32(k), error(nil)
		if err != nil {
			return err
		}
		vw, err := wire.NewValueMap(_Map_I32_I32_MapItemList(v)), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_I32_Map_I32_I32_MapItemList) Size() int {
	return len(m)
}

func (_Map_I32_Map_I32_I32_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_I32_Map_I32_I32_MapItemList) ValueType() wire.Type {
	return wire.TMap
}

func (_Map_I32_Map_I32_I32_MapItemList) Close() {
}

func (v *TestMapMapResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_I32_Map_I32_I32_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestMapMapResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_I32_Map_I32_I32_Read(m wire.MapItemList) (map[int32]map[int32]int32, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}
	if m.ValueType() != wire.TMap {
		return nil, nil
	}
	o := make(map[int32]map[int32]int32, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI32(), error(nil)
		if err != nil {
			return err
		}
		v, err := _Map_I32_I32_Read(x.Value.GetMap())
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *TestMapMapResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_I32_Map_I32_I32_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestMapMapResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *TestMapMapResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("TestMapMapResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestMapMapResult) MethodName() string {
	return "testMapMap"
}

func (v *TestMapMapResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestMapMapHelper = struct {
	IsException    func(error) bool
	Args           func(hello *int32) *TestMapMapArgs
	WrapResponse   func(map[int32]map[int32]int32, error) (*TestMapMapResult, error)
	UnwrapResponse func(*TestMapMapResult) (map[int32]map[int32]int32, error)
}{}

func init() {
	TestMapMapHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	TestMapMapHelper.Args = func(hello *int32) *TestMapMapArgs {
		return &TestMapMapArgs{Hello: hello}
	}
	TestMapMapHelper.WrapResponse = func(success map[int32]map[int32]int32, err error) (*TestMapMapResult, error) {
		if err == nil {
			return &TestMapMapResult{Success: success}, nil
		}
		return nil, err
	}
	TestMapMapHelper.UnwrapResponse = func(result *TestMapMapResult) (success map[int32]map[int32]int32, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}