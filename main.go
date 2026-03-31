package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
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

// 파일명을 .json으로 변경
const fileName = "history.json"

// [기능 분리] 저장 로직
func saveAsJSON(history []Calculation) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// [기능 분리] 불러오기 로직
func loadFromJSON() ([]Calculation, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Calculation{}, nil // 파일이 없으면 빈 슬라이스 반환
		}
		return nil, err
	}

	var history []Calculation
	err = json.Unmarshal(data, &history)
	return history, err
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

func main() {
	// 1. 시작하자마자 기존 기록 불러오기
	history, err := loadFromJSON()
	if err != nil {
		fmt.Println("⚠️ 데이터를 불러오는 중 오류 발생:", err)
		history = []Calculation{} // 에러 발생 시 빈 상태로 시작
	}

	// 2. 기존 기록이 있다면 출력
	if len(history) > 0 {
		fmt.Println("📜 [이전 연산 기록을 불러왔습니다]")
		for i, c := range history {
			fmt.Printf("[%d] %s\n", i+1, c.ToString())
		}
	} else {
		fmt.Println("✨ 기존 기록이 없습니다. 새로운 계산을 시작합니다.")
	}

	for {
		var n1, n2 float64
		var op, input string

		fmt.Print("\n첫 번째 숫자 (종료: q): ")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "q" {
			break
		}
		fmt.Sscanf(input, "%f", &n1)

		fmt.Print("연산자: ")
		fmt.Scanln(&op)
		fmt.Print("두 번째 숫자: ")
		fmt.Scanln(&n2)

		res, err := calculate(n1, n2, op)
		if err != nil {
			fmt.Println("오류:", err)
			continue
		}

		// 4. 새로운 구조체 인스턴스 생성 및 슬라이스 추가
		calc := Calculation{
			Num1:   n1,
			Num2:   n2,
			Op:     op,
			Result: res,
		}
		history = append(history, calc)

		fmt.Printf("결과: %s\n", calc.ToString())
	}

	// 4. 종료 전 저장
	if len(history) > 0 {
		fmt.Print("\n💾 기록을 저장하시겠습니까? (y/n): ")
		var saveConfirm string
		fmt.Scanln(&saveConfirm)

		if strings.ToLower(saveConfirm) == "y" {
			if err := saveAsJSON(history); err != nil {
				fmt.Println("❌ 저장 실패:", err)
			} else {
				fmt.Println("🎉 history.json에 성공적으로 저장되었습니다!")
			}
		}
	}
}
