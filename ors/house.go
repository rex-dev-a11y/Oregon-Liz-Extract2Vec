package ors

import (
	"log"
	"os"
	"time"
)

var configsLogger = log.New(os.Stdout, "Configs > ", log.LstdFlags)

type HouseBill struct {
	URL				 ORS_URL
	Code			 string
	Title			 string
	RawText			 string
	Sections		 []string
	Fetched			 bool
	Idle			 bool
	UpdatedLast	     time.Time
	HashSum			 []uint32
}

type ORS_URL int

type ORS_URL_OPTIONS struct {
	Base string
	Path string
}

const (
	sodaAPI  = "https://data.oregon.gov/resource/sitn-tqhd.json"
	baseUrl = "https://olis.leg.state.or.us/"
	downloadMeasureSegment =  "/Downloads/MeasureDocument/"
)

//func (o ORS_URL) toEncoded(options ORS_URL_OPTIONS) url.URL {
//	return url.URL{ Host: options.Base, Path: options.Path }
//}

type Legislature int

const (
	L2015  			Legislature  = iota + 1
	L2017
	L2019
)

var legislatures = []string{"2019R1", "2017R1", "2017R1"}

func (l Legislature) toString() string {
	return legislatures[l]
}

func (o ORS_URL) build(options map[string]string) string {

	return "https://olis.leg.state.or.us/liz/" + options["liz"] + "/Downloads/MeasureDocument/" + options["code"] + "/Introduced"
}

func (hb *HouseBill) fetch(code int) {
	//brow := surf.NewBrowser()
	//opts := ORS_URL_OPTIONS{baseUrl, hb}
	//url := hb.URL.build(opts)
	//brow.Open(url)
	//body := brow.Body
	//configsLogger.Println(body)
}

