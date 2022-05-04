package helper

import (
	"testing"
)

func TestInArray(t *testing.T) {
	type args struct {
		a interface{}
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContain(t *testing.T) {
	type args struct {
		searchkey interface{}
		target    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.searchkey, tt.args.target); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}
