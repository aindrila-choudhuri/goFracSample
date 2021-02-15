package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many numbers do you want to enter -> ")

	var fracText []string

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	numLen, err := strconv.Atoi(text)

	if err == nil {
		for i := 0; i < numLen; i++ {
			fmt.Print("-> ")
			num, _ := reader.ReadString('\n')
			num = strings.Replace(num, "\n", "", -1)
			fracText = append(fracText, num)
		}
	}

	numArr := 1
	denArr := 1

	for i := 0; i < len(fracText); i++ {
		fracArr := strings.Split(fracText[i], "/")

		num, numErr := strconv.Atoi(fracArr[0])
		if numErr == nil {
			numArr = numArr * num
		}

		den, denErr := strconv.Atoi(fracArr[1])
		if denErr == nil {
			denArr = denArr * den
		}
	}

	fmt.Println("numArr : ", numArr)
	fmt.Println("denArr : ", denArr)

	hcf := HCF([]int{numArr, denArr}, 2)
	fmt.Println("hcf : ", hcf)

	numArr = numArr / hcf
	denArr = denArr / hcf

	n1 := strconv.Itoa(numArr)
	d1 := strconv.Itoa(denArr)

	fmt.Println("res : " + n1 + "/" + d1)

}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}
	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

// HCF of Denominator
func HCF(den []int, N int) int {
	ans := den[0]
	for i := 1; i < N; i++ {
		ans = gcd(den[i], ans)
	}
	return ans
}
