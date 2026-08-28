package main

import (
	"os"
)

type outStream struct {
	stream *os.File
	path   string
}

func (o *outStream) String() string {
	defer o.stream.Close()

	// if o.stream is nil, default to os.Stdout
	if o.stream == nil {
		o.stream = os.Stdout
		o.path = "/dev/stdout"
	}
	return o.path
}

func (o *outStream) Set(s string) error {
	//Create a new file at path "s"
	out, err := os.Create(s)

	if err != nil {
		return err
	}

	o.path = s
	o.stream = out

	return nil
}
