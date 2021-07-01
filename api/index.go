package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`for/(\d{4}/\d{2}/\d{2})/`)

}

type data struct {
	Data map[string]map[string]int `json:"data"`
}


func Handler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	query := r.URL.Query() //Get URL Params(type map)
	user, ok := query["user"]
	if !ok || len(user) == 0 {
		http.Error(w, "no users", 400)
		return
	}
	log.Println(user)

	userData := scrapeMain(user[0])
	js, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, "internal error", 500)
		return
	}
	w.Write(js)

}

func scrapeMain(user string) data {
	store := map[string]map[string]int{}
	ch := make(chan string)
	go scrape("https://letterboxd.com/"+ user +"/films/diary/", ch)

	for elem := range ch {
		yearDate := strings.SplitAfterN(elem, "/", 2)
		if store[yearDate[0][:4]] == nil {
			store[yearDate[0][:4]] = make(map[string]int)
		}
		inner := store[yearDate[0][:4]]
		inner[yearDate[1]] += 1

	}
	return data{Data: store}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func scrape(url string, ch chan string) {
	wg := sync.WaitGroup{}
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1000})
	c.OnHTML("td.td-day.diary-day.center", func(e *colly.HTMLElement) {
		run := func() {
			wg.Add(1)
			dateUrl := e.ChildAttr("a[href]", "href")
			match := re.FindStringSubmatch(dateUrl)
			ch <- match[1]
			wg.Done()
		}
		go run()
	})

	c.OnHTML("div.paginate-nextprev", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr("a.next", "href")
		e.Request.Visit(e.Request.AbsoluteURL(nextPage))
	})

	c.Visit(url)
	wg.Wait()
	c.Wait()
	close(ch)
}
