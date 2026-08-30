package main

import (
	"fmt"
	"strconv"
)

type flagWidth struct {
	width int
}

func (f *flagWidth) String() string {
	if f.width == 0 {
		f.width = 80
	}
	return strconv.Itoa(f.width)
}

func (f *flagWidth) Set(s string) error {
	w, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	if w > 0 && w < 250 {
		f.width = w
	} else {
		return fmt.Errorf("width must be an integer between 0 and 250, received %d", w)
	}

	return nil
}
