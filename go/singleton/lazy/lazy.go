package lazy

import (
	"fmt"
	"sync"

	"github.com/yeluyang/design_patterns/go/singleton"
)

var (
	instance      singleton.Singleton
	instanceMutex sync.Mutex
)

type singletonGetter struct{}

func New() singleton.SingletonGetter {
	return &singletonGetter{}
}

func (*singletonGetter) Get() singleton.Singleton {
	if instance == nil {
		instanceMutex.Lock()
		if instance == nil {
			instance = newSingletonInstance()
		}
		defer instanceMutex.Unlock()
	}
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
