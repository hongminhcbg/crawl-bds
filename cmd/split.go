package main

import (
	"bufio"
	"crawl/store"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const fileName = "./results/2019_3"

const rent = "thuê"
const Rent = "Thuê"

func getInforType(input string) string {
	args := strings.Split(input, ",")
	if len(args) < 2 {
		return ""
	}

	return args[1]
}

func isRentInfor(input string) bool {
	inforType := getInforType(input)

	if strings.Contains(inforType, rent) || strings.Contains(inforType, Rent) {
		fmt.Println("[DEBUG][RENT] ", input)
		return true
	}

	fmt.Println("[DEBUG][SELL] ", input)
	return false
}

func main() {
	fileNameRoot := fileName + ".csv"

	fileNameRent := fileName + "_thue.csv"
	fileNameSell := fileName + "_ban.csv"

	chanRent := make(chan string, 3000)
	chanSell := make(chan string, 3000)

	go store.StoreToCSV(chanRent, fileNameRent)
	go store.StoreToCSV(chanSell, fileNameSell)

	file, err := os.Open(fileNameRoot)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if isRentInfor(line) {
			chanRent <- line + "\n"
			continue
		}

		chanSell <- line + "\n"
	}

	time.Sleep(10 * time.Second)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
