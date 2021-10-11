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

// A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014
type HeaderworksDatatypeUUID struct {
	Uuid string `json:"uuid"`
}

const HeaderworksDatatypeUUIDAvroCRC64Fingerprint = "\xabK\x8cN\xf4L\xb6S"

func NewHeaderworksDatatypeUUID() HeaderworksDatatypeUUID {
	r := HeaderworksDatatypeUUID{}
	r.Uuid = ""
	return r
}

func DeserializeHeaderworksDatatypeUUID(r io.Reader) (HeaderworksDatatypeUUID, error) {
	t := NewHeaderworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeHeaderworksDatatypeUUIDFromSchema(r io.Reader, schema string) (HeaderworksDatatypeUUID, error) {
	t := NewHeaderworksDatatypeUUID()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeHeaderworksDatatypeUUID(r HeaderworksDatatypeUUID, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Uuid, w)
	if err != nil {
		return err
	}
	return err
}

func (r HeaderworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeHeaderworksDatatypeUUID(r, w)
}

func (r HeaderworksDatatypeUUID) Schema() string {
	return "{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"headerworks.datatype.UUID\",\"type\":\"record\"}"
}

func (r HeaderworksDatatypeUUID) SchemaName() string {
	return "headerworks.datatype.UUID"
}

func (_ HeaderworksDatatypeUUID) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetInt(v int32)       { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetLong(v int64)      { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetString(v string)   { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *HeaderworksDatatypeUUID) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Uuid}

		return w

	}
	panic("Unknown field index")
}

func (r *HeaderworksDatatypeUUID) SetDefault(i int) {
	switch i {
	case 0:
		r.Uuid = ""
		return
	}
	panic("Unknown field index")
}

func (r *HeaderworksDatatypeUUID) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ HeaderworksDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) HintSize(int)                     { panic("Unsupported operation") }
func (_ HeaderworksDatatypeUUID) Finalize()                        {}

func (_ HeaderworksDatatypeUUID) AvroCRC64Fingerprint() []byte {
	return []byte(HeaderworksDatatypeUUIDAvroCRC64Fingerprint)
}

func (r HeaderworksDatatypeUUID) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["uuid"], err = json.Marshal(r.Uuid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *HeaderworksDatatypeUUID) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["uuid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Uuid); err != nil {
			return err
		}
	} else {
		r.Uuid = ""
	}
	return nil
}
