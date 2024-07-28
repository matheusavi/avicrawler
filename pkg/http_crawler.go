package pkg

import (
	"bufio"
	"net/http"
	"regexp"
	"strings"
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

	content, new_urls, err := fetchContent(url)

	if err != nil {
		return err
	}

	urls.mutex.Lock()
	urls.urls[url] = content
	urls.mutex.Unlock()

	for _, value := range new_urls {
		log.Info(value)
		crawl(value, depth-1, urls)
	}

	return nil
}

var r, _ = regexp.Compile("\"(https:\\/\\/[^\"]*)\"")

func fetchContent(url string) (content string, urls []string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", nil, err
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	result := make([]string, 0)

	var builder strings.Builder

	for scanner.Scan() {
		newLine := scanner.Text()
		builder.WriteString(newLine)
		urls := r.FindAllString(newLine, -1)

		if len(urls) > 0 {
			result = append(result, urls...) //Push this through a pointer?
		}
	}

	if err := scanner.Err(); err != nil {
		return "", nil, err
	}

	return builder.String(), result, nil
}
