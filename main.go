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
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15,
	"XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25,
	"XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
	"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35,
	"XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
	"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45,
	"XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
	"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55,
	"LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
	"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65,
	"LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
	"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75,
	"LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
	"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85,
	"LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
}

var arabicToRoman = []struct {
	value  int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
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
	for _, pair := range arabicToRoman {
		for arabic >= pair.value {
			roman += pair.symbol
			arabic -= pair.value
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

	if num1 == 0 || num2 == 0 {
		return "", errors.New("операнд не может быть равен нулю")
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
