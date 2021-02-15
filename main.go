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

	var fracText []string

	// since as per requirement multiplyFractions function needs to take two fractions as input that's why the array length is 2
	// in case we want to increase the array number or if we want to take input from users how many fractions they want to multiply
	// that can be customized by making the array length as per user input and passing the appropriate parameters to multiplyFractions
	for i := 0; i < 2; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		num, _ := reader.ReadString('\n')
		num = strings.Replace(num, "\n", "", -1)
		fracText = append(fracText, num)
	}

	res := multiplyFractions(fracText[0], fracText[1])
	fmt.Println("Result is : ", res)
}

// multiplyFractions takes two fractions and multiply them.
func multiplyFractions(fractions ...string) string {
	var (
		numArr []int
		denArr []int
	)

	numMul := 1
	denMul := 1

	if len(fractions) > 0 {
		for i := 0; i < len(fractions); i++ {
			getFractionArrays(fractions[i], &numArr, &denArr)
		}
	}

	for i := 0; i < len(numArr); i++ {
		numMul = numMul * numArr[i]
	}

	for i := 0; i < len(denArr); i++ {
		denMul = denMul * denArr[i]
	}

	numMul, denMul = reduce(numMul, denMul)

	n1 := strconv.Itoa(numMul)
	d1 := strconv.Itoa(denMul)

	return n1 + "/" + d1
}

// getFractionArrays gets the fraction in string format, splits into numerator and denominator and pushes to appropriate array
func getFractionArrays(fraction string, numArr *[]int, denArr *[]int) {
	fracArr := strings.Split(fraction, "/")
	num, numErr := strconv.Atoi(fracArr[0])
	if numErr == nil {
		*numArr = append(*numArr, num)
	}

	den, denErr := strconv.Atoi(fracArr[1])
	if denErr == nil {
		*denArr = append(*denArr, den)
	}
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
func reduce(nums ...int) (int, int) {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = gcd(nums[i], ans)
	}

	nums[0] = nums[0] / ans
	nums[1] = nums[1] / ans

	return nums[0], nums[1]
}
