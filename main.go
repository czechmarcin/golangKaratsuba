package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt64(input string) int64 {
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Print(`"` + input + `" nie jest liczbą`)
		os.Exit(1)
	}
	return number
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Podaj pierwszą liczbę: ")
	input1, _ := reader.ReadString('\n')
	number1 := parseInt64(input1)
	fmt.Print("Podaj drugą liczbę: ")
	input2, _ := reader.ReadString('\n')
	number2 := parseInt64(input2)
	result := Karatsuba(number1, number2)
	fmt.Println(`Iloczyn podanych liczb wynosi `, result)
	fmt.Println(`Wciśnij enter aby wyłączyć program`)
	a, _ := reader.ReadString('\n')
	fmt.Print(a)
	os.Exit(1)
}
