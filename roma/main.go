package main

func main() {

}

var roman = map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
var romanlist = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var quotient int
	var r []rune
	for _, i := range romanlist {
		quotient = num / i
		num = num % i
		for j := 0; j < quotient; j++ {
			r = append(r, []rune(roman[i])...)
		}

		if num == 0 {
			break
		}
	}

	return string(r)
}

var roman2int = map[rune]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}

func romanToInt(roman string) int {
	var sum, v int
	for _, i := range roman {
		if roman2int[i] > v {
			sum = sum + (roman2int[i] - v)
		} else {
			sum = sum + roman2int[i]
		}
	}
	return sum
}
