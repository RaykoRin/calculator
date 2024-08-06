package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func isRoman(numeral string) bool {
	_, exists := romanToArabic[numeral]
	return exists
}

func toArabic(roman string) (int, error) {
	if value, exists := romanToArabic[roman]; exists {
		return value, nil
	}
	return 0, errors.New("недопустимое римское число")
}

func toRoman(arabic int) (string, error) {
	if arabic < 1 {
		return "", errors.New("результат меньше единицы в римских цифрах недопустим")
	}
	roman := ""
	for arabic > 0 {
		for arabicValue, romanValue := range arabicToRoman {
			if arabic >= arabicValue {
				roman = romanValue
				arabic -= arabicValue
				break
			}
		}
	}
	return roman, nil
}

func calculate(expression string) (string, error) {
	expression = strings.TrimSpace(expression)

	var operator string
	if strings.Contains(expression, "+") {
		operator = "+"
	} else if strings.Contains(expression, "-") {
		operator = "-"
	} else if strings.Contains(expression, "*") {
		operator = "*"
	} else if strings.Contains(expression, "/") {
		operator = "/"
	} else {
		return "", errors.New("недопустимый оператор")
	}

	parts := strings.Split(expression, operator)
	if len(parts) != 2 {
		return "", errors.New("ожидается два операнда и один оператор")
	}

	operand1 := strings.TrimSpace(parts[0])
	operand2 := strings.TrimSpace(parts[1])

	var num1, num2 int
	var err error
	romanMode := false

	if isRoman(operand1) && isRoman(operand2) {
		num1, _ = toArabic(operand1)
		num2, _ = toArabic(operand2)
		romanMode = true
	} else if !isRoman(operand1) && !isRoman(operand2) {
		num1, err = strconv.Atoi(operand1)
		if err != nil {
			return "", errors.New("недопустимое арабское число")
		}
		num2, err = strconv.Atoi(operand2)
		if err != nil {
			return "", errors.New("недопустимое арабское число")
		}
	} else {
		return "", errors.New("используются одновременно разные системы счисления")
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("деление на ноль недопустимо")
		}
		result = num1 / num2
	default:
		return "", errors.New("недопустимый оператор")
	}

	if romanMode {
		if result < 1 {
			return "", errors.New("результат меньше единицы в римских цифрах недопустим")
		}
		romanResult, err := toRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	}

	return strconv.Itoa(result), nil
}

func main() {
	var input string
	fmt.Println("Введите выражение:")
	fmt.Scanln(&input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
