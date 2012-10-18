package main

import (
	"bytes"
	. "launchpad.net/gocheck"
	"os/exec"
)

type MainSuite struct{}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestOutput(c *C) {
	if *output {
		return
	}
	args := []string{"test", "-gocheck.vv", "-output", "OutputSuite"}
	cmd := exec.Command("go", args...)
	output, _ := cmd.CombinedOutput()

	rd := bytes.NewReader(output)
	suite := &testsuite{Name: "main"}
	suite.read(rd)

	c.Check(suite.Failures, Equals, 2)
	c.Check(suite.Errors, Equals, 1)
	c.Check(suite.Skipped, Equals, 1)
	c.Check(suite.Tests, Equals, 9)

	for _, test := range suite.TestCase {
		if test.Skipped != nil {
			c.Check(test.Classname, Equals, "OutputSuite")
			c.Check(test.Name, Equals, "TestSkip")
		}
	}
}
