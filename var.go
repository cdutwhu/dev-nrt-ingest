package main

import (
	"fmt"
	"strings"
)

var (
	fPln       = fmt.Println
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
	dataCheck = false
	// SetDataCheck : check xml & json in parsing
	SetDataCheck = func(check bool) {
		dataCheck = check
	}
)
