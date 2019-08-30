package ors

import (
	"fmt"
	"html/template"
	"net/http"
)

type VolumeOrder int

const TitlesAndChaptersPdfUrl = "https://www.oregonlegislature.gov/bills_laws/BillsLawsEDL/2013ORS_TitlesChapters.pdf"

const (
	One 	VolumeOrder = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
	Twelve
	Thirteen
	Fourteen
	Fifteen
	Sixteen
	Seventeen
)

type LawSection struct {
	head template.HTML
	body template.HTML
	foot template.HTML
}

type LawDocument struct {
	Order int
	Title string
	Header template.HTML
	Sections []LawSection
}

type Chapter struct {
	Order int
	Title string
	Content LawDocument
}

type ChapterSection struct {
	Order int
	Title string
	Chapters map[int]Chapter
}

type Volume struct {
	Order int
	Title string
	ChaptersSections map[int]ChapterSection
}

var VolumeTitles = map[int]string{ 0: "Title 1", 1: "Title 2" }

func (v VolumeOrder) Title() string {
	fmt.Println(v)
	return VolumeTitles[int(v)]
}

type options struct {
	IsWhole bool
	Meta map[string]string
	Payload map[string]string
}

func (v *Volume) Write(options options) {

}

func (v *Volume) Init() {
	titleChaps, err := http.Get(TitlesAndChaptersPdfUrl)
	if err != nil {
		configsLogger.Println("Error", err)
	}
	configsLogger.Println("Title PDF?", titleChaps)
}
