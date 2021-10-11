// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     superhero.avsc
 */
package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type Superpower struct {
	Id int32 `json:"id"`

	Name string `json:"name"`

	Damage float32 `json:"damage"`

	Energy float32 `json:"energy"`

	Passive bool `json:"passive"`
}

const SuperpowerAvroCRC64Fingerprint = "\r+\x84\x8ee\xc6a\n"

func NewSuperpower() Superpower {
	r := Superpower{}
	return r
}

func DeserializeSuperpower(r io.Reader) (Superpower, error) {
	t := NewSuperpower()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSuperpowerFromSchema(r io.Reader, schema string) (Superpower, error) {
	t := NewSuperpower()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSuperpower(r Superpower, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Damage, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Energy, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Passive, w)
	if err != nil {
		return err
	}
	return err
}

func (r Superpower) Serialize(w io.Writer) error {
	return writeSuperpower(r, w)
}

func (r Superpower) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"damage\",\"type\":\"float\"},{\"name\":\"energy\",\"type\":\"float\"},{\"name\":\"passive\",\"type\":\"boolean\"}],\"name\":\"com.model.Superpower\",\"type\":\"record\"}"
}

func (r Superpower) SchemaName() string {
	return "com.model.Superpower"
}

func (_ Superpower) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Superpower) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Superpower) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Superpower) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Superpower) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Superpower) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Superpower) SetString(v string)   { panic("Unsupported operation") }
func (_ Superpower) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Superpower) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Name}

		return w

	case 2:
		w := types.Float{Target: &r.Damage}

		return w

	case 3:
		w := types.Float{Target: &r.Energy}

		return w

	case 4:
		w := types.Boolean{Target: &r.Passive}

		return w

	}
	panic("Unknown field index")
}

func (r *Superpower) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Superpower) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Superpower) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Superpower) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Superpower) HintSize(int)                     { panic("Unsupported operation") }
func (_ Superpower) Finalize()                        {}

func (_ Superpower) AvroCRC64Fingerprint() []byte {
	return []byte(SuperpowerAvroCRC64Fingerprint)
}

func (r Superpower) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["damage"], err = json.Marshal(r.Damage)
	if err != nil {
		return nil, err
	}
	output["energy"], err = json.Marshal(r.Energy)
	if err != nil {
		return nil, err
	}
	output["passive"], err = json.Marshal(r.Passive)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Superpower) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		return fmt.Errorf("no value specified for id")
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
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["damage"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Damage); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for damage")
	}
	val = func() json.RawMessage {
		if v, ok := fields["energy"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Energy); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for energy")
	}
	val = func() json.RawMessage {
		if v, ok := fields["passive"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Passive); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for passive")
	}
	return nil
}
