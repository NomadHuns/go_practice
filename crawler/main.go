package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 사이트 상태를 체크하는 함수
func checkStatus(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // 함수 종료 시 WaitGroup 카운터 감소

	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		results <- fmt.Sprintf("[ERROR] %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("[%d OK] %s (%v)", resp.StatusCode, url, elapsed)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.go.dev",
		"https://www.stackoverflow.com",
		"https://www.netflix.com",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // 결과 전달을 위한 채널

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)                         // 실행할 고루틴 개수 추가
		go checkStatus(url, &wg, results) // 고루틴 실행
	}

	// 모든 고루틴이 끝날 때까지 기다림
	go func() {
		wg.Wait()
		close(results) // 더 이상 보낼 데이터가 없으면 채널을 닫음
	}()

	// 결과 출력
	for res := range results {
		fmt.Println(res)
	}

	fmt.Printf("\n전체 소요 시간: %v\n", time.Since(startTime))
}
