package pkg

import (
	"bufio"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/avicrawler/db"
	"github.com/gofiber/fiber/v2/log"
)

type SafeMap struct {
	urls  map[string]string
	mutex sync.Mutex
}

type Url struct {
	url   string
	depth int
}

func CrawlUrl(url string, depth int) (nUrls int) {
	safeMap := SafeMap{
		urls:  make(map[string]string, 0),
		mutex: sync.Mutex{},
	}

	urlChan := make(chan Url)
	var wg sync.WaitGroup
	wg.Add(1)
	go crawl(url, depth, &safeMap, urlChan, &wg)

	go func() {
		wg.Wait()
		close(urlChan)
	}()

	i := 1
	for val := range urlChan {
		i += 1
		wg.Add(1)
		go crawl(val.url, val.depth, &safeMap, urlChan, &wg)
	}
	return i
}

func crawl(url string, depth int, urls *SafeMap, urlsChannel chan Url, wg *sync.WaitGroup) error {
	defer wg.Done()

	if depth <= 0 {
		return nil
	}

	urls.mutex.Lock()
	if _, exists := urls.urls[url]; exists {
		urls.mutex.Unlock()
		return nil
	}
	urls.urls[url] = ""
	urls.mutex.Unlock()

	content, err := fetchContent(url, depth, urlsChannel)

	if err != nil {
		log.Error(err)
		return err
	}

	db.SaveContent(url, content)

	return nil
}

var r = regexp.MustCompile(`https:\/\/[^"]*`)

func fetchContent(url string, depth int, urlsChannel chan Url) (content string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	var builder strings.Builder

	for {
		newLine, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		builder.WriteString(newLine)
		urls := r.FindAllString(newLine, -1)

		for _, val := range urls {
			urlsChannel <- Url{
				url:   val,
				depth: depth - 1,
			}
		}
	}

	return builder.String(), nil
}
