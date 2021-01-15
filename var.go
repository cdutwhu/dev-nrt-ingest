package main

import (
	"fmt"
	"strings"

	"github.com/gosuri/uiprogress"
)

var (
	fPln       = fmt.Println
	fPf        = fmt.Printf
	sTrimLeft  = strings.TrimLeft
	sTrimRight = strings.TrimRight
	sHasSuffix = strings.HasSuffix
	sContains  = strings.Contains

	dbgPln = func(do bool, a ...interface{}) (n int, err error) {
		if do {
			return fPln(a...)
		}
		return 0, nil
	}
)

const (
	attrPrefix   = ""
	contAttrName = "value"
	attrNameOfID = "RefId"
)

var (
	validate = false
	// CheckData : validate xml & json in parsing
	CheckData = func(check bool) {
		validate = check
	}
)

var (
	uip      *uiprogress.Progress
	bar      *uiprogress.Bar
	procsize uint64
	probar   bool
)
