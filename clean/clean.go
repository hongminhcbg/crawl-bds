package clean

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func clean(rawPageContent string, result chan string) error {
	var validID = regexp.MustCompile(`^[0-9]{2}/[0-9]{2}/[0-9]{4}$`)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawPageContent))
	if err != nil {
		return err
	}

	// var date string
	// var catalog string
	// var address string
	// var phone string
	// var content string
	// var district string
	// var price string

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		date := "";
		title := ""
		address := ""
		phone := ""
		price := ""
		square := ""
		detail := ""
		s.Find("td").Each(func(itd int, s *goquery.Selection) {

			s.Find("span").Each(func(i int, s *goquery.Selection) {
				rawText := strings.TrimSpace(s.Text())
				//fmt.Printf("[BD][itd = %d][ispan = %d] %s\n", itd, i, rawText)

				switch i {
				case 0:
					if itd == 1 {
						title = rawText
					} else {
						address = rawText
					}
				case 1:
					phone = rawText
				case 3:
					price = rawText
				case 4:
					square = rawText
				}
			})

			s.Find("p").Each(func(i int, s *goquery.Selection) {
				rawStr := s.Text()
				rawStr = strings.TrimSpace(rawStr)
				//fmt.Printf("rawstr = '%s'\n", rawStr)

				if validID.MatchString(rawStr) {
					date = rawStr;
				}
			})
	
		})

		fmt.Println("[DB] date = ", date)
		fmt.Println("[DB] title = ", title)
		fmt.Println("[DB] phone = ", phone)
		fmt.Println("[DB] price = ", price)
		fmt.Println("[DB] square = ", square)
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
