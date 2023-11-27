package impl

import (
	"fmt"

	"github.com/yeluyang/design_patterns/go/bridge"
)

type (
	refinedAbstraction struct {
		a bridge.ImplementorA
		b bridge.ImplementorB
	}
	a struct{}
	b struct{}
)

func newA() bridge.ImplementorA { return &a{} }
func (*a) Supply() string       { return "a" }

func newB() bridge.ImplementorB { return &b{} }
func (*b) Supply() string       { return "b" }

func newRefinedAbstraction() bridge.Abstraction                     { return &refinedAbstraction{} }
func (r *refinedAbstraction) SetImplementorA(a bridge.ImplementorA) { r.a = a }
func (r *refinedAbstraction) SetImplementorB(b bridge.ImplementorB) { r.b = b }
func (r *refinedAbstraction) Opertaion() {
	fmt.Printf("refinedAbstraction:%s:%s\n", r.a.Supply(), r.b.Supply())
}
