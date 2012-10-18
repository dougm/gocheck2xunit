package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type reason struct {
	Message string `xml:",chardata"`
}

type testcase struct {
	Classname string  `xml:"classname,attr"`
	Name      string  `xml:"name,attr"`
	Time      string  `xml:"time,attr,omitempty"`
	Failure   *reason `xml:"failure,omitempty"`
	Error     *reason `xml:"error,omitempty"`
	Skipped   *reason `xml:"skipped,omitempty"`
}

type testsuite struct {
	Name     string      `xml:"name,attr"`
	Tests    int         `xml:"tests,attr"`
	Failures int         `xml:"failures,attr"`
	Errors   int         `xml:"errors,attr"`
	Skipped  int         `xml:"skipped,attr"`
	TestCase []*testcase `xml:"testcase"`
}

func (suite *testsuite) start(line string) *testcase {
	name := strings.Split(line[strings.LastIndex(line, " ")+1:], ".")
	test := &testcase{Classname: name[0], Name: name[1]}
	suite.TestCase = append(suite.TestCase, test)
	suite.Tests++
	return test
}

func (suite *testsuite) read(in io.Reader) {
	var test *testcase
	output := new(bytes.Buffer)
	reader := bufio.NewReader(in)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		switch {
		case strings.HasPrefix(line, "START:"):
			output.Reset()
			test = suite.start(line[0 : len(line)-1])
		case strings.HasPrefix(line, "PASS:"):
		case strings.HasPrefix(line, "FAIL:"):
			suite.Failures++
			test.Failure = &reason{Message: output.String()}
		case strings.HasPrefix(line, "PANIC:"):
			suite.Errors++
			test.Error = &reason{Message: output.String()}
		case strings.HasPrefix(line, "SKIP:"):
			suite.Skipped++
			test.Skipped = &reason{}
		default:
			output.Write([]byte(line))
		}
	}

}

func (suite *testsuite) write(out io.Writer) {
	data, err := xml.MarshalIndent(suite, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(out, xml.Header, string(data), "\n")
}

func main() {
	suite := &testsuite{Name: "gocheck"}
	suite.read(os.Stdin)
	suite.write(os.Stdout)
	os.Exit(suite.Failures)
}
