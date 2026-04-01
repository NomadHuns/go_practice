package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 1. 로깅 미들웨어: 요청 시간과 메서드를 기록
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("시작: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r) // 다음 미들웨어 또는 핸들러 호출

		log.Printf("완료: %s 소요시간: %v", r.URL.Path, time.Since(start))
	})
}

// 2. 타임아웃 미들웨어: 모든 요청에 Context 타임아웃 설정
func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 전체 요청 처리 시간을 3초로 제한
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		// 타임아웃이 적용된 새로운 Context를 요청 객체에 주입
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// 3. 실제 비즈니스 로직 핸들러
func helloHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(2 * time.Second): // 2초 걸리는 작업 시뮬레이션
		fmt.Fprintln(w, "안녕하세요! 작업이 완료되었습니다.")
	case <-r.Context().Done(): // 3초 타임아웃에 걸릴 경우
		log.Println("핸들러: 클라이언트 요청 취소 또는 타임아웃")
		http.Error(w, "요청 처리 시간 초과", http.StatusRequestTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// 4. 미들웨어 체이닝 (오른쪽에서 왼쪽 순으로 감싸짐)
	// Logging -> Timeout -> helloHandler 순서로 실행
	finalHandler := LoggingMiddleware(TimeoutMiddleware(mux))

	fmt.Println("서버가 :8080에서 시작됩니다...")
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
