package v2

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yeluyang/design_patterns/go/factory"
)

func TestFactory(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}

type FactoryTestSuite struct {
	suite.Suite
	factory factory.Factory
}

func (s *FactoryTestSuite) SetupSuite() {
	s.factory = new(factoryV2)
}

func (s *FactoryTestSuite) TestA() {
	s.Require().Equal("A.v2:", s.factory.A().A())
	s.Require().Equal("A.v2:default", s.factory.ADefault().A())
	s.Require().Equal("A.v2:start,params,end,", s.factory.A("start", "params", "end").A())
	s.Require().Equal("A.v2:from_config:path=/etc/config.toml", s.factory.AFromConfig("/etc/config.toml").A())
}

func (s *FactoryTestSuite) TestB() {
	s.Require().Equal("B.v2:", s.factory.B().B())
	s.Require().Equal("B.v2:default", s.factory.BDefault().B())
	s.Require().Equal("B.v2:start,params,end,", s.factory.B("start", "params", "end").B())
	s.Require().Equal("B.v2:from_config:path=/etc/config.toml", s.factory.BFromConfig("/etc/config.toml").B())
}
