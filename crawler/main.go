package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	URL        string
	StatusCode int
}

// 1. 워커 함수: jobs 채널에서 일을 받아 처리하고 결과를 results 채널로 보냄
func worker(id int, jobs <-chan string, results chan<- Result) {
	for url := range jobs {
		fmt.Printf("워커 %d가 %s 처리 중...\n", id, url)
		resp, err := http.Get(url)

		status := 0
		if err == nil {
			status = resp.StatusCode
			resp.Body.Close()
		}

		results <- Result{URL: url, StatusCode: status}
	}
}

func main() {
	urls := []string{
		"https://google.com", "https://github.com", "https://golang.org",
		"https://facebook.com", "https://twitter.com", "https://naver.com",
	}

	const workerCount = 3 // 동시에 일할 워커 수
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	// 2. 워커 고루틴들을 미리 실행
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}

	// 3. 일거리(URL)를 전송
	for _, url := range urls {
		jobs <- url
	}
	close(jobs) // 모든 일을 다 넣었으므로 jobs 채널을 닫음

	// 4. 결과 수집
	for i := 0; i < len(urls); i++ {
		res := <-results
		fmt.Printf("결과 수신 -> URL: %s, 상태: %d\n", res.URL, res.StatusCode)
	}
}
