package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 구조체와 인터페이스: Go는 클래스 대신 구조체를 사용합니다.
type User struct {
	ID   int
	Name string
}

// 2. 고루틴을 이용한 가상의 DB 저장 함수
func saveUser(id int, name string, wg *sync.WaitGroup, results chan<- string) {
	// 함수 종료 시 WaitGroup 카운트 감소
	defer wg.Done()

	// 가상의 네트워크/DB 지연 시간 시뮬레이션 (PostgreSQL 저장 가정)
	time.Sleep(time.Millisecond * 500)

	result := fmt.Sprintf("유저 %d(%s) 저장 완료", id, name)

	// 3. 채널을 통해 결과 전송
	results <- result
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	// 고루틴 동기화를 위한 WaitGroup
	var wg sync.WaitGroup
	// 결과 수집을 위한 채널
	results := make(chan string, len(users))

	fmt.Println("데이터 저장을 시작합니다...")
	start := time.Now()

	for _, u := range users {
		wg.Add(1)
		// 4. 'go' 키워드 하나로 비동기 실행 (고루틴)
		go saveUser(u.ID, u.Name, &wg, results)
	}

	// 모든 고루틴이 끝날 때까지 대기하고 채널 닫기
	go func() {
		wg.Wait()
		close(results)
	}()

	// 5. 채널에 쌓인 결과 출력
	for res := range results {
		fmt.Println("-", res)
	}

	fmt.Printf("전체 소요 시간: %v\n", time.Since(start))
	fmt.Println("모든 작업이 안전하게 종료되었습니다.")
}
