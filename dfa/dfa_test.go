package dfa

import "testing"

func TestNew(t *testing.T) {
	var (
		a = State("1")
		b = State("2")
		c = State("3")

		alphabetZero = Alphabet("0")
		alphabetOne  = Alphabet("1")
	)
	var states = []State{
		a, b, c,
	}

	var alphabets = []Alphabet{
		alphabetZero,
		alphabetOne,
	}

	var transition = map[State]Transition{
		a: Transition{
			alphabetZero: a,
			alphabetOne:  b,
		},
		b: Transition{
			alphabetZero: c,
			alphabetOne:  a,
		},
		c: Transition{
			alphabetZero: b,
			alphabetOne:  c,
		},
	}
	var initState = a
	var finalStats = []State{
		a,
	}

	dfa := New(states, alphabets, transition, initState, finalStats)

	tests := []struct {
		text string
		want bool
	}{
		{"00000", true},
		{"0011", true},
		{"01110", false},
		{"1001", true},
		{"1110101110", true},
	}

	for _, tt := range tests {
		if got := dfa.Run(tt.text); got != tt.want {
			t.Errorf("Error::: text %s, got = %v,  want = %v", tt.text, got, tt.want)
		}
	}

}
