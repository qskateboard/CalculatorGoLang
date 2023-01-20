package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

var maxTable = []int{
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

var numInv = map[int]string{
	90: "XC",
	50: "L",
	40: "XL",
	10: "X",
	9:  "IX",
	5:  "V",
	4:  "IV",
	1:  "I",
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func convertToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func getFromRoman(str string) int {
	roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	val, ok := roman[str]
	if ok {
		return val
	}
	return 0
}

func main() {
	mathChars := []string{"+", "-", "*", "/"}
	buf := bufio.NewReader(os.Stdin)
	always := false
	for {
		fmt.Println("Input: ")
		input, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			text := strings.TrimSpace(string(input))
			chr := strings.Split(text, " ")
			if len(chr) == 3 && contains(mathChars, chr[1]) {
				romanChecker := 0
				firstNum, err := strconv.Atoi(chr[0])
				if err != nil {
					firstRoman := getFromRoman(chr[0])
					if firstRoman != 0 {
						romanChecker++
						firstNum = firstRoman
					} else {
						fmt.Println("На вход доступны только числа от 1 до 10 включительно")
						continue
					}
				}
				secondNum, err := strconv.Atoi(chr[2])
				if err != nil {
					secondRoman := getFromRoman(chr[2])
					if secondRoman != 0 {
						romanChecker++
						secondNum = secondRoman
					} else {
						fmt.Println("На вход доступны только числа от 1 до 10 включительно")
						continue
					}
				}
				if romanChecker == 0 || romanChecker == 2 {
					if firstNum < 1 || firstNum > 10 || secondNum < 1 || secondNum > 10 {
						fmt.Println("На вход доступны только числа от 1 до 10 включительно")
						continue
					}

					if romanChecker == 2 {
						firstNum = getFromRoman(chr[0])
						secondNum = getFromRoman(chr[2])
					}

					var result int = 0
					if chr[1] == "+" {
						result = firstNum + secondNum
					}
					if chr[1] == "-" {
						result = firstNum - secondNum
					}
					if chr[1] == "*" {
						result = firstNum * secondNum
					}
					if chr[1] == "/" {
						result = firstNum / secondNum
					}
					if romanChecker == 2 {
						if result < 1 {
							fmt.Println("В римской системе счисления нет отрицательных цифр")
						} else {
							fmt.Println(convertToRoman(result))
						}
					} else {
						fmt.Println(result)
					}

				} else {
					fmt.Println("Одновременно нельзя использовать разные системы счисления.")
					continue
				}

			} else {
				fmt.Println("Строка не является математической операцией.")
			}
		}
		if !always {
			break
		}
	}
}
