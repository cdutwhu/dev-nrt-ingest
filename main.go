package main

import (
	"time"

	"github.com/cdutwhu/gotil/misc"
	"github.com/cdutwhu/nrt-temp/store"
	xt "github.com/cdutwhu/xml-tool"
)

func main() {
	defer misc.TrackTime(time.Now())

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

	// ------------------------------------- //

	// ingest := store.NewSyncMap()

	// ---------------- //

	// ingest, err := store.NewBadgerDB("./db")
	// if err != nil {
	// 	panic(err)
	// }
	// defer ingest.Close()
	// defer ingest.Flush()

	// ---------------- //

	ingest, err := store.NewLocalFile("./file/sif.json")
	if err != nil {
		panic(err)
	}
	defer ingest.FlushClose()

	// ------------------------------------- //

	scan("./sif.xml", true, true, ingest)
}
