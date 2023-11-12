package prototype

type (
	Prototype[T any] interface {
		Clone() Prototype[T]
		Prototype() Prototype[T]
		Underlying() *T
	}

	PrototypeSingletonFactory interface {
		Get(key any) Prototype[any]
	}
)
