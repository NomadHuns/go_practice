package main

import (
	"fmt"
	"html/template" // HTML 템플릿 처리를 위해 추가
	"net/http"
	"strconv"
)

// 화면에 전달할 데이터 구조체
type PageData struct {
	Num1, Num2, Result float64
	Op, Error          string
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
			return 0, fmt.Errorf("0으로 나눌 수 없습니다")
		}
		return n1 / n2, nil
	default:
		return 0, fmt.Errorf("잘못된 연산자")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 템플릿 파일 파싱
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "템플릿을 찾을 수 없습니다", http.StatusInternalServerError)
		return
	}

	// 2. 폼 데이터 읽기 (GET 방식)
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	data := PageData{}

	// 3. 값이 있을 때만 계산 수행
	if aStr != "" && bStr != "" {
		a, _ := strconv.ParseFloat(aStr, 64)
		b, _ := strconv.ParseFloat(bStr, 64)
		res, calcErr := calculate(a, b, op)

		data = PageData{Num1: a, Num2: b, Op: op, Result: res}
		if calcErr != nil {
			data.Error = calcErr.Error()
		}
	}

	// 4. 템플릿에 데이터 주입하여 응답
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", mainHandler) // 모든 접속을 핸들러로 연결

	fmt.Println("🌐 서버 시작: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
