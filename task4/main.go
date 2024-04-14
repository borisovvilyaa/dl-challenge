package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var (
	cache = make(map[string][]string)
	mu    sync.Mutex
)

// fetchArticleLinks fetches links from a given URL
func fetchArticleLinks(url string) ([]string, error) {
	mu.Lock()
	defer mu.Unlock()

	// Check if links are cached
	if links, ok := cache[url]; ok {
		return links, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page: %s", resp.Status)
	}

	re := regexp.MustCompile(`<a\s+(?:[^>]*?\s+)?href="([^"]*)"[^>]*>(.*?)</a>`)

	links := make([]string, 0)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			matches := re.FindAllStringSubmatch(string(buf[:n]), -1)
			for _, match := range matches {
				link := match[1]
				if regexp.MustCompile(`^/wiki/[^:]+$`).MatchString(link) {
					links = append(links, "https://en.wikipedia.org"+link)
				}
			}
		}
		if err != nil {
			break
		}
	}

	// Cache the links
	cache[url] = links

	return links, nil
}

// findArticle checks if a specific article is present in the links retrieved from a given URL
func findArticle(url string, articleTitle string) (bool, string, error) {
	links, err := fetchArticleLinks(url)
	if err != nil {
		return false, "", err
	}

	for _, link := range links {
		if link == "https://en.wikipedia.org/wiki/"+articleTitle {
			return true, url, nil
		}
	}

	return false, "", nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var url string
	fmt.Print("Enter the Wikipedia article URL: ")
	fmt.Scanln(&url)

	found := false
	foundURL := ""

	// Try to find the article "Adolf_Hitler" in the Wikipedia links, up to 6 iterations
	for i := 0; i < 6; i++ {
		found, foundURL, _ = findArticle(url, "Adolf_Hitler")
		if found {
			break
		}

		links, err := fetchArticleLinks(url)
		if err != nil {
			log.Fatal("Error fetching links:", err)
		}

		randomIndex := rand.Intn(len(links))
		url = links[randomIndex]
	}

	if found {
		fmt.Println(foundURL)
	} else {
		fmt.Println("Hitler not found.")
	}
}
