package store

import (
	"io/ioutil"
	"log"
	"os"
)

const initFile = "\"Ngày\",\"Danh mục\",\"Nội dung\" \n"

func StoreToCSV(queue chan string, filename string) {

	ioutil.WriteFile(filename, []byte(initFile), 0666)

	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for row := range queue {
		f.WriteString(row)
	}
}
