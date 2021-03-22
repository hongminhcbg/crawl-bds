package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"crawl/clean"
	"crawl/crawler"
	"crawl/store"
)

const requestURL = "https://batdongsanchinhchu.vn/bds"
const configURL = "/config.json"
const maxPage = 10000
const startPage = 1

type Config struct {
	Cookie    string `json:"cookie"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	config, err := getConfig(dir + configURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[INFOR] start with config %+v\n", config)
	queueRawData := make(chan string, 50)
	queueSaveCSV := make(chan string, 50)
	pages := make(chan int, 50)
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

	filename := fmt.Sprintf("%s/results/%d_result.csv", dir, time.Now().Unix())
	crawler := crawler.NewCrawler(requestURL, config.Cookie)
	fmt.Println(config)
	go crawler.Start(pages, queueRawData)

	go store.StoreToCSV(queueSaveCSV, filename)
	go clean.Consume(queueRawData, queueSaveCSV, config.StartDate, config.EndDate)

	defer close(queueSaveCSV)
	defer close(queueRawData)
	defer close(pages)

	for {
		time.Sleep(1 * time.Hour)
	}
}
func getConfig(urlConfig string) (*Config, error) {
	config, err := ioutil.ReadFile(urlConfig)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(config, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
