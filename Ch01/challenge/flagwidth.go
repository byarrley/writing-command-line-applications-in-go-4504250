package main

import (
	"fmt"
	"strconv"
)

type flagWidth struct {
	width int
}

func (a *flagWidth) String() string {
	if a.width == 0 {
		a.width = 80
	}
	return strconv.Itoa(a.width)
}

func (a *flagWidth) Set(s string) error {
	w, err := strconv.Atoi(s)

	if err != nil {
		return err
	}

	if w > 0 && w < 250 {
		a.width = w
	} else {
		return fmt.Errorf("width must be an integer between 0 and 250, received %d", w)
	}

	return nil
}
