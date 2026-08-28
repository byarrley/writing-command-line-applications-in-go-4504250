package main

import (
	"os"
)

type outStream struct {
	stream *os.File
}

func (o *outStream) String() string {
	//TODO: determine if there's a better place for this defer
	defer o.stream.Close()

	// if o.stream is nil, default to os.Stdout
	if o.stream == nil {
		o.stream = os.Stdout
	}
	return o.stream.Name()
}

func (o *outStream) Set(s string) error {
	//Create a new file at path "s"
	out, err := os.Create(s)

	if err != nil {
		return err
	}

	o.stream = out
	return nil
}
