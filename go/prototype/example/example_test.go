package example_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yeluyang/design_patterns/go/prototype/example"
)

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	factory *example.PrototypeFactory
}

func (s *TestSuite) SetupSuite() {
	s.factory = &example.PrototypeFactory{}
}

func (s *TestSuite) Test() {
	prototypeRoot, ok := s.factory.Get(example.KeyEmpty)
	s.Require().True(ok)

	objA := prototypeRoot.Clone()
	s.Require().Equal(prototypeRoot.Underlying(), objA.Underlying())
	s.Require().NotSame(prototypeRoot.Underlying(), objA.Underlying())
	s.Require().Same(prototypeRoot, objA.Prototype())
	s.Require().Nil(objA.Prototype().Prototype())

	objB := prototypeRoot.Clone()
	s.Require().Equal(objA, objB)
	s.Require().NotSame(objA, objB)

	prototypeRoot.Underlying().Sclar = "sclar"
	prototypeRoot.Underlying().Compound = []byte("compound")
	prototypeRoot.Underlying().Structure.Sclar = "structure.sclar"
	prototypeRoot.Underlying().Structure.Compound = []byte("structure.compound")
	s.Require().NotEqual(prototypeRoot.Underlying(), objA.Underlying())
	s.Require().NotEqual(prototypeRoot.Underlying(), objB.Underlying())

	objA = prototypeRoot.Clone()
	s.Require().Equal(prototypeRoot.Underlying(), objA.Underlying())
	s.Require().NotSame(prototypeRoot.Underlying(), objA.Underlying())
	s.Require().Same(prototypeRoot, objA.Prototype())
	s.Require().Nil(objA.Prototype().Prototype())

	objB = prototypeRoot.Clone()
	s.Require().Equal(objA, objB)
	s.Require().NotSame(objA, objB)

	objAA := objA.Clone()
	s.Require().Equal(objA.Underlying(), objAA.Underlying())
	s.Require().NotSame(objA.Underlying(), objAA.Underlying())
	s.Require().Same(objA, objAA.Prototype())
	s.Require().Same(prototypeRoot, objAA.Prototype().Prototype())
	s.Require().Nil(objAA.Prototype().Prototype().Prototype())

	objAB := objA.Clone()
	s.Require().Equal(objAA, objAB)
	s.Require().NotSame(objAA, objAB)
}
