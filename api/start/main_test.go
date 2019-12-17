package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	got, _ := Handler(nil, &Request{})
	if got.StatusCode != 200 {
		t.Fail()
	}
}