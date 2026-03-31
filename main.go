package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("--- Go 계산기 프로그램 ---")

	// 1. 첫 번째 숫자 입력
	fmt.Print("첫 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num1)

	// 2. 연산자 입력
	fmt.Print("연산자를 입력하세요 (+, -, *, /): ")
	fmt.Scanln(&operator)

	// 3. 두 번째 숫자 입력
	fmt.Print("두 번째 숫자를 입력하세요: ")
	fmt.Scanln(&num2)

	// 4. 연산 수행 및 결과 출력
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		// 에러 처리: 0으로 나누기 방지
		if num2 == 0 {
			fmt.Println("오류: 0으로 나눌 수 없습니다.")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("잘못된 연산자입니다.")
	}
}
