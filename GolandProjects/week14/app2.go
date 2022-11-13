package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http" // 네트워크계층/전송계층

)

type WebPage struct {
	URL string
	Size int
}

func responseSize(url string, channel chan WebPage)  {		// ?? 연결 역할을 해줄 channel 변수
	res, e := http.Get(url)

	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()

	body, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}
	channel <- WebPage{URL: url, Size: len(body)}		// page와 pages channel이 연결되었다.

	// fmt.Println(url, len(body))
}



func main() {
	urls := []string{"http://www.inhatc.ac.kr", "http://www.harvard.edu", "http://www.naver.com", "http://www.daum.net"}
	pages := make(chan WebPage)		// 비어있는 slice 만들 때 make

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <- pages	// pages에서 하나식 channel을 가져온다.
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	//go responseSize("http://www.inhatc.ac.kr")
	//go responseSize("http://www.harvard.edu")
	//go responseSize("http://www.naver.com")
	//go responseSize("http://www.daum.net")
	//하드코딩 제거

	// 채널이 있다면 sleep을 안 써도 된다.
}
