package main

import "testing"

/*
Write tests that will execute rot13.
- Test with text as argument
  - Test with various inputs (table test, see t.Run)

- Test with text from standard input
*/
type testCase struct {
	name string
	in   string
	want string
}

func TestRot13Func(t *testing.T) {
	tests := []testCase{
		{"initial", "test", "grfg"},
		{"sub1", "Usenet", "Hfrarg"},
		{"sub2", "equivalent", "rdhvinyrag"},
		{"sub3", "printing", "cevagvat"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := rot13(test.in)

			if got != test.want {
				t.Fatalf("in=%s, want=%s, got=%s", test.in, test.want, got)
			}
		})
	}
}
