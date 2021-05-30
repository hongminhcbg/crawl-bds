package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

)

type UtilsClient struct {
	url string
	cookie string
}

func NewUtilsClient(url string, cookie string) *UtilsClient {
	return &UtilsClient{
		url:    url,
		cookie: cookie,
	}
}

func (c *UtilsClient) GetString(uri string) (string, error) {
	getPageURL := fmt.Sprintf("%s/%s", c.url, uri)
	request, err := http.NewRequest(http.MethodGet, getPageURL, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("cookie", c.cookie)
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