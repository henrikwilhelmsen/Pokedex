package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"lots of space": {
			input: "   hello    world  ",
			want:  []string{"hello", "world"}},
		"one word, no space": {
			input: "foo",
			want:  []string{"foo"}},
		"variable casing": {
			input: "  PICHU pikachu charizaRD",
			want:  []string{"pichu", "pikachu", "charizard"}},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(testCase.input)
			msg := fmt.Sprintf(
				"Result does not match expected output, expected: %#v, got: %#v",
				testCase.want,
				got,
			)
			if len(got) != len(testCase.want) {
				t.Fatal(msg)
			}
			for i := range got {
				actualWord := got[i]
				expectedWord := testCase.want[i]
				if actualWord != expectedWord {
					t.Fatal(msg)
				}
			}
		})
	}
}
