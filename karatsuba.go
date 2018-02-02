package main

import "math"

func getDigitCount(num int64) uint {
	var digitCount uint
	if num == 0 {
		return 1
	}
	if num < 0 {
		num = -num
	}
	for num > 0 {
		digitCount++
		num = num / 10
	}
	return digitCount
}

func splitNumber(num int64, digits uint) (int64, int64) {
	divisor := int64(math.Pow(10, float64(digits)))

	if num >= divisor {
		return num / divisor, num % divisor
	}
	return 0, num
}

// Zwraca iloczyn parametrów inputX, inputY
func Karatsuba(inputX int64, inputY int64) int64 {

	isInputSigned := (inputX > 0 && inputY < 0) || (inputX < 0 && inputY > 0)
	var x int64
	var y int64

	if inputX < 0 {
		x = -inputX
	} else {
		x = inputX
	}
	if inputY < 0 {
		y = -inputY
	} else {
		y = inputY
	}

	if x < 10 || y < 10 {
		return x * y
	}

	xDigits := getDigitCount(x)
	yDigits := getDigitCount(y)
	var maxDigits uint
	if xDigits >= yDigits {
		maxDigits = xDigits / 2
	} else {
		maxDigits = yDigits / 2
	}
	xHigh, xLow := splitNumber(x, maxDigits)
	yHigh, yLow := splitNumber(y, maxDigits)

	z1 := Karatsuba(xLow, yLow)
	z2 := Karatsuba(xLow+xHigh, yLow+yHigh)
	z3 := Karatsuba(xHigh, yHigh)

	result := (z3 * int64(math.Pow(10, float64(2*maxDigits)))) + (z2-z3-z1)*int64(math.Pow(10, float64(maxDigits))) + z1
	if isInputSigned {
		result *= -1
	}
	return result

}
