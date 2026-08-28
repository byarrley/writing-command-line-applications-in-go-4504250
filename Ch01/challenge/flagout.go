package main

import (
	"os"
)

type flagOut struct {
	stream *os.File
}

func (a *flagOut) String() string {
	//TODO: determine if there's a better place for this defer
	defer a.stream.Close()

	// if a.stream is nil, default to os.Stdout
	if a.stream == nil {
		a.stream = os.Stdout
	}
	return a.stream.Name()
}

func (a *flagOut) Set(s string) error {
	//Create a new file at path "s"
	out, err := os.Create(s)

	if err != nil {
		return err
	}

	a.stream = out
	return nil
}
