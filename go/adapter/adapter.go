package adapter

type (
	Foo interface {
		FooBar(string) (string, error)
		FooBuzz(string) (string, error)
	}
)
