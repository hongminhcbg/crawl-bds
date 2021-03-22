package store

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const initFile = `"Ngày","Danh mục", "Địa chỉ", "Quận","Số điện thoại","Giá","Nội dung"`

func StoreToCSV(queue chan string, filename string) {

	ioutil.WriteFile(filename, []byte(initFile+"\n"), 0666)

	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for {
		select {
		case row := <-queue:
			f.WriteString(row)
		case <-time.After(30 * time.Minute):
			os.Exit(2)
		}
	}
}
