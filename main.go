package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	// as per requirement 2 numbers need to multiplied but for more than two numbers also this program can be used
	fmt.Print("How many numbers do you want to enter -> ")

	// fracText will hold the input fractions in string format
	var fracText []string

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	numLen, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	} else {
		for i := 0; i < numLen; i++ {
			fmt.Printf("Enter number %d in fraction format: ", i+1)
			num, _ := reader.ReadString('\n')
			num = strings.Replace(num, "\n", "", -1)
			if len(num) > 0 {
				fracText = append(fracText, num)
			}
		}
	}
	res := multiplyFractions(fracText...)
	fmt.Println("Result is : ", res)
}

// multiplyFractions takes fractions in string format and multiply them.
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

// getFractionArrays receives the fraction in string format, splits into numerator and denominator, pushes to appropriate arrays and returns
func getFractionArrays(fraction string, numArr *[]int, denArr *[]int) {
	if strings.Index(fraction, "/") > -1 {
		fracArr := strings.Split(fraction, "/")
		num, numErr := strconv.Atoi(fracArr[0])
		if numErr != nil {
			log.Fatal(numErr)
		} else {
			*numArr = append(*numArr, num)
		}

		den, denErr := strconv.Atoi(fracArr[1])
		if denErr != nil {
			log.Fatal(denErr)
		} else {
			*denArr = append(*denArr, den)
		}
	} else {
		log.Fatal("Oops! You need to enter the numbers in fraction format")
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
