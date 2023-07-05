package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func crawl(url string) {
	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to crawl %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body of %s: %v\n", url, err)
		return
	}

	// Parse the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("Failed to parse HTML of %s: %v\n", url, err)
		return
	}

	// Find and process the desired elements in the document
	// Here's an example of printing all the <a> tags' href attribute values
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			fmt.Println(href)
		}
	})

	// Save the HTML content to a file
	fileName := "result.html"
	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		fmt.Printf("Failed to save HTML content to %s: %v\n", fileName, err)
		return
	}

	fmt.Printf("Crawling of %s completed. Results saved in %s\n", url, fileName)
}

func main() {
	url := "https://sequence.day/" // Replace with the desired URL
	crawl(url)
}
