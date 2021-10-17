package main

import "testing"

func TestRetirement(t *testing.T) {
	type parameters struct {
		currentAge int16
		retirementAge int16
	}

	tests := []struct {
		name string
		args parameters
		expected int16
}{
		{
			name: "test not yet retired",
			args: parameters{currentAge: 25, retirementAge: 65},
			expected: 40,
		},
		{
			name: "test already retired",
			args: parameters{currentAge: 66, retirementAge: 65},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := CalcRetirement(test.args.currentAge, test.args.retirementAge)
			if test.expected != actual {
				t.Errorf("CalcRetirement() = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestOutput(t *testing.T) {
	type parameters struct {
		yearsLeft int16
	}

	tests := []struct {
		name string
		args parameters
		expected string
	}{
		{
			name: "test not yet retired output",
			args: parameters{yearsLeft: 40},
			expected: "You have 40 years left until you can retire. It's 2021, so you can retire in 2061.",
		},
		{
			name: "test already retired output",
			args: parameters{yearsLeft: -1},
			expected: "You have no more years left to work. It's 2021, so you can already retire.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Output(test.args.yearsLeft)
			if test.expected != actual {
				t.Errorf("Output() = %v, want %v", actual, test.expected)
			}
		})
	}
}