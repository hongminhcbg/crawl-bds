package main

import (
	"fmt"
	"time"

	"crawl/clean"
	"crawl/crawler"
	"crawl/store"
)

const cookie = "__cfduid=d47121bbf4a11649bb87c1c2897ebe3e11612182600; __sbmask=acqglygatyyalzijgctv@usqxgkanhrmoasbblkauw@4oXYd/I+Tcg67DWfaP8bWMWRbttC0LBlg++Lbg%3D%3D; redux_update_check=3.6.18; wordpress_logged_in_3748aa90f9091fbd66dfda219c76b982=mr.nvlam%40gmail.com%7C1612540418%7CDbkstd8guNkfTJm7JMf2phXXNnjMcVKiOWG9QPhsXxS%7C370b577e747843797de3f46a672194e87335bd101203caa120b1333b18f1bff3; __atuvc=32%7C5; __atuvs=601ac6c313ff0a7c002"

const requestURL = "https://batdongsanchinhchu.vn/bds"
const maxPage = 20000
const startPage = 1

func main() {

	queueRawData := make(chan string, 5000)
	queueSaveCSV := make(chan string, 5000)
	pages := make(chan int, 5000)

	currentPage := startPage
	go func(chan int) {
		for currentPage < maxPage {
			select {
			case pages <- currentPage:
				fmt.Printf("[INFOR] start get page[%d]\n", currentPage)
				currentPage++
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}(pages)

	filename := fmt.Sprintf("./results/%d_result.csv", time.Now().Unix())
	crawler := crawler.NewCrawler(requestURL, cookie)

	go crawler.Start(pages, queueRawData)
	go crawler.Start(pages, queueRawData)
	go crawler.Start(pages, queueRawData)

	go store.StoreToCSV(queueSaveCSV, filename)
	go clean.Consume(queueRawData, queueSaveCSV)

	defer close(queueSaveCSV)
	defer close(queueRawData)
	defer close(pages)

	for {
		time.Sleep(1 * time.Hour)
	}
}
