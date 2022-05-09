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
	jobnum := 0 // This will be our total jobs scraped counter
	fName := os.Args[1] // reading the file name given, should be data.csv if you use my script
	// I don't check whether it's data.csv or not just in case you want to use your own script to expand 
	file, err := os.Create(fName) 
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close() // This is so handy coming from C

	writer := csv.NewWriter(file) 
	defer writer.Flush()
	c := colly.NewCollector() // Creating a new collector
	// Now this part is only here in case I want to add delays later, for now they work with 0 delay, if you ever get 
	// "Too many requests" error message, try adding a delay (1 second for example)
	c.Limit(&colly.LimitRule{
		Delay:       0 * time.Second, 
		RandomDelay: 0 * time.Second,
	})
	c.SetRequestTimeout(20 * time.Second) // Almost never happens if you have internet 
	c.IgnoreRobotsTxt = true
	extensions.RandomUserAgent(c)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		i := 0
		// Now for this part, the collector looks for a div that has the data-entity-urn attribute
		// When it finds it, it scrapes the contents of its children
		e.ForEach("div[data-entity-urn]", func(_ int, el *colly.HTMLElement) { 
			time := el.ChildText("time") // not used really but might come in handy later
			id := el.Attr("data-entity-urn") // The id of the job listing, used for the link
			name := el.ChildText("h3") // A more appropriate name should be title but eh
			company := el.ChildText("h4")
			location := el.ChildText("span.job-search-card__location")
			// Sometimes it detects empty lines (I have no idea why), so this if statement saves us from the panic that insues
			// due to `id[18:]
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
	// Check the picture in README.md in order to know what to do here
	c.Visit("link1")
	c.Visit("link2")
	c.Visit("link3")
	c.Visit("link4")
	c.Visit("link5")
	c.Visit("link6")
	fmt.Println("Total jobs scraped :", jobnum)
}
