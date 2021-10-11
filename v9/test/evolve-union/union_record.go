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

type UnionRecord struct {
	A string `json:"a"`

	Id *UnionNullInt `json:"id"`

	Name *UnionNullString `json:"name"`
}

const UnionRecordAvroCRC64Fingerprint = "\xfeS\x1bd\xa1\xfc͒"

func NewUnionRecord() UnionRecord {
	r := UnionRecord{}
	r.Id = NewUnionNullInt()

	r.Id = nil
	r.Name = NewUnionNullString()

	r.Name = nil
	return r
}

func DeserializeUnionRecord(r io.Reader) (UnionRecord, error) {
	t := NewUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUnionRecordFromSchema(r io.Reader, schema string) (UnionRecord, error) {
	t := NewUnionRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUnionRecord(r UnionRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.A, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r UnionRecord) Serialize(w io.Writer) error {
	return writeUnionRecord(r, w)
}

func (r UnionRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":\"string\"},{\"default\":null,\"name\":\"id\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"name\",\"type\":[\"null\",\"string\"]}],\"name\":\"UnionRecord\",\"type\":\"record\"}"
}

func (r UnionRecord) SchemaName() string {
	return "UnionRecord"
}

func (_ UnionRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ UnionRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ UnionRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ UnionRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ UnionRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ UnionRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ UnionRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ UnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.A}

		return w

	case 1:
		r.Id = NewUnionNullInt()

		return r.Id
	case 2:
		r.Name = NewUnionNullString()

		return r.Name
	}
	panic("Unknown field index")
}

func (r *UnionRecord) SetDefault(i int) {
	switch i {
	case 1:
		r.Id = nil
		return
	case 2:
		r.Name = nil
		return
	}
	panic("Unknown field index")
}

func (r *UnionRecord) NullField(i int) {
	switch i {
	case 1:
		r.Id = nil
		return
	case 2:
		r.Name = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ UnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ UnionRecord) Finalize()                        {}

func (_ UnionRecord) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecordAvroCRC64Fingerprint)
}

func (r UnionRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["a"], err = json.Marshal(r.A)
	if err != nil {
		return nil, err
	}
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *UnionRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["a"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.A); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for a")
	}
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullInt()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		r.Name = NewUnionNullString()

		r.Name = nil
	}
	return nil
}
