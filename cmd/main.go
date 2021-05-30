package main

import (
	"fmt"
	"time"

	"crawl/clean"
	"crawl/crawler"
	"crawl/store"
)


const cookie = "bat_dong_san_chinh_chu_session=eyJpdiI6IktZOVlQTU5VTUR3a2U4aWdBaEZTXC9RPT0iLCJ2YWx1ZSI6IkZLQm5zMndacTlxckp2XC95RzVPc3IxR3o1WnJTZHNGcTlxRVl1YUhjUFVqXC80MklmMzJUU2dsWFlubDNVS1hzSSIsIm1hYyI6IjgyMmM2MGUyM2QyYzdiOWZkNGUwNjk2ODgzZDQxMjI4ZGVjOTI3ZGI0ZTVjOWIyNWYwMDdlOTE4NmRlN2Y3ZDkifQ%3D%3D; expires=Sun, 30-May-2022 02:46:21 GMT; Max-Age=7200; path=/; httponly"
const requestURL = "https://batdongsanchinhchu.vn/product"
const maxPage = 60000
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
