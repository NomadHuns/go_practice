package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Calculation struct {
	Num1   float64 `json:"num1"`
	Num2   float64 `json:"num2"`
	Op     string  `json:"operator"`
	Result float64 `json:"result"`
}

type PageData struct {
	Current Calculation
	History []Calculation
	Error   string
}

const fileName = "history.json"

// 파일 관리 함수들
func saveToFile(history []Calculation) {
	data, _ := json.MarshalIndent(history, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func loadFromFile() []Calculation {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Calculation{}
	}
	var history []Calculation
	json.Unmarshal(data, &history)
	return history
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")

	// 1. 기존 기록 무조건 불러오기
	allHistory := loadFromFile()

	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	data := PageData{History: allHistory}

	// 2. 계산 요청이 들어온 경우
	if aStr != "" && bStr != "" {
		a, _ := strconv.ParseFloat(aStr, 64)
		b, _ := strconv.ParseFloat(bStr, 64)

		// 간단 계산 로직 (에러 체크 생략 버전)
		res := 0.0
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b != 0 {
				res = a / b
			} else {
				data.Error = "0으로 나눌 수 없음"
			}
		}

		if data.Error == "" {
			newCalc := Calculation{Num1: a, Num2: b, Op: op, Result: res}
			data.Current = newCalc

			// 3. 새 기록 추가 및 파일 저장
			allHistory = append([]Calculation{newCalc}, allHistory...) // 최신순 정렬
			if len(allHistory) > 5 {
				allHistory = allHistory[:5]
			} // 5개만 유지
			saveToFile(allHistory)
			data.History = allHistory
		}
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", mainHandler)
	fmt.Println("🌐 http://localhost:8080 에서 확인하세요!")
	http.ListenAndServe(":8080", nil)
}
