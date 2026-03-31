package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 1. 결과를 담을 구조체 정의
type Result struct {
	URL          string
	StatusCode   int
	Duration     time.Duration
	ErrorMessage error
}

func checkStatus(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	start := time.Now()
	// 타임아웃 설정을 추가한 클라이언트 (권장사항)
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	elapsed := time.Since(start)

	// 2. 구조체 인스턴스를 생성하여 채널로 전송
	if err != nil {
		results <- Result{
			URL:          url,
			ErrorMessage: err,
			Duration:     elapsed,
		}
		return
	}
	defer resp.Body.Close()

	results <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   elapsed,
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",               // 간혹 접속이 느릴 수 있음
		"https://this-site-does-not-exist.com", // 에러 테스트용
	}

	var wg sync.WaitGroup
	// 3. Result 타입을 주고받는 채널 생성
	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go checkStatus(url, &wg, results)
	}

	// 고루틴 감시자
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("--- 크롤링 결과 분석 ---")

	// 4. 채널에서 구조체를 꺼내어 활용
	for res := range results {
		if res.ErrorMessage != nil {
			fmt.Printf("❌ 실패: %-30s | 에러: %v\n", res.URL, res.ErrorMessage)
		} else {
			fmt.Printf("✅ 성공: %-30s | 상태: %d | 소요시간: %v\n",
				res.URL, res.StatusCode, res.Duration)
		}
	}
}
