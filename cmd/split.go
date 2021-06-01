package main

import (
	"bufio"
	"crawl/store"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"strconv"
	"regexp"
)

const fileName = "./results/2021"

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

func getMonth(input string) int {
	month := input[4:6]
	monthInt, _ := strconv.Atoi(month)
	fmt.Println("full date = ", input[:13], ", monthInt = ", monthInt)

	return monthInt;
	return 0
}

func main() {
	fileName1Rent := fileName + "_1_thue.csv"
	fileName1Sell := fileName + "_1_ban.csv"
	fileName2Rent := fileName + "_2_thue.csv"
	fileName2Sell := fileName + "_2_ban.csv"
	fileName3Rent := fileName + "_3_thue.csv"
	fileName3Sell := fileName + "_3_ban.csv"
	fileName4Rent := fileName + "_4_thue.csv"
	fileName4Sell := fileName + "_4_ban.csv"
	fileName5Rent := fileName + "_5_thue.csv"
	fileName5Sell := fileName + "_5_ban.csv"
	fileName6Rent := fileName + "_6_thue.csv"
	fileName6Sell := fileName + "_6_ban.csv"
	fileName7Rent := fileName + "_7_thue.csv"
	fileName7Sell := fileName + "_7_ban.csv"
	fileName8Rent := fileName + "_8_thue.csv"
	fileName8Sell := fileName + "_8_ban.csv"
	fileName9Rent := fileName + "_9_thue.csv"
	fileName9Sell := fileName + "_9_ban.csv"
	fileName10Rent := fileName + "_10_thue.csv"
	fileName10Sell := fileName + "_10_ban.csv"
	fileName11Rent := fileName + "_11_thue.csv"
	fileName11Sell := fileName + "_11_ban.csv"
	fileName12Rent := fileName + "_12_thue.csv"
	fileName12Sell := fileName + "_12_ban.csv"

	chanRent1 := make(chan string, 20)
	chanSell1 := make(chan string, 20)
	chanRent2 := make(chan string, 20)
	chanSell2 := make(chan string, 20)
	chanRent3 := make(chan string, 20)
	chanSell3 := make(chan string, 20)
	chanRent4 := make(chan string, 20)
	chanSell4 := make(chan string, 20)
	chanRent5 := make(chan string, 20)
	chanSell5 := make(chan string, 20)
	chanRent6 := make(chan string, 20)
	chanSell6 := make(chan string, 20)
	chanRent7 := make(chan string, 20)
	chanSell7 := make(chan string, 20)
	chanRent8 := make(chan string, 20)
	chanSell8 := make(chan string, 20)
	chanRent9 := make(chan string, 20)
	chanSell9 := make(chan string, 20)
	chanRent10 := make(chan string, 20)
	chanSell10 := make(chan string, 20)
	chanRent11 := make(chan string, 20)
	chanSell11 := make(chan string, 20)
	chanRent12 := make(chan string, 20)
	chanSell12 := make(chan string, 20)

	go store.StoreToCSV(chanRent1, fileName1Rent)
	go store.StoreToCSV(chanSell1, fileName1Sell)
	go store.StoreToCSV(chanRent2, fileName2Rent)
	go store.StoreToCSV(chanSell2, fileName2Sell)
	go store.StoreToCSV(chanRent3, fileName3Rent)
	go store.StoreToCSV(chanSell3, fileName3Sell)
	go store.StoreToCSV(chanRent4, fileName4Rent)
	go store.StoreToCSV(chanSell4, fileName4Sell)
	go store.StoreToCSV(chanRent5, fileName5Rent)
	go store.StoreToCSV(chanSell5, fileName5Sell)
	go store.StoreToCSV(chanRent6, fileName6Rent)
	go store.StoreToCSV(chanSell6, fileName6Sell)
	go store.StoreToCSV(chanRent7, fileName7Rent)
	go store.StoreToCSV(chanSell7, fileName7Sell)
	go store.StoreToCSV(chanRent8, fileName8Rent)
	go store.StoreToCSV(chanSell8, fileName8Sell)
	go store.StoreToCSV(chanRent9, fileName9Rent)
	go store.StoreToCSV(chanSell9, fileName9Sell)
	go store.StoreToCSV(chanRent10, fileName10Rent)
	go store.StoreToCSV(chanSell10, fileName10Sell)
	go store.StoreToCSV(chanRent11, fileName11Rent)
	go store.StoreToCSV(chanSell11, fileName11Sell)
	go store.StoreToCSV(chanRent12, fileName12Rent)
	go store.StoreToCSV(chanSell12, fileName12Sell)

	fileNameRoot := fileName + ".csv"
	file, err := os.Open(fileNameRoot)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var validID = regexp.MustCompile(`^[0-9]{2}/[0-9]{2}/[0-9]{4}$`)
	var line string
	var lineNow string
	for scanner.Scan() {
		lineNow = scanner.Text()

		if len(lineNow) < 11 || !validID.MatchString(lineNow[1:11]) {
			line = line + lineNow + "\n"
			continue
		}

		if len(line) == 0 {
			continue
		}

		line_temp := line
		line = lineNow
		month := getMonth(line_temp)
		if isRentInfor(line_temp) {
			switch (month) {
				case 1:
								chanRent1 <- line_temp + "\n"
				case 2:
								chanRent2 <- line_temp + "\n"
				case 3:
								chanRent3 <- line_temp + "\n"
				case 4:
								chanRent4 <- line_temp + "\n"
				case 5:
								chanRent5 <- line_temp + "\n"
				case 6:
								chanRent6 <- line_temp + "\n"
				case 7:
								chanRent7 <- line_temp + "\n"
				case 8:
								chanRent8 <- line_temp + "\n"
				case 9:
								chanRent9 <- line_temp + "\n"
				case 10:
								chanRent10 <- line_temp + "\n"
				case 11:
								chanRent11 <- line_temp + "\n"
				case 12:
								chanRent12 <- line_temp + "\n"			
			}
			continue
		}

		switch (month) {
			case 1:
							chanSell1 <- line_temp + "\n"
			case 2:
							chanSell2 <- line_temp + "\n"
			case 3:
							chanSell3 <- line_temp + "\n"
			case 4:
							chanSell4 <- line_temp + "\n"
			case 5:
							chanSell5 <- line_temp + "\n"
			case 6:
							chanSell6 <- line_temp + "\n"
			case 7:
							chanSell7 <- line_temp + "\n"
			case 8:
							chanSell8 <- line_temp + "\n"
			case 9:
							chanSell9 <- line_temp + "\n"
			case 10:
							chanSell10 <- line_temp + "\n"
			case 11:
							chanSell11 <- line_temp + "\n"
			case 12:
							chanSell12 <- line_temp + "\n"
		}
}

	time.Sleep(10 * time.Second)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
