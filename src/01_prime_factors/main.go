package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPrimeFactors(num int) []int {
	factors := []int{}
	divisor := 2
	for divisor <= num {
		if (num % divisor) == 0 {
			factors = append(factors, divisor)
			num = num / 2
		}
		divisor++
	}
	return factors
}

func main() {
	fmt.Println("Hello")
	reader := bufio.NewReader(os.Stdin)
	condition := true
	if condition {
		fmt.Println("Pass a number:")
		numInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: It isn't a number!\n", err)
		}
		anInt, err := strconv.Atoi(strings.TrimSpace(numInput))
		if err != nil {
			fmt.Println("Error: Fail to convert into intiger\n", err)
		} else {
			factorsList := getPrimeFactors(anInt)
			fmt.Println("Factors' list: ", factorsList)
		}
	}
}
