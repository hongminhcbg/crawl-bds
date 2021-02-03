package main

import (
	"fmt"
	"time"

	"crawl/store"
)

const cookie = "__cfduid=d47121bbf4a11649bb87c1c2897ebe3e11612182600; __sbmask=acqglygatyyalzijgctv@usqxgkanhrmoasbblkauw@4oXYd/I+Tcg67DWfaP8bWMWRbttC0LBlg++Lbg%3D%3D; wordpress_test_cookie=WP+Cookie+check; wordpress_logged_in_3748aa90f9091fbd66dfda219c76b982=mr.nvlam%40gmail.com%7C1612455365%7CocjNTXiQPUhxmftozhDkbtmfSaMcQKffb6OkDDgdERK%7C7b66b2a7605bb4073fe6b6216cb82368f167b44bb851b7f5d5772111ea3c68fe; __atuvc=28%7C5; __atuvs=60198364c0e89a3a000"
const requestURL = "https://batdongsanchinhchu.vn/bds"

func main() {
	fmt.Println("hello world")

	queue := make(chan string, 5000)
	for i := 1; i < 10; i++ {
		queue <- fmt.Sprintf("%d,%d,%d\n", i, i+1, i+2)
	}

	filename := fmt.Sprintf("./results/%d_result.csv", time.Now().Unix())
	go store.StoreToCSV(queue, filename)

	time.Sleep(30 * time.Second)
	close(queue)
}
