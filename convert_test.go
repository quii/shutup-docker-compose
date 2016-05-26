package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestItReplacesPortsWithExpose(t *testing.T) {
	in := `
e2e:
  ports:
    - "8040:8040"
  command: echo "ignored"
	`

	expected := `
e2e:
  expose:
    - "8040"
  command: echo "ignored"
	`

	out := &bytes.Buffer{}

	convert(strings.NewReader(in), out)

	if out.String() != expected+"\n" {
		t.Error("Port field was not converted into expose, expected", expected, "got", out.String())
	}
}
