package main

import (
	"errors"
	"fmt"
	"strings"
)

// 1. 계산 기록을 담을 구조체 정의
type Calculation struct {
	Num1   float64
	Num2   float64
	Op     string
	Result float64
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

func main() {
	// 3. 문자열 슬라이스가 아닌 'Calculation 구조체 슬라이스' 선언
	var history []Calculation

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

	// 최종 기록 출력
	fmt.Println("\n--- 전체 계산 로그 ---")
	for i, c := range history {
		fmt.Printf("[%d] %8.2f %s %8.2f = %10.2f\n", i+1, c.Num1, c.Op, c.Num2, c.Result)
	}
}
