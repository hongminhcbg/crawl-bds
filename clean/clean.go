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

	var data string

	doc.Find(".item-bds").Each(func(i int, s *goquery.Selection) {
		oneRow := ""
		// get date
		date := s.Find(".text-center").Text()
		oneRow += fmt.Sprintf(`"%s",`, date[:10])

		//infor type
		typeInfor := s.Find(".has-bg").Find("span").Text()
		oneRow += fmt.Sprintf(`"%s",`, standardizeSpaces(typeInfor))

		// get data
		data = s.Find(".bds-item-content-high").Text() + "\n" + s.Find(".item-bds-more").Text()
		oneRow += fmt.Sprintf(`"%s"`, standardizeSpaces(data))
		oneRow += "\n"

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
