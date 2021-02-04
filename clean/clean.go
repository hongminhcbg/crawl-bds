package clean

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func clean(rawPageContent string, result chan string) error {
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

func Consume(queue chan string, result chan string) {
	for rawPageConnect := range queue {
		err := clean(rawPageConnect, result)
		if err != nil {
			fmt.Printf("[CLEAN][ERROR] clean raw data error %s\n", err.Error())
			continue
		}

		fmt.Println("[CLEAN][INFOR] clean success")
	}
}
