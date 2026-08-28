package main

import (
	"os"
)

type argOut struct {
	stream *os.File
}

func (o *argOut) String() string {
	//TODO: determine if there's a better place for this defer
	defer o.stream.Close()

	// if o.stream is nil, default to os.Stdout
	if o.stream == nil {
		o.stream = os.Stdout
	}
	return o.stream.Name()
}

func (o *argOut) Set(s string) error {
	//Create a new file at path "s"
	out, err := os.Create(s)

	if err != nil {
		return err
	}

	o.stream = out
	return nil
}
