package main

import (
	"regexp"
	"testing"
)

func Test_hello_with_name(t *testing.T) {
	name := "MyName"
	want := regexp.MustCompile(`\b` + name + `\b`)
	n, err := Hello(name)
	if !want.MatchString(n) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, n, err, want)
	}
}

func Test_hello(t *testing.T) {
	n, err := Hello("")
	if n != "" && err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, n, err)
	}
}
