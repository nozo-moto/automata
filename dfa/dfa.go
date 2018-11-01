package dfa

import "fmt"

type State string

type Alphabet string

func (a *Alphabet) String() string {
	return string(*a)
}

func newAlphabet(word string) Alphabet {
	var a Alphabet
	a = Alphabet(word)
	return a
}

type Transition map[Alphabet]State

type DeterministicFiniteAutomata struct {
	States      []State
	Alphabets   []Alphabet
	Transitions map[State]Transition
	InitState   State
	FinalStates []State
}

func New(states []State, alphabets []Alphabet, transitions map[State]Transition, initState State, finalStates []State) *DeterministicFiniteAutomata {
	return &DeterministicFiniteAutomata{
		States:      states,
		Alphabets:   alphabets,
		Transitions: transitionsi,
		InitState:   initState,
		FinalStates: finalStates,
	}
}

func (d *DeterministicFiniteAutomata) Run(text string) bool {
	var state = d.InitState
	for _, t := range text {
		if d.containAlphabet(t) {
			return false
		} else {
			state = d.Transitions[state][newAlphabet(string(t))]
		}
	}

	return d.containFinalState(state)
}

func (d *DeterministicFiniteAutomata) containAlphabet(word rune) bool {
	for _, v := range d.Alphabets {
		if v.String() == fmt.Sprint(word) {
			return true
		}
	}

	return false
}

func (d *DeterministicFiniteAutomata) containFinalState(state State) bool {
	for _, v := range d.FinalStates {
		if v == state {
			return true
		}
	}

	return false
}
