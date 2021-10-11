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

type ComCompanySharedTypeTwo struct {
	Type ComCompanySharedSomeEnum `json:"type"`
}

const ComCompanySharedTypeTwoAvroCRC64Fingerprint = "rG\xf7\xb2\x8f\xac|\x98"

func NewComCompanySharedTypeTwo() ComCompanySharedTypeTwo {
	r := ComCompanySharedTypeTwo{}
	return r
}

func DeserializeComCompanySharedTypeTwo(r io.Reader) (ComCompanySharedTypeTwo, error) {
	t := NewComCompanySharedTypeTwo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComCompanySharedTypeTwoFromSchema(r io.Reader, schema string) (ComCompanySharedTypeTwo, error) {
	t := NewComCompanySharedTypeTwo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComCompanySharedTypeTwo(r ComCompanySharedTypeTwo, w io.Writer) error {
	var err error
	err = writeComCompanySharedSomeEnum(r.Type, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComCompanySharedTypeTwo) Serialize(w io.Writer) error {
	return writeComCompanySharedTypeTwo(r, w)
}

func (r ComCompanySharedTypeTwo) Schema() string {
	return "{\"fields\":[{\"name\":\"type\",\"type\":{\"name\":\"SomeEnum\",\"symbols\":[\"X\",\"Y\"],\"type\":\"enum\"}}],\"name\":\"com.company.shared.TypeTwo\",\"type\":\"record\"}"
}

func (r ComCompanySharedTypeTwo) SchemaName() string {
	return "com.company.shared.TypeTwo"
}

func (_ ComCompanySharedTypeTwo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetString(v string)   { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComCompanySharedTypeTwo) Get(i int) types.Field {
	switch i {
	case 0:
		w := ComCompanySharedSomeEnumWrapper{Target: &r.Type}

		return w

	}
	panic("Unknown field index")
}

func (r *ComCompanySharedTypeTwo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ComCompanySharedTypeTwo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ComCompanySharedTypeTwo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComCompanySharedTypeTwo) Finalize()                        {}

func (_ ComCompanySharedTypeTwo) AvroCRC64Fingerprint() []byte {
	return []byte(ComCompanySharedTypeTwoAvroCRC64Fingerprint)
}

func (r ComCompanySharedTypeTwo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["type"], err = json.Marshal(r.Type)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComCompanySharedTypeTwo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["type"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Type); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for type")
	}
	return nil
}
