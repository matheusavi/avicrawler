package pkg

import (
	"bufio"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

type SafeMap struct {
	urls  map[string]string
	mutex sync.Mutex
}

func CrawlUrl(url string, depth int) {
	safeMap := SafeMap{
		urls:  make(map[string]string, 0),
		mutex: sync.Mutex{},
	}
	crawl(url, depth, &safeMap)
}

func crawl(url string, depth int, urls *SafeMap) error {
	if depth <= 0 {
		return nil
	}

	urls.mutex.Lock()
	if _, exists := urls.urls[url]; exists {
		return nil
	}
	urls.mutex.Unlock()

	content, err := fetchContent(url)

	if err != nil {
		return err
	}

	log.Info(content)

	urls.mutex.Lock()
	urls.urls[url] = content
	urls.mutex.Unlock()

	new_urls := getUrlsFromContent(content)

	for _, value := range new_urls {
		crawl(value, depth-1, urls)
	}

	return nil
}

func fetchContent(url string) (content string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		content += scanner.Text() //TODO: Improve string concat performance
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func getUrlsFromContent(content string) []string {
	result := make([]string, 0)
	return result
}
