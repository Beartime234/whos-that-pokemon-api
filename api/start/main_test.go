package main

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	got, _ := Handler(nil, &Request{})
	log.Print(got)
}