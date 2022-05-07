package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	jobnum := 0
	fName := os.Args[1]
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		Delay:       0 * time.Second,
		RandomDelay: 0 * time.Second,
	})
	c.SetRequestTimeout(20 * time.Second)
	c.IgnoreRobotsTxt = true
	extensions.RandomUserAgent(c)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		i := 0
		e.ForEach("div[data-entity-urn]", func(_ int, el *colly.HTMLElement) {
			time := el.ChildText("time")
			id := el.Attr("data-entity-urn")
			name := el.ChildText("h3")
			company := el.ChildText("h4")
			location := el.ChildText("span.job-search-card__location")
			if len(id) != 0 && len(time) != 0 && len(company) != 0 && len(name) != 0 && len(location) != 0 {
				if strings.Contains(time, "minute") || time == "1 hour ago" {
					link := "https://www.linkedin.com/jobs/view/" + id[18:] + " "
					writer.Write([]string{
						link, name, company, location,
					})
					i++
				}
			}
		})
		fmt.Println("LinkedIn Scrapping Complete, scraped jobs:", i)
		jobnum += i
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("link1")
	c.Visit("link2")
	c.Visit("link3")
	c.Visit("link4")
	c.Visit("link5")
	c.Visit("link6")
	fmt.Println("Total jobs scraped :", jobnum)
}
