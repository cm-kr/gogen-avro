// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type Record1 struct {
	Intfield int32 `json:"intfield"`
}

const Record1AvroCRC64Fingerprint = "5\x1f\v\x05O7Z\x02"

func NewRecord1() Record1 {
	r := Record1{}
	return r
}

func DeserializeRecord1(r io.Reader) (Record1, error) {
	t := NewRecord1()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecord1FromSchema(r io.Reader, schema string) (Record1, error) {
	t := NewRecord1()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecord1(r Record1, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Intfield, w)
	if err != nil {
		return err
	}
	return err
}

func (r Record1) Serialize(w io.Writer) error {
	return writeRecord1(r, w)
}

func (r Record1) Schema() string {
	return "{\"fields\":[{\"name\":\"intfield\",\"type\":\"int\"}],\"name\":\"record1\",\"type\":\"record\"}"
}

func (r Record1) SchemaName() string {
	return "record1"
}

func (_ Record1) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Record1) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Record1) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Record1) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Record1) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Record1) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Record1) SetString(v string)   { panic("Unsupported operation") }
func (_ Record1) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Record1) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Intfield}

		return w

	}
	panic("Unknown field index")
}

func (r *Record1) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Record1) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Record1) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Record1) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Record1) HintSize(int)                     { panic("Unsupported operation") }
func (_ Record1) Finalize()                        {}

func (_ Record1) AvroCRC64Fingerprint() []byte {
	return []byte(Record1AvroCRC64Fingerprint)
}

func (r Record1) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["intfield"], err = json.Marshal(r.Intfield)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Record1) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["intfield"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Intfield); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for intfield")
	}
	return nil
}
