package main

import (
	"errors"
	"fmt"
	"strings"
)

// 계산 로직을 담당하는 함수
// 결과값(float64)과 에러(error)를 반환합니다.
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
			// 에러 객체를 생성하여 반환
			return 0, errors.New("0으로 나눌 수 없습니다")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("지원하지 않는 연산자입니다")
	}
}

func main() {
	fmt.Println("=== Go 계산기 프로그램 (종료: q 입력) ===")

	for {
		var input1, input2, op string
		var n1, n2 float64

		// 1. 첫 번째 숫자 입력 (또는 종료 문자)
		fmt.Print("\n첫 번째 숫자 (혹은 'q'로 종료): ")
		fmt.Scanln(&input1)
		if strings.ToLower(input1) == "q" {
			break
		}
		// 문자열을 숫자로 변환 (단순화를 위해 Sscanf 사용)
		fmt.Sscanf(input1, "%f", &n1)

		// 2. 연산자 입력
		fmt.Print("연산자 (+, -, *, /): ")
		fmt.Scanln(&op)

		// 3. 두 번째 숫자 입력
		fmt.Print("두 번째 숫자: ")
		fmt.Scanln(&input2)
		fmt.Sscanf(input2, "%f", &n2)

		// 4. 함수 호출 및 결과 처리
		result, err := calculate(n1, n2, op)

		if err != nil {
			// 에러가 발생한 경우 (nil이 아닌 경우)
			fmt.Printf("⚠️  오류 발생: %v\n", err)
		} else {
			// 정상적인 경우
			fmt.Printf("✅ 결과: %.2f %s %.2f = %.2f\n", n1, op, n2, result)
		}
	}

	fmt.Println("프로그램을 종료합니다. 이용해 주셔서 감사합니다!")
}
