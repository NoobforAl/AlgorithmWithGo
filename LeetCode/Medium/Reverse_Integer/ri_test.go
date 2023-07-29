package main

import (
	"bytes"
	"testing"

	"github.com/kr/pretty"
)

func TestCase1(t *testing.T) {
	var err bytes.Buffer
	pretty.Fdiff(&err, reverse(-12345), -54321)
	if err.String() != "" {
		t.Fatal(err)
	}
}

func TestCase2(t *testing.T) {
	var err bytes.Buffer
	pretty.Fdiff(&err, reverse(55213), 31255)
	if err.String() != "" {
		t.Fatal(err)
	}
}
