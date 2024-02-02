package main

import "testing"

func Test_findClosestValue(t *testing.T) {
	type args struct {
		values []float32
		value  float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "when values length is 0",
		},
		{
			name: "when values length is 1",
		},
		{
			name: "when values length > 2",
		},
		{
			name: "when values array have duplicate values",
		},
		{
			name: "when values array have both negative & positive values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestValue(tt.args.values, tt.args.value); got != tt.want {
				t.Errorf("findClosestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
