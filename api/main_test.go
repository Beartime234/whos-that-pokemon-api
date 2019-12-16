package main

import "testing"

func Test_testMe(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name:"TestTest", want:"Hi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testMe(); got != tt.want {
				t.Errorf("testMe() = %v, want %v", got, tt.want)
			}
		})
	}
}