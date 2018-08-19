package main

import (
	"fmt"
	"testing"
)

func Test_fun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		name: "dalizi",
	}
	for tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
		})
	}
}
