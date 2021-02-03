package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type crawImpl struct {
	URL    string
	Cookie string
}

type Crawler interface {
	Start(pageChan chan int, queue chan string)
}

func NewCrawler(url, cookie string) Crawler {
	return &crawImpl{
		URL:    url,
		Cookie: cookie,
	}
}

func (c *crawImpl) getPageData(pageIndex int) (string, error) {
	var getPageURL string
	if pageIndex <= 1 {
		getPageURL = c.URL
	} else {
		getPageURL = c.URL + fmt.Sprintf("/page/%d", pageIndex)
	}

	request, err := http.NewRequest(http.MethodGet, getPageURL, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("cookie", c.Cookie)
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	httpClient := http.Client{
		Timeout: 80 * time.Second,
	}
	resp, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status code is %d, not OK", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *crawImpl) Start(pageChan chan int, queue chan string) {
	for pageID := range pageChan {
		rawData, err := c.getPageData(pageID)
		if err != nil {
			fmt.Printf("[CRAWL][ERROR] get page %d error %s\n", pageID, err.Error())
			continue
		}

		fmt.Printf("[CRAWL][INFOR] get page %d success\n", pageID)
		queue <- rawData
	}
}
