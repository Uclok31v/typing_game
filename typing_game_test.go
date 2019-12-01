package main

import (
	"bytes"
	"io"
	"testing"
)

func Test_main(t *testing.T){
	cases := []struct{
		name string
		r io.Reader
	}
}

func Test_input(t *testing.T) {
	cases := []struct {
		name     string
		r        io.Reader
		expected string
	}{
		{name: "true", r: bytes.NewBufferString("survey"), expected: "survey"},
	}

	for _, c := range cases {
		t.Helper()
		t.Run(c.name, func(t *testing.T) {
			ch := input(c.r)
			actual := <-ch
			if c.expected != actual {
				t.Errorf("want %s, actual %s", c.expected, actual)
			}
		})
	}
}
