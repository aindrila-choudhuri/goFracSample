package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	infiniteString  = "infinite"
	undefinedString = "undefined"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	// as per requirement 2 numbers need to multiplied but for more than two numbers also this program can be used
	fmt.Print("How many numbers do you want to enter -> ")

	// fracText will hold the input fractions in string format
	var fracText []string

	text, _ := reader.ReadString('\n')
	text = strings.Trim(strings.Replace(text, "\n", "", -1), " ")

	numLen, err := strconv.Atoi(strings.Trim(text, " "))
	if err != nil {
		log.Fatal(err)
	} else {
		for i := 0; i < numLen; i++ {
			fmt.Printf("Enter number %d in fraction format: ", i+1)
			num, _ := reader.ReadString('\n')
			num = strings.Trim(strings.Replace(num, "\n", "", -1), " ")
			if !isValidInput(num) {
				return
			}
			if len(num) > 0 {
				fracText = append(fracText, num)
			}
		}
	}
	res := multiplyFractions(fracText...)
	fmt.Println("Result is : ", res)
}

// isValidInput checks whether / is in correct place
func isValidInput(numText string) bool {
	if strings.Index(numText, "/") < 1 || strings.Count(numText, "/") > 1 {
		fmt.Println("Oops! You need to enter the numbers in fraction format")
		return false
	}

	return true
}

// multiplyFractions takes fractions in string format and multiply them.
func multiplyFractions(fractions ...string) string {
	numMul := 1
	denMul := 1

	if len(fractions) > 0 {
		for i := 0; i < len(fractions); i++ {
			getFractionMultiplicationResult(fractions[i], &numMul, &denMul)
		}
	}

	// wg := &sync.WaitGroup{}
	// if len(fractions) > 0 {
	// 	wg.Add(len(fractions))
	// 	for i := 0; i < len(fractions); i++ {
	// 		go func(fraction string, numMul *int, denMul *int) {
	// 			getFractionMultiplicationResult(fraction, numMul, denMul)
	// 			wg.Done()
	// 		}(fractions[i], &numMul, &denMul)
	// 	}
	// 	wg.Wait()
	// }

	numMul, denMul = reduce(numMul, denMul)

	res := ""

	if numMul == 0 {
		res = infiniteString
	} else if denMul == 0 {
		res = undefinedString
	} else {
		n1 := strconv.Itoa(numMul)
		d1 := strconv.Itoa(denMul)
		res = n1 + "/" + d1
	}

	return res
}

// getFractionMultiplicationResult splits each fraction into numerator and denominator and multiplies
func getFractionMultiplicationResult(fraction string, numMul, denMul *int) {
	fracArr := strings.Split(fraction, "/")
	num, numErr := strconv.Atoi(fracArr[0])
	if numErr != nil {
		log.Fatal(numErr)
	} else {
		*numMul = *numMul * num
	}

	den, denErr := strconv.Atoi(fracArr[1])
	if denErr != nil {
		log.Fatal(denErr)
	} else {
		*denMul = *denMul * den
	}
}

// gcd will determine the greatest common divisor
func gcd(a, b int) int {
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

// reduce reduces the fraction to its lowest form
func reduce(nums ...int) (int, int) {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = gcd(nums[i], ans)
	}

	nums[0] = nums[0] / ans
	nums[1] = nums[1] / ans

	return nums[0], nums[1]
}
