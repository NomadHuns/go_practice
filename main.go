package main

import (
	"errors"
	"fmt"
	"strings"
)

// 계산 함수 (이전과 동일)
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
			return 0, errors.New("0으로 나눌 수 없습니다")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("지원하지 않는 연산자입니다")
	}
}

func main() {
	// 1. 기록을 저장할 슬라이스 선언
	var history []string

	fmt.Println("=== Go 계산기 (기록 저장 기능 포함) ===")

	for {
		var input1, input2, op string
		var n1, n2 float64

		fmt.Print("\n첫 번째 숫자 (종료: q): ")
		fmt.Scanln(&input1)
		if strings.ToLower(input1) == "q" {
			break
		}
		fmt.Sscanf(input1, "%f", &n1)

		fmt.Print("연산자 (+, -, *, /): ")
		fmt.Scanln(&op)

		fmt.Print("두 번째 숫자: ")
		fmt.Scanln(&input2)
		fmt.Sscanf(input2, "%f", &n2)

		result, err := calculate(n1, n2, op)

		if err != nil {
			fmt.Printf("⚠️  오류: %v\n", err)
		} else {
			// 2. 결과를 문자열로 만들어 기록에 추가
			record := fmt.Sprintf("%.2f %s %.2f = %.2f", n1, op, n2, result)
			history = append(history, record) // 슬라이스에 추가

			fmt.Printf("✅ 결과: %s\n", record)
		}
	}

	// 3. 루프 종료 후 한꺼번에 출력
	fmt.Println("\n============================")
	fmt.Println("        📜 연산 기록        ")
	fmt.Println("============================")

	if len(history) == 0 {
		fmt.Println("저장된 기록이 없습니다.")
	} else {
		// range를 사용하여 슬라이스 순회
		for i, v := range history {
			fmt.Printf("%d. %s\n", i+1, v)
		}
	}
	fmt.Println("============================")
}
