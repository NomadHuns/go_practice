package main

import (
	"fmt"
	"strings" // 문자열 처리를 위해 추가
)

func main() {
	for { // <--- 무한 루프 시작
		var num1, num2 float64
		var operator string

		fmt.Println("\n--- 계산기를 시작합니다 (종료하려면 'q' 입력) ---")

		fmt.Print("첫 번째 숫자: ")
		_, err := fmt.Scanln(&num1)
		if err != nil { // 숫자가 아닌 값이 들어오면 종료 확인 로직으로 이동
			break
		}

		fmt.Print("연산자 (+, -, *, /): ")
		fmt.Scanln(&operator)

		fmt.Print("두 번째 숫자: ")
		fmt.Scanln(&num2)

		switch operator {
		case "+":
			fmt.Printf("결과: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("결과: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("결과: %.2f\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("오류: 0으로 나눌 수 없습니다.")
			} else {
				fmt.Printf("결과: %.2f\n", num1/num2)
			}
		default:
			fmt.Println("잘못된 연산자입니다.")
		}

		// 계속할지 물어보기
		var choice string
		fmt.Print("계속하시겠습니까? (y/n): ")
		fmt.Scanln(&choice)

		// 대문자/소문자 구분 없이 처리하기 위해 Lower 사용
		if strings.ToLower(choice) == "n" {
			fmt.Println("프로그램을 종료합니다.")
			break // <--- 루프 탈출
		}
	}
}
