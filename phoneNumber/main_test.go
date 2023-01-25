package main

import "testing"

type TestNumbers struct {
	input string
	want  string
}

func TestNormalize(t *testing.T) {
	tests := []TestNumbers{
		{"01234567890", "01234567890"},
		{"01234567890", "01234567890"},
		{"(012)34567890", "01234567890"},
		{"012-3456-7890", "01234567890"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := normalize(test.input)
			if got != test.want {
				t.Errorf("test error, input: %s, want: %s, got: %s ", test.input, test.want, got)
			}
		})
	}

}
