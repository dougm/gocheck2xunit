# gocheck2xunit

Converts `go test -gocheck.vv` output to XML for use by the Jenkins [xunit
plugin](https://wiki.jenkins-ci.org/display/JENKINS/xUnit+Plugin).

Similar to [go2xunit](https://bitbucket.org/tebeka/go2xunit), but parses output
from [gocheck](http://labix.org/gocheck)

## Install

	go get github.com/dougm/gocheck2xunit

## Usage

	go test -gocheck.vv | gocheck2xunit > tests.xml
