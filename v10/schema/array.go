package schema

import (
	"fmt"
	"github.com/actgardner/gogen-avro/v10/generator"
)

type ArrayField struct {
	itemType   AvroType
	definition map[string]interface{}
}

func NewArrayField(itemType AvroType, definition map[string]interface{}) *ArrayField {
	return &ArrayField{
		itemType:   itemType,
		definition: definition,
	}
}

func (r *ArrayField) Name() string {
	return "Array" + r.itemType.Name()
}

func (r *ArrayField) filename() string {
	return generator.ToSnake(r.Name()) + ".go"
}

func (r *ArrayField) GoType() string {
	return fmt.Sprintf("[]%v", r.itemType.GoType())
}

func (r *ArrayField) SerializerMethod() string {
	return fmt.Sprintf("write%v", r.Name())
}

func (r *ArrayField) ItemType() AvroType {
	return r.itemType
}

func (r *ArrayField) Attribute(name string) interface{} {
	return r.definition[name]
}

func (r *ArrayField) Definition(scope map[QualifiedName]interface{}) (interface{}, error) {
	def := copyDefinition(r.definition)
	var err error
	def["items"], err = r.itemType.Definition(scope)
	if err != nil {
		return nil, err
	}
	return def, nil
}

func (r *ArrayField) ConstructorMethod() string {
	return fmt.Sprintf("make(%v, 0)", r.GoType())
}

func (r *ArrayField) DefaultValue(lvalue string, rvalue interface{}) (string, error) {
	items, ok := rvalue.([]interface{})
	if !ok {
		return "", fmt.Errorf("Expected array as default for %v, got %v", lvalue, rvalue)
	}

	setters := fmt.Sprintf("%v = make(%v,%v)\n", lvalue, r.GoType(), len(items))
	for i, item := range items {
		if c, ok := getConstructableForType(r.itemType); ok {
			setters += fmt.Sprintf("%v[%v] = %v\n", lvalue, i, c.ConstructorMethod())
		}

		setter, err := r.itemType.DefaultValue(fmt.Sprintf("%v[%v]", lvalue, i), item)
		if err != nil {
			return "", err
		}

		setters += setter + "\n"
	}
	return setters, nil
}

func (r *ArrayField) WrapperType() string {
	if union, ok := r.itemType.(*UnionField); ok {
		if union.IsSimpleNullUnion() && r.itemType.IsPrimitive() {
			return fmt.Sprintf("ArrayOfNullable%vUnion", union.itemType[union.NonNullIndex()].Name())
		}
	}

	return fmt.Sprintf("%vWrapper", r.Name())
}

func (r *ArrayField) WrapperPointer() bool {
	return false
}

func (r *ArrayField) IsReadableBy(f AvroType) bool {
	if union, ok := f.(*UnionField); ok {
		for _, t := range union.AvroTypes() {
			if r.IsReadableBy(t) {
				return true
			}
		}
	}

	if reader, ok := f.(*ArrayField); ok {
		return r.ItemType().IsReadableBy(reader.ItemType())
	}
	return false
}

func (r *ArrayField) ItemConstructable() string {
	if constructor, ok := getConstructableForType(r.itemType); ok {
		return fmt.Sprintf("v = %v\n", constructor.ConstructorMethod())
	}
	return ""
}

func (r *ArrayField) Children() []AvroType {
	return []AvroType{r.itemType}
}

func (r *ArrayField) UnionKey() string {
	return "array"
}

func (r *ArrayField) IsPrimitive() bool { return false }

func (r *ArrayField) IsSimpleNullUnion() bool {
	if unionField, ok := r.itemType.(*UnionField); !ok {
		return ok
	} else {
		return unionField.IsSimpleNullUnion()
	}
}

func (r *ArrayField) SimpleNullUnionNullIndex() int {
	if unionField, ok := r.itemType.(*UnionField); ok {
		return unionField.NullIndex()
	}
	return -1
}

func (r *ArrayField) SimpleNullUnionNonNullIndex() int {
	if unionField, ok := r.itemType.(*UnionField); ok {
		return unionField.NonNullIndex()
	}
	return -1
}

func (r *ArrayField) SimpleNullUnionItemType() string {
	if unionField, ok := r.itemType.(*UnionField); ok {
		return unionField.itemType[unionField.NonNullIndex()].Name()
	}
	return ""
}
