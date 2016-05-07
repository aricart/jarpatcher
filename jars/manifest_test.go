package jars

import (
	"testing"
)

func TestStub(t *testing.T) {

	mf := Manifest{}
	mf.Parse("Hello: World")

	if mf.Map["Hello"] == "" {
		t.Fail()
	}

	if mf.Map["Hello"] != "World" {
		t.Fail()
	}

	mf.PrintHeaders()
}
