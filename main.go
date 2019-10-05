package main

import (
	"fmt"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/agent"
	"log"
	"oregon-law-templating/clients"
	"oregon-law-templating/ors"
	"oregon-law-templating/ors/volume"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Task struct {
	close chan struct{}
	wg     sync.WaitGroup
	Vol 	ors.Volume
	ticker *time.Ticker
}

//type ActiveStatues struct {
//	Vols [...]Volume
//}

func (t *Task) get() {

}

func updateLegalDocuments() {

}



var logger = log.New(os.Stdout, "Main > ", log.LstdFlags)

func main() {
	goSurf()
	//goSurfLiz()

	store := clients.Store{}
	store.InitializeViperEndpoints()
	urls := store.Vi.Get("urls")

	logger.Println(urls)
	 ApiBase := (urls.SODAApi).string

	vol := ors.Volume{}
	vol.Init()

	task := &Task{
		close: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 2),
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.Run() }()

	select {
	case sig := <-c:
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		task.Stop()
	}
}

func goSurfLiz() {
	bow := surf.NewBrowser()

	err := bow.Open("https://api.oregonlegislature.gov/odata/odataservice.svc/")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fmt.Println(bow.Title())

	for _, link := range bow.Links() {

		logger.Println(link.Text)
	}
}

type SurfNode int

const (
	Last SurfNode  = iota + 1
	First
	Previous
	Current
	Next
	Node
)

type LinkNode struct {
	Url string
	NextNodes map[string]SurfNode
	LastNode SurfNode
}

type MissionMap struct {
	entryPoint LinkNode
	priorityNodes []LinkNode
	history []string
	URLS string
	path string
	current LinkNode
}

func goSurf() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.AOL())
	//err := bow.Open()
	err := bow.Open("https://www.oregonlaws.org/oregon_revised_statutes")
	if err != nil {
		panic(err)
	}

	logger.Println(bow.Body())

	// Outputs: "reddit: the front page of the internet"
	fmt.Println(bow.Title())

	for _, link := range bow.Links() {
		time.Sleep(time.Second * 3)

		//fmt.Println(link.Text)
		//fmt.Println(link.URL.String())
		for i := 0; i < len(link.URL.String()); i++ {

			volMatched := false
			volumeTest := "tattoo"
			test := link.URL.String()
			for j := 0; j < len(volumeTest); j++ {

				if string([]rune(test)[i + j]) != string([]rune(volumeTest)[j]) {
					//fmt.Println(string([]rune(test)[i + j]) + " - does not match - " + string([]rune(volumeTest)[j]))
					break
				}
				fmt.Println(string([]rune(test)[i + j]) + " - does match - " + string([]rune(volumeTest)[j]))
				println("matched: \n" + link.Text)
				volMatched = true

			}

			if volMatched {

				err := bow.Open(link.URL.String())
				if err != nil {
					println(err)
				}


				for _, links2 := range bow.Links() {

					fmt.Println(links2.Text)
					fmt.Println(links2.URL.String())
				}

			}
		}
	}

	fmt.Println(ors.One.Title())

	// Click the link for the newest submissions.
	//bow.Click("a.new")

	// Outputs: "newest submissions: reddit.com"
	//fmt.Println(bow.Title())

	// Log in to the site.
	//fm, _ := bow.Form("form.login-form")
	//fm.Input("user", "JoeRedditor")
	//fm.Input("passwd", "d234rlkasd")
	//if fm.Submit() != nil {
	//	panic(err)
	//}

	// Go back to the "newest submissions" page, bookmark it, and
	// print the title of every link on the page.
	//bow.Back()
	//bow.Bookmark("reddit-new")
	//bow.Find("a.title").Each(func(_ int, s *goquery.Selection) {
	//	fmt.Println(s.Text())
	//})
}

func (t *Task) Run() {
	for {
		select {
		case <-t.close:
			return
		case <-t.ticker.C:
			handle()
		}
	}
}

func (t *Task) Stop() {
	close(t.close)
	t.wg.Wait()
}

func handle() {
	for i := 0; i < 5; i++ {
		fmt.Print("#")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println()
}
