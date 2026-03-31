package main

import (
	"bufio"
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

// 파일에 기록을 저장하는 전용 함수
func saveHistory(history []Calculation) error {
	// 1. 파일 열기 (없으면 생성, 있으면 끝에 추가, 쓰기 전용 권한 0644)
	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 2. defer: 함수가 끝나기 직전에 무조건 파일을 닫도록 예약 (중요!)
	defer file.Close()

	fmt.Fprintln(file, "\n--- 새로운 세션 기록 ---")
	for _, c := range history {
		// 3. fmt.Fprintln을 사용해 콘솔이 아닌 'file' 객체에 쓰기
		_, err := fmt.Fprintln(file, c.ToString())
		if err != nil {
			return err
		}
	}
	return nil
}

// 파일에서 이전 기록을 읽어와 출력하는 함수
func loadHistory() {
	// 1. 파일 열기 (읽기 전용)
	file, err := os.Open("history.txt")
	if err != nil {
		// 파일이 없으면 에러가 나지만, 처음 실행 시에는 당연한 것이므로 무시합니다.
		if os.IsNotExist(err) {
			fmt.Println("📜 이전 기록이 없습니다. 새로운 세션을 시작합니다.")
			return
		}
		fmt.Println("⚠️ 파일을 읽는 중 오류 발생:", err)
		return
	}
	defer file.Close()

	fmt.Println("============================")
	fmt.Println("      📂 이전 연산 기록      ")
	fmt.Println("============================")

	// 2. Scanner를 사용하여 한 줄씩 읽기
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 한 줄 읽어서 콘솔에 출력
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("⚠️ 파일 내용 읽기 오류:", err)
	}
	fmt.Println("============================\n")
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
	loadHistory()

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
	if len(history) > 0 {
		fmt.Println("파일에 기록을 저장 중...")
		err := saveHistory(history)
		if err != nil {
			fmt.Println("파일 저장 실패:", err)
		} else {
			fmt.Println("history.txt 파일에 저장 완료!")
		}
	}
}
