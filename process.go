package main

import (
	"sync"

	xt "github.com/cdutwhu/xml-tool"
)

// var ops uint64
var mutex = &sync.Mutex{}

func proc(params ...interface{}) error {

	var (
		wg       *sync.WaitGroup
		id       = params[1].(int64)   // threadID
		xml      = params[2].(string)  // input xml
		cvt2json = params[3].(bool)    // whether to convert to json
		ingest   = params[4].(IIngest) // ingest interface
	)

	if params[0] != nil {
		wg = params[0].(*sync.WaitGroup) // WaitGroup
	}

	if wg != nil {
		defer func() {
			mutex.Lock()
			wg.Done()
			mutex.Unlock()
		}()
	}

	// atomic.AddUint64(&ops, 1)
	dbgPln(false, "---@---", id)

	if wg != nil {
		mutex.Lock()
		wg.Add(1)
		mutex.Unlock()
	}

	if cvt2json {
		parse4json(xml, ingest)
	} else {
		parse(xt.Fmt(xml), ingest) // if store xml, Fmt it. if store json, do not need fmt
	}
	return nil
}
