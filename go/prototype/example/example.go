package example

import (
	"github.com/yeluyang/design_patterns/go/prototype"
	"github.com/yeluyang/design_patterns/go/prototype/wrappers/deepcopy"
)

var prototypes map[any]prototype.Prototype[prototypeObject]

type (
	keyTypeEmpty   struct{}
	keyTypeDefault struct{}
)

var (
	KeyEmpty   = keyTypeEmpty{}
	KeyDefault = keyTypeDefault{}
)

func init() {
	prototypes = map[any]prototype.Prototype[prototypeObject]{
		KeyEmpty: deepcopy.NewWrapper[prototypeObject](prototypeObject{}),
		KeyDefault: deepcopy.NewWrapper[prototypeObject](prototypeObject{
			Sclar:    "sclar",
			Compound: []byte("compund"),
			Structure: struct {
				Sclar    string
				Compound []byte
			}{
				Sclar:    "structure.sclar",
				Compound: []byte("structure.compund"),
			},
		}),
	}
}

type prototypeObject struct {
	Sclar     string
	Compound  []byte
	Structure struct {
		Sclar    string
		Compound []byte
	}
}

type PrototypeFactory struct{}

func (*PrototypeFactory) Get(key any) (prototype.Prototype[prototypeObject], bool) {
	p, ok := prototypes[key]
	return p, ok
}
