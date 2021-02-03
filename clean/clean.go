package clean

import (
	"fmt"
)

func clean(rawPageConnect string) (string, error) {
	return "", nil
}

func Consume(queue chan string, result chan string) {
	for rawPageConnect := range queue {
		cleanData, err := clean(rawPageConnect)
		if err != nil {
			fmt.Printf("[CLEAN][ERROR] clean raw data error %s\n", err.Error())
			continue
		}

		fmt.Println("[CLEAN][INFOR] clean success, ", cleanData)
		result <- cleanData
	}
}
