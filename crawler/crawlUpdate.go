package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type crawlUpdateImpl struct {
	pageFormat string
	baseURL    string
}

func NewCrawlerUpdate(pageFormat, baseURL string) Crawler {
	return &crawlUpdateImpl{
		pageFormat: pageFormat,
		baseURL:    baseURL,
	}
}

func (c *crawlUpdateImpl) doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status resp %d not OK", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *crawlUpdateImpl) getPage(pageIndex int) ([]byte, error) {
	pageURL := fmt.Sprintf(c.pageFormat, pageIndex)
	return c.doGetRequest(pageURL)
}

func (c *crawlUpdateImpl) getInforHref(rawPage []byte) []string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(rawPage))
	if err != nil {
		return nil
	}

	results := make([]string, 0)

	doc.Find(".product-item.clearfix").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Find("a").Nodes {
			for _, att := range node.Attr {
				if att.Key == "href" {
					results = append(results, att.Val)
				}
			}
		}
	})

	return results
}

func (c *crawlUpdateImpl) getSingleInfor(inforURL string) ([]byte, error) {
	getURL := c.baseURL + inforURL
	return c.doGetRequest(getURL)
}

func (c *crawlUpdateImpl) Start(pageChan chan int, rawResults chan string) {
	for pageIndex := range pageChan {
		pageRaw, err := c.getPage(pageIndex)
		if err != nil {
			log.Printf("[ERROR] pageID = %d, error := %s\n", pageIndex, err.Error())
			continue
		}

		log.Printf("[INFOR] get page success pageID = %d\n", pageIndex)
		inforIDs := c.getInforHref(pageRaw)

		for _, inforID := range inforIDs {
			inforRaw, err := c.getSingleInfor(inforID)
			if err != nil {
				log.Printf("[ERROR] infor href = %s, error := %s\n", inforID, err.Error())
				continue
			}

			log.Printf("[INFOR] GET SUCCESS infor href = %s\n", inforID)
			rawResults <- string(inforRaw)
		}
	}
}
