package main

import (
	"io"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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
		{"simple", "test", "grfg"},
		{"all_caps", "USENET", "HFRARG"},
		{"special_chars", "</>", "</>"},
		{"numerals", "12345", "12345"},
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

func TestRot13Exe(t *testing.T) {
	tests := []testCase{
		{"simple", "test", "grfg"},
		{"all_caps", "USENET", "HFRARG"},
		{"special_chars", "</>", "</>"},
		{"numerals", "12345", "12345"},
	}

	exe := buildExe(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := exec.Command(exe, test.in).CombinedOutput()

			require.NoError(t, err, "run:\n%s", string(got))

			if strings.TrimSpace(string(got)) != test.want {
				t.Errorf("in=%s, want=%s, got=%s", test.in, test.want, got)
			}
		})
	}
}

func TestRot13Stdin(t *testing.T) {
	//Adapted from example at https://pkg.go.dev/os/exec#example-Cmd.StdinPipe
	tests := []testCase{
		{"simple", "test", "grfg"},
		{"all_caps", "USENET", "HFRARG"},
		{"special_chars", "</>", "</>"},
		{"numerals", "12345", "12345"},
	}

	exe := buildExe(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command(exe)      //set up cmd
			stdin, err := cmd.StdinPipe() //prep stdin

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, test.in) //write input to stdin in goroutine
			}()

			got, err := cmd.CombinedOutput() //run the command, output is merged stdout/stderr (2>&1)
			require.NoError(t, err, "run:\n%s", string(got))

			if strings.TrimSpace(string(got)) != test.want {
				t.Errorf("in=%s, want=%s, got=%s", test.in, test.want, got)
			}
		})
	}
}

func buildExe(t *testing.T) string {
	exe := path.Join(t.TempDir(), "rot13")
	output, err := exec.Command("go", "build", "-o", exe).CombinedOutput()
	require.NoErrorf(t, err, "build:\n%s", string(output))

	return exe
}
