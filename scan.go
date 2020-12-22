package main

import (
	"bufio"
	"os"
	"sync"

	xt "github.com/cdutwhu/xml-tool"
)

// scan :
func scan(xmlpath string, cvt2json, async bool, ingest IIngest) {
	file, err := os.Open(xmlpath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	br := bufio.NewReader(file)
	count := int64(0)
	var dataTypes = []string{
		"NAPStudentResponseSet",
		"NAPEventStudentLink",
		"StudentPersonal",
		"NAPTestlet",
		"NAPTestItem",
		"NAPTest",
		"NAPCodeFrame",
		"SchoolInfo",
		"NAPTestScoreSummary",
	}

	if async {
		var wg sync.WaitGroup
		for ele := range xt.StreamEle(br, dataTypes...) {
			count++
			go proc(&wg, count, ele, cvt2json, ingest)
		}
		wg.Wait()
	} else {
		for ele := range xt.StreamEle(br) {
			count++
			proc(nil, count, ele, cvt2json, ingest)
		}
	}

	fPln("total", count)
}
