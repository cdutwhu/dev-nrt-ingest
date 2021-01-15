package main

import (
	"flag"
	"os"
	"time"

	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/nrt-temp/store"
	xt "github.com/cdutwhu/xml-tool"
)

func main() {
	defer misc.TrackTime(time.Now())

	xmlPathPtr := flag.String("input", "./rrd.xml", "path of input xml file")
	asyncPtr := flag.Bool("async", true, "async process")
	dataChkPtr := flag.Bool("check", false, "validate json and xml data")
	storeTypePtr := flag.String("store", "map", "store type [map, badger, file]")
	probarPtr := flag.Bool("bar", true, "show progress bar")
	flag.Parse()

	if !fileExists(*xmlPathPtr) {
		fPln("xml file is not exist, file path [-input] is needed")
		return
	}
	CheckData(*dataChkPtr)
	probar = *probarPtr

	// ------------------------------------- //

	var ingest IIngest
	var err error

	switch *storeTypePtr {
	case "map":
		ingest = store.NewSyncMap()

	case "badger":
		ingest, err = store.NewBadgerDB("./db")
		if err != nil {
			panic(err)
		}
		defer ingest.(*store.BadgerDB).Close()
		defer ingest.(*store.BadgerDB).Flush()

	case "file":
		ingest, err = store.NewLocalFile("./file/sif.json")
		if err != nil {
			panic(err)
		}
		defer ingest.(*store.LocalFile).FlushClose()

	default:
		fPln("[-store] is needed and from [map badger file]")
		return
	}

	// -------------------- //

	xt.SetSlim(true)
	xt.SetIgnrAttr(
		"xsi:nil",
		"xmlns:xsd",
		"xmlns:xsi",
		"xmlns",
	)
	xt.SetSuffix4List(`List`)
	xt.SetAttrPrefix(attrPrefix)
	xt.SetContAttrName(contAttrName)

	fPf("%d elements saved into [%s]\n", scan(*xmlPathPtr, true, *asyncPtr, ingest), *storeTypePtr)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
