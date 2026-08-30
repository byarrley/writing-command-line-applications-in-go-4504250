package main

import (
	"os"
)

type flagOut struct {
	path string
}

func (f *flagOut) String() string {
	// if a.path is empty, default to os.Stdout
	if f.path == "" {
		f.path = os.Stdout.Name()
	}

	return f.path
}

func (f *flagOut) Set(s string) error {
	f.path = s

	return nil
}
