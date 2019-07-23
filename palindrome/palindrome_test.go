package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name  string
		input int64
		want  bool
	}{
		{name: "returns false if input is negative", input: -676766, want: false},
		{name: "input is even digits and symmetrical", input: 1111, want: true},
		{name: "input is odd digits and symmetrical", input: 11111, want: true},
		{name: "input is even digits and non-symmetrical", input: 222333, want: false},
		{name: "input is odd digits and non-symmetrical", input: 2223335, want: false},
		{name: "another odd symmetry", input: 7776777, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
