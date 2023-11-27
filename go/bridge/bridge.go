package bridge

type (
	ImplementorA interface{ Supply() string }
	ImplementorB interface{ Supply() string }

	Abstraction interface {
		SetImplementorA(a ImplementorA)
		SetImplementorB(b ImplementorB)

		Opertaion()
	}
)
