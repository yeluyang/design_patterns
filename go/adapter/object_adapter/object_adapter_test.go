package objectadapter

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yeluyang/design_patterns/go/adapter"
)

func TestFooAdapter(t *testing.T) {
	suite.Run(t, new(fooAdapterTestSuite))
}

type fooAdapterTestSuite struct {
	suite.Suite
	adapter adapter.Foo
}

func (s *fooAdapterTestSuite) SetupSuite() {
	s.adapter = newFooAdapter(&barAdaptee{}, &buzzAdaptee{})
}

func (s *fooAdapterTestSuite) TestBarAdapter() {
	r, err := s.adapter.FooBar("foo")
	s.Require().NoError(err)
	s.Require().Equal("bar:foo", r)
}

func (s *fooAdapterTestSuite) TestBuzzAdapter() {
	r, err := s.adapter.FooBuzz("foo")
	s.Require().NoError(err)
	s.Require().Equal("buzz:foo", r)
}
