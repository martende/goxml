package goxml_test

import (
	. "launchpad.net/gocheck"
	"testing"
)

type S struct{}

func TestAll(t *testing.T) { TestingT(t) }

var _ = Suite(&S{})
