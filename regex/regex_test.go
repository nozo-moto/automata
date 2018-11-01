package regex

import (
	"testing"
)

func Test(t *testing.T) {
	var (
		regex1 = "a*"
	)
	tests := []struct {
		Token   TokenType
		Literal string
	}{
		{
			Token:   IDENT,
			Literal: "a",
		},
		{
			Token:   ASTERISK,
			Literal: "*",
		},
	}

	input := New(regex1)
	for _, tt := range tests {
		token := input.NextToken()
		if token.Literal != tt.Literal {
			t.Errorf("expected token is %v, got %v", tt.Literal, token.Literal)
		}
		if token.Literal != tt.Literal {
			t.Errorf("expected literal is %v, got %v", tt.Literal, token.Literal)
		}
	}
}
