package main

import (
	"time"

	"github.com/Knetic/govaluate"
)

type Default struct {
	Name  string
	Value interface{}
	//~ Type  string
}

type Check struct {
	Desc    string
	If      *govaluate.EvaluableExpression
	Classes []string
}

type Probe struct {
	Name      string
	Script    string
	Targets   []string
	Delay     time.Duration
	Timeout   time.Duration
	Arguments string
	Defaults  []*Default
	Checks    []*Check
}
