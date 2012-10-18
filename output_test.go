package main

import (
	"flag"
	. "launchpad.net/gocheck"
	"testing"
)

var output = flag.Bool("output", false, "Include output tests")

type OutputSuite struct{}

var _ = Suite(&OutputSuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *OutputSuite) SetUpSuite(c *C) {
	if !*output {
		c.Skip("-output not provided")
	}
}

func (s *OutputSuite) TestCheckFail(c *C) {
	c.Check(nil, NotNil)
}

func (s *OutputSuite) TestAssertFail(c *C) {
	c.Check("nope", Matches, "y.*ep")
	c.Assert(nil, NotNil, Commentf("sorry"))
}

func (s *OutputSuite) TestSkip(c *C) {
	c.Skip("because")
}

func (s *OutputSuite) TestPanic(c *C) {
	var nilstring string
	c.Check(nilstring[0], Equals, ' ')
}

func (s *OutputSuite) TestPass1(c *C) {
	c.Check(nil, IsNil)
}

func (s *OutputSuite) TestPass2(c *C) {
	c.Check(c, NotNil)
}

func (s *OutputSuite) TestPass3(c *C) {
	c.Check(c, Equals, c)
}
