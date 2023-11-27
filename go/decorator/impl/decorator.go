package impl

import (
	"fmt"

	"github.com/yeluyang/design_patterns/go/decorator"
)

type (
	component struct{}

	decorateAround struct{ inner decorator.Component }
	decorateBefore struct{ inner decorator.Component }
	decorateAfter  struct{ inner decorator.Component }
)

func newComponent() decorator.Component { return &component{} }
func (c *component) Operation()         { fmt.Println("component") }

func newDecorateAround(inner decorator.Component) decorator.Component { return &decorateAround{inner} }
func (d *decorateAround) Operation() {
	fmt.Println("decorateAround.before")
	d.inner.Operation()
	fmt.Println("decorateAround.after")
}

func newDecorateBefore(inner decorator.Component) decorator.Component { return &decorateBefore{inner} }
func (d *decorateBefore) Operation() {
	fmt.Println("decorateBefore.before")
	d.inner.Operation()
}

func newDecorateAfter(inner decorator.Component) decorator.Component { return &decorateAfter{inner} }
func (d *decorateAfter) Operation() {
	d.inner.Operation()
	fmt.Println("decorateAfter.after")
}

func new() decorator.Component {
	return newDecorateAround(
		newDecorateBefore(
			newDecorateAfter(
				newComponent(),
			),
		),
	)
}
