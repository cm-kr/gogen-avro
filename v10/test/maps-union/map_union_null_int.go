// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
	"io"
)

func writeMapUnionNullInt(r map[string]*int32, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err := vm.WriteLong(int64(len(r)), w)
		if err != nil || len(r) == 0 {
			return err
		}
		if e == nil {
			err = vm.WriteLong(0, w)
			if err != nil {
				return err
			}
		} else {
			err = vm.WriteLong(int64(-1), w)
			if err != nil {
				return err
			}
			err = vm.WriteInt(*e, w)
		}
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapUnionNullIntWrapper struct {
	Target *map[string]*int32
	keys   []string
	values []*int32
}

func (_ *MapUnionNullIntWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapUnionNullIntWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapUnionNullIntWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]*int32, 0, s)
	}
}

func (r *MapUnionNullIntWrapper) NullField(_ int) {
	r.values[len(r.values)-1] = nil
}

func (r *MapUnionNullIntWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapUnionNullIntWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v int32
	r.values = append(r.values, &v)
	return &types.Int{Target: r.values[len(r.values)-1]}
}

func (_ *MapUnionNullIntWrapper) AppendArray() types.Field { panic("Unsupported operation") }
