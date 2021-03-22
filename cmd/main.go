package main

import (
	"fmt"
	"time"

	"crawl/clean"
	"crawl/crawler"
	"crawl/store"
)

const cookie = "__cfduid=d23ca80214b2bfdfca69e3882e171fc7b1616416043; __atssc=messenger%3B1; redux_update_check=3.6.18; wordpress_logged_in_3748aa90f9091fbd66dfda219c76b982=cherryupvietnam%40gmail.com%7C1616588970%7CjA4AzOKdrxVN4gWDq9p2ivE8VRL2cZMhxSV1npBr7sS%7Cd7457184cac3e187b670937c792ac6c4f501c0c3183630b0c6854bb8eeb085c3; __atuvc=8%7C12; __atuvs=60588d3b3888ed7f007"

const requestURL = "https://batdongsanchinhchu.vn/bds"
const maxPage = 60000
const startPage = 2100

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
