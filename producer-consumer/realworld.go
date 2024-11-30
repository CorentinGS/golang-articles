package producer_consumer

import (
	"fmt"
	"time"
)

type Page struct {
	URL     string
	Content string
}

func scraper(urls []string, pages chan<- Page) {
	for _, url := range urls {
		// Simulate web scraping
		time.Sleep(100 * time.Millisecond)
		pages <- Page{
			URL:     url,
			Content: fmt.Sprintf("Content from %s", url),
		}
	}
	close(pages)
}

func processor(pages <-chan Page, results chan<- string) {
	for page := range pages {
		// Simulate content processing
		time.Sleep(200 * time.Millisecond)
		results <- fmt.Sprintf("Processed %s: %s", page.URL, page.Content)
	}
	close(results)
}

type RealWorldPattern struct{}

func (p RealWorldPattern) Execute() {
	urls := []string{"https://example1.com", "https://example2.com"}
	pages := make(chan Page)
	results := make(chan string)

	go scraper(urls, pages)
	go processor(pages, results)

	for result := range results {
		fmt.Println(result)
	}
}
