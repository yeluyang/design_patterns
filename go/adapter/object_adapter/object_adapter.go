package objectadapter

import (
	"bytes"

	"github.com/yeluyang/design_patterns/go/adapter"
)

type barAdaptee struct{}

func (b *barAdaptee) Bar(param []byte) ([]byte, error) {
	buf := bytes.NewBufferString("bar:")
	buf.Write(param)
	return buf.Bytes(), nil
}

type buzzAdaptee struct{}

func (b *buzzAdaptee) Buzz(param []byte) ([]byte, error) {
	buf := bytes.NewBufferString("buzz:")
	buf.Write(param)
	return buf.Bytes(), nil
}

type barAdapter struct {
	adaptee *barAdaptee
}

func newBarAdapter(bar *barAdaptee) *barAdapter {
	return &barAdapter{bar}
}

func (b *barAdapter) FooBar(param string) (string, error) {
	r, err := b.adaptee.Bar([]byte(param))
	if err != nil {
		return "", err
	}
	return string(r), nil
}

type buzzAdapter struct {
	adaptee *buzzAdaptee
}

func newBuzzAdapter(buzz *buzzAdaptee) *buzzAdapter {
	return &buzzAdapter{buzz}
}

func (b *buzzAdapter) FooBuzz(param string) (string, error) {
	r, err := b.adaptee.Buzz([]byte(param))
	if err != nil {
		return "", err
	}
	return string(r), nil
}

type fooAdapter struct {
	*barAdapter
	*buzzAdapter
}

func newFooAdapter(bar *barAdaptee, buzz *buzzAdaptee) adapter.Foo {
	return &fooAdapter{barAdapter: newBarAdapter(bar), buzzAdapter: newBuzzAdapter(buzz)}
}
