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

type NestedRecord struct {
	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField Bytes `json:"BytesField"`
}

const NestedRecordAvroCRC64Fingerprint = "\x81\x8d\xc3K?\xe83\xcc"

func NewNestedRecord() NestedRecord {
	r := NestedRecord{}
	return r
}

func DeserializeNestedRecord(r io.Reader) (NestedRecord, error) {
	t := NewNestedRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNestedRecordFromSchema(r io.Reader, schema string) (NestedRecord, error) {
	t := NewNestedRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNestedRecord(r NestedRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r NestedRecord) Serialize(w io.Writer) error {
	return writeNestedRecord(r, w)
}

func (r NestedRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}"
}

func (r NestedRecord) SchemaName() string {
	return "NestedRecord"
}

func (_ NestedRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NestedRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NestedRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NestedRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NestedRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NestedRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NestedRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ NestedRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.StringField}

		return w

	case 1:
		w := types.Boolean{Target: &r.BoolField}

		return w

	case 2:
		w := BytesWrapper{Target: &r.BytesField}

		return w

	}
	panic("Unknown field index")
}

func (r *NestedRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NestedRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NestedRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NestedRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NestedRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NestedRecord) Finalize()                        {}

func (_ NestedRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NestedRecordAvroCRC64Fingerprint)
}

func (r NestedRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["StringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["BoolField"], err = json.Marshal(r.BoolField)
	if err != nil {
		return nil, err
	}
	output["BytesField"], err = json.Marshal(r.BytesField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NestedRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["StringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StringField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BoolField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BoolField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BoolField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BytesField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BytesField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BytesField")
	}
	return nil
}
