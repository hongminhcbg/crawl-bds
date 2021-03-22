package clean

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getDayFromDate(data string) int {
	args := strings.Split(data, "/")
	if len(args) != 3 {
		return 0
	}
	day, _ := strconv.Atoi(args[0])
	month, _ := strconv.Atoi(args[1])
	year, _ := strconv.Atoi(args[2])

	return day + 31*month + 366*year
}

func compareDate(start, end, b string) int {

	startDay := getDayFromDate(start)
	endDay := getDayFromDate(end)

	if endDay < startDay {
		panic("start day = " + start + ", end day = " + end + ".... invalid")
	}

	input := getDayFromDate(b)
	if input <= endDay && input >= startDay {
		return 0
	}

	if input < startDay {
		return -1
	}

	return 1
}

func clean(rawPageContent string, result chan string, StartDate string, endDate string) error {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawPageContent))
	if err != nil {
		return err
	}

	var date string
	var catalog string
	var address string
	var phone string
	var content string
	var district string
	var price string

	doc.Find(".item-bds").Each(func(i int, s *goquery.Selection) {
		date = ""
		catalog = ""
		address = ""
		phone = ""
		content = ""
		price = ""

		date = s.Find(".text-center").Text()[:10]
		fmt.Println("[DB][Compare date] ======>", date, StartDate)
		compareStartDateWithDate := compareDate(StartDate, endDate, date)
		if compareStartDateWithDate == 1 {
			fmt.Println(" ============> download success <============ ")
			time.Sleep(20 * time.Second)
			os.Exit(2)
		}

		if compareStartDateWithDate == -1 {
			return
		}

		//infor type
		catalog = s.Find(".has-bg").Find("span").Text()

		s.Find(".item-bds-info").Find(".bds-item-content").Find(".sub-table").Find("td").Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				address = s.Text()
				return
			}

			if i == 3 {
				phone = s.Text()
				return
			}
		})

		s.Find(".item-bds-more").Find("b").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				district = s.Text()
				return
			}

			price = s.Text()
		})

		// get data
		content = s.Find(".item-bds-title").Find("a").Text() + s.Find(".bds-item-content-high").Find("p").Text()
		content = standardizeSpaces(content)
		oneRow := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\" \n", date, catalog, address, district, phone, price, content)

		fmt.Println("[DEBUG] ", oneRow)

		result <- oneRow
	})

	return nil
}

func Consume(queue chan string, result chan string, StartDate string, endDate string) {
	for rawPageConnect := range queue {
		err := clean(rawPageConnect, result, StartDate, endDate)
		if err != nil {
			fmt.Printf("[CLEAN][ERROR] clean raw data error %s\n", err.Error())
			continue
		}

		fmt.Println("[CLEAN][INFOR] clean success")
	}
}
