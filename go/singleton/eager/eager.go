package eager

import (
	"fmt"

	"github.com/yeluyang/design_patterns/go/singleton"
)

func init() {
	instance = newSingletonInstance()
}

var instance singleton.Singleton

type singletonGetter struct{}

func New() singleton.SingletonGetter {
	return &singletonGetter{}
}

func (*singletonGetter) Get() singleton.Singleton {
	return instance
}

type singletonInstance struct {
	content string
}

func newSingletonInstance() singleton.Singleton {
	return &singletonInstance{content: "singletonInstance"}
}

func (i *singletonInstance) Do() {
	fmt.Println(i.content)
}
