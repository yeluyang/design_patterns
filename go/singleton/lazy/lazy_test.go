package lazy

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yeluyang/design_patterns/go/singleton"
)

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite

	getter singleton.SingletonGetter
}

func (s *TestSuite) SetupSuite() {
	s.getter = New()
}

func (s *TestSuite) Test() {
	s.Require().Same(s.getter.Get(), s.getter.Get())
}
