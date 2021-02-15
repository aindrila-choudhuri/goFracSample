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

	// fracText will hold the input fractions in string format
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

	var (
		numArr []int
		denArr []int
	)

	for i := 0; i < len(fracText); i++ {
		fracArr := strings.Split(fracText[i], "/")

		num, numErr := strconv.Atoi(fracArr[0])
		if numErr == nil {
			numArr = append(numArr, num)
		}

		den, denErr := strconv.Atoi(fracArr[1])
		if denErr == nil {
			denArr = append(denArr, den)
		}
	}

	res := multiplyFractions(numArr, denArr)
	fmt.Println("Result is : ", res)
}

func multiplyFractions(numArr []int, denArr []int) string {
	numMul := 1
	denMul := 1

	for i := 0; i < len(numArr); i++ {
		numMul = numMul * numArr[i]
	}

	for i := 0; i < len(denArr); i++ {
		denMul = denMul * denArr[i]
	}

	hcf := reduce([]int{numMul, denMul}, 2)
	fmt.Println("hcf : ", hcf)

	numMul = numMul / hcf
	denMul = denMul / hcf

	n1 := strconv.Itoa(numMul)
	d1 := strconv.Itoa(denMul)

	return n1 + "/" + d1
}

// gcd will determine the greatest common divisor
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

// reduce will determine the greatest common factor (GCF)
func reduce(den []int, N int) int {
	ans := den[0]
	for i := 1; i < N; i++ {
		ans = gcd(den[i], ans)
	}
	return ans
}
