package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {

	/*
	 This is a snippet of code used to familiarize myself with the colly web scraping framework
	*/

	var ratings = []string{}
	// Set the domain for the colly web scraping framework
	web_scraper := colly.NewCollector(
		colly.AllowedDomains("https://openlibrary.org", "openlibrary.org"),
	)
	// Scrape the reader ratings
	web_scraper.OnHTML(".readers-stats ", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, elem *colly.HTMLElement) {
			ratings = append(ratings, e.DOM.Find("span").Text())
		})
	})
	// Scrape the book plot
	var plot string = ""
	web_scraper.OnHTML(".book-description-content", func(e *colly.HTMLElement) {
		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			plot = e.Text
		})
	})
	// Rate limit the colly web scraping requests
	web_scraper.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	web_scraper.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting...", request.URL.String())
	})
	// Scrape the book data for Klara and the Sun
	web_scraper.Visit("https://openlibrary.org/works/OL20883297W/")

	// Parse and format the book ratings data
	rating := strings.ReplaceAll(ratings[len(ratings)-1], "★", "")
	rating_float, _ := strconv.ParseFloat(rating, 64)
	rating_float = math.Ceil(rating_float*100) / 100
	rating = strconv.FormatFloat(rating_float, 'E', -1, 64)
	rating = strings.ReplaceAll(rating, "E+00", "")

	fmt.Println("Book Ratings : " + rating)
	fmt.Println("Book Plot : " + plot)
}
