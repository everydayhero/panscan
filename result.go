package main

import (
	"fmt"
)

type Result struct {
	Source string
	Loc    string
	Value  string
}

func (r Result) String() string {
	return fmt.Sprintf("[%s:%s] %s", r.Source, r.Loc, r.Value)
}
