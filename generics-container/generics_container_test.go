package main

import (
	"testing"
)

func TestContainer_Put(t *testing.T) {
	type args struct {
		elem interface{}
	}
	tests := []struct {
		name string
		c    *Container
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Put(tt.args.elem)
		})
	}
}

func Test_main(t *testing.T) {

	t.Run("Main", func(t *testing.T) {
		main()
	})
}
