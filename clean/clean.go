package clean

import (
	"fmt"
	"regexp"
	"strings"

	"crawl/clients"
	"github.com/PuerkitoBio/goquery"
)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func clean(rawPageContent string, result chan string, utilsClient *clients.UtilsClient) error {
	var validID = regexp.MustCompile(`^[0-9]{2}/[0-9]{2}/[0-9]{4}$`)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawPageContent))
	if err != nil {
		return err
	}

	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		date := ""
		title := ""
		address := ""
		phone := ""
		price := ""
		detail := ""
		s.Find("td").Each(func(itd int, s *goquery.Selection) {
			s.Find("span").Each(func(i int, s *goquery.Selection) {
				rawText := strings.TrimSpace(s.Text())

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
				}
			})

			s.Find("p").Each(func(i int, s *goquery.Selection) {
				rawStr := s.Text()
				rawStr = strings.TrimSpace(rawStr)

				if validID.MatchString(rawStr) {
					date = rawStr
				}
			})

			s.Find("div button").Each(func(i int, s *goquery.Selection) {
				slug, _ :=  s.Attr("data-slug")
				slug = strings.TrimSpace(slug)
				//fmt.Println("[INFOR] get-slug success = ", slug, ", raw_text in slug = ", s.Text())
				if len(slug) > 0 {
					rawSlug, err := utilsClient.GetString(slug)
					if err != nil {
						fmt.Println("[ERROR] get-slug error = ", err.Error())
					} else {
						detail = rawSlug
					}
				}
			})


		})

		oneRow := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\" \n", date, title, address, address, phone, price, detail)
		fmt.Println("[DB] data = ", oneRow)
		result <- oneRow
	})


	return nil
}

func Consume(queue chan string, result chan string, utilsClient *clients.UtilsClient) {
	for rawPageConnect := range queue {
		err := clean(rawPageConnect, result, utilsClient)
		if err != nil {
			fmt.Printf("[CLEAN][ERROR] clean raw data error %s\n", err.Error())
			continue
		}

		fmt.Println("[CLEAN][INFOR] clean success")
	}
}
