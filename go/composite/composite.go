package composite

type (
	ID        string
	Component interface {
		ID() ID
		Operation()
	}
	Composite interface {
		Component
		Add(c Composite)
		Child(id ID) Composite
		Remove(id ID)
	}
)
