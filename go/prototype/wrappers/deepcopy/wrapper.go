package deepcopy

import (
	"github.com/mohae/deepcopy"

	"github.com/yeluyang/design_patterns/go/prototype"
)

type PrototypeWrapper[T any] struct {
	inner     T
	prototype *PrototypeWrapper[T]
}

func NewWrapper[T any](inner T) prototype.Prototype[T] {
	return &PrototypeWrapper[T]{
		inner:     inner,
		prototype: nil,
	}
}

func (p *PrototypeWrapper[T]) Clone() prototype.Prototype[T] {
	w := &PrototypeWrapper[T]{
		prototype: p,
	}
	w.inner = deepcopy.Copy(p.inner).(T)
	return w
}

func (p *PrototypeWrapper[T]) Prototype() prototype.Prototype[T] {
	return p.prototype
}

func (p *PrototypeWrapper[T]) Underlying() *T {
	return &p.inner
}
