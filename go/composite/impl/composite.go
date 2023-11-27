package impl

import (
	"fmt"

	"github.com/yeluyang/design_patterns/go/composite"
)

type (
	leaf struct {
		id string
	}

	container struct {
		id       string
		children []composite.Composite
	}
)

func newLeaf(id string) composite.Composite            { return &leaf{id} }
func (l *leaf) Child(composite.ID) composite.Composite { panic("unreachable") }
func (l *leaf) Add(composite.Composite)                { panic("unreachable") }
func (l *leaf) Remove(composite.ID)                    { panic("unreachable") }
func (l *leaf) ID() composite.ID                       { return composite.ID(l.id) }
func (l *leaf) Operation() {
	fmt.Printf("leaf[%s].operation\n", l.id)
}

func newContainer(id string) composite.Composite { return &container{id, []composite.Composite{}} }
func (l *container) ID() composite.ID            { return composite.ID(l.id) }
func (o *container) Add(c composite.Composite)   { o.children = append(o.children, c) }

func (o *container) Remove(id composite.ID) {
	for _, l := range o.children {
		if l.ID() == id {
			// remove this component
		}
	}
}

func (o *container) Child(id composite.ID) composite.Composite {
	for _, l := range o.children {
		if l.ID() == id {
			return l
		}
	}
	return nil
}

func (o *container) Operation() {
	for _, l := range o.children {
		l.Operation()
	}
	fmt.Printf("container[%s].operation\n", o.id)
}
