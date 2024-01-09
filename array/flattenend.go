// You can edit this code!
// Click here and start typing.
package main

import (
	"reflect"
	"testing"
)

func Flatten(m interface{}) []interface{} {
	arr := make([]interface{}, 0)

	switch v := m.(type) {
	case int, string, float32:
		arr = append(arr, v)
	case []interface{}:
		for _, value := range v {
			if value == nil {
				continue
			}

			list := Flatten(value)
			arr = append(arr, list...)
		}
	}

	return arr
}

var testCases = []struct {
	description string
	input       interface{}
	expected    []interface{}
}{
	{
		description: "empty",
		input:       []interface{}{},
		expected:    []interface{}{},
	},
	{
		description: "no nesting",
		input:       []interface{}{0, 1, 2},
		expected:    []interface{}{0, 1, 2},
	},
	{
		description: "flattens a nested array",
		input:       []interface{}{[]interface{}{[]interface{}{}}},
		expected:    []interface{}{},
	},
	{
		description: "flattens array with just integers present",
		input:       []interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8},
		expected:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "5 level nesting",
		input:       []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, 100, 4, []interface{}{[]interface{}{[]interface{}{50}}}}, -2},
		expected:    []interface{}{0, 2, 2, 3, 8, 100, 4, 50, -2},
	},
	{
		description: "6 level nesting",
		input:       []interface{}{1, []interface{}{2, []interface{}{[]interface{}{3}}, []interface{}{4, []interface{}{[]interface{}{5}}}, 6, 7}, 8},
		expected:    []interface{}{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		description: "null values are omitted from the final result",
		input:       []interface{}{1, 2, interface{}(nil)},
		expected:    []interface{}{1, 2},
	},
	{
		description: "consecutive null values at the front of the list are omitted from the final result",
		input:       []interface{}{interface{}(nil), interface{}(nil), 3},
		expected:    []interface{}{3},
	},
	{
		description: "consecutive null values in the middle of the list are omitted from the final result",
		input:       []interface{}{1, interface{}(nil), interface{}(nil), 4},
		expected:    []interface{}{1, 4},
	},
	{
		description: "6 level nest list with null values",
		input:       []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, []interface{}{[]interface{}{100}}, interface{}(nil), []interface{}{[]interface{}{interface{}(nil)}}}, -2},
		expected:    []interface{}{0, 2, 2, 3, 8, 100, -2},
	},
	{
		description: "all values in nested list are null",
		input:       []interface{}{interface{}(nil), []interface{}{[]interface{}{[]interface{}{interface{}(nil)}}}, interface{}(nil), interface{}(nil), []interface{}{[]interface{}{interface{}(nil), interface{}(nil)}, interface{}(nil)}, interface{}(nil)},
		expected:    []interface{}{},
	},
}

func TestFlatten(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Flatten(tc.input); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Flatten(%v) = %v, want: %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Flatten(tc.input)
		}
	}
}
