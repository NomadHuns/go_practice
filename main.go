package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// 1. 계산 기록을 담을 구조체 정의
type Calculation struct {
	Num1, Num2, Result float64
	Op                 string
}

// 2. 구조체 데이터를 예쁘게 출력해주는 메서드 (선택 사항)
func (c Calculation) ToString() string {
	return fmt.Sprintf("%.2f %s %.2f = %.2f", c.Num1, c.Op, c.Num2, c.Result)
}

func calculate(n1, n2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errors.New("0으로 나눌 수 없음")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("잘못된 연산자")
	}
}

// 웹 요청을 처리하는 핸들러 함수
func calcHandler(w http.ResponseWriter, r *http.Request) {
	// 1. URL 파라미터 읽기 (예: ?a=10&b=5&op=+)
	query := r.URL.Query()

	aStr := query.Get("a")
	bStr := query.Get("b")
	op := query.Get("op")

	// 2. 문자열 숫자를 float64로 변환
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)

	// 3. 계산 수행
	res, err := calculate(a, b, op)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. 결과를 JSON으로 응답
	resultObj := Calculation{Num1: a, Num2: b, Op: op, Result: res}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultObj)
}

func main() {
	// 주소와 핸들러 연결
	http.HandleFunc("/calculate", calcHandler)

	fmt.Println("🚀 서버가 http://localhost:8080 에서 시작되었습니다!")
	// 서버 실행
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("서버 실행 실패:", err)
	}
}
