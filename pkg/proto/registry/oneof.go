package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

type Oneof interface {
	Protobuf() *protobuf.Oneof

	GetFieldByName(name string) (*OneofField, bool)

	GetFieldByLine(line int) (*OneofField, bool)
}

type oneof struct {
	protoOneofField *protobuf.Oneof

	fieldNameToField map[string]*OneofField

	lineToField map[int]*OneofField

	mu *sync.RWMutex
}

var _ Oneof = (*oneof)(nil)

func NewOneof(protoOneofField *protobuf.Oneof) Oneof {
	oneof := &oneof{
		protoOneofField: protoOneofField,

		fieldNameToField: make(map[string]*OneofField),

		lineToField: make(map[int]*OneofField),
	}

	for _, e := range protoOneofField.Elements {
		v, ok := e.(*protobuf.OneOfField)
		if !ok {
			continue
		}
		f := NewOneofField(v)
		oneof.fieldNameToField[v.Name] = f
		oneof.lineToField[v.Position.Line] = f
	}

	return oneof
}

func (o *oneof) Protobuf() *protobuf.Oneof {
	return o.protoOneofField
}

// GetFieldByName gets EnumField  by provided name.
// This ensures thread safety.
func (o *oneof) GetFieldByName(name string) (*OneofField, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	f, ok := o.fieldNameToField[name]
	return f, ok
}

// GetFieldByName gets MapField by provided line.
// This ensures thread safety.
func (o *oneof) GetFieldByLine(line int) (*OneofField, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	f, ok := o.lineToField[line]
	return f, ok
}

type OneofField struct {
	ProtoOneOfField *protobuf.OneOfField
}

func NewOneofField(protoOneOfField *protobuf.OneOfField) *OneofField {
	return &OneofField{
		ProtoOneOfField: protoOneOfField,
	}
}
