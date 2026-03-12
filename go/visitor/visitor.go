package visitor

// Element 接口
type Element interface {
	Accept(v Visitor)
}

// Visitor 接口
type Visitor interface {
	VisitA(a *ConcreteElementA)
	VisitB(b *ConcreteElementB)
}

// ConcreteElementA 具体元素A
type ConcreteElementA struct {
	Value string
}

func (a *ConcreteElementA) Accept(v Visitor) {
	v.VisitA(a)
}

// ConcreteElementB 具体元素B
type ConcreteElementB struct {
	Value int
}

func (b *ConcreteElementB) Accept(v Visitor) {
	v.VisitB(b)
}
