package main

import (
	"fmt"
	"math/big"
	"slices"
)

func main() {
	fmt.Println("Task 1\n", task1(25))
	fmt.Println("Task 2\n", task2([]int{42, 12, 18}))
	fmt.Println("Task 3\n", task3(11, 20))
	fmt.Println("Task 4")
	task4(5)
}

func task1(num int) string {
	n := num % 100
	if n < 0 {
		n = -n
	}
	var str string;

	switch {
	case n <= 20:
		switch {
		case n == 1:
			str = "компьютер"
		case n >= 2 && n <= 4:
			str = "компьютера"
		default: 
			str = "компьютеров"
		}

	default:
		n %= 10
		switch {
		case n == 1:
			str = "компьютер"
		case n >= 2 && n <= 4:
			str = "компьютера"
		default: 
			str = "компьютеров"
		}
	}
	return fmt.Sprintf("%d %s", num, str)
}

func task2(array []int) []int {
	dividers := make([][]int, len(array))

	for index, elem := range array {
		for n := 1; n <= elem; n++ {
			if (elem % n == 0) {
				dividers[index] = append(dividers[index], n)
			}
		}
	}

	res := dividers[0]
	for i := 1; i < len(array); i++ {
		arr := []int{}
		for _, elem := range dividers[i] {
			if slices.Contains(res, elem) {
				arr = append(arr, elem)
			}
		}
		res = arr
	}

	return res
}

func task3(num1 int, num2 int) []int {
	res := []int{}
	n1 := int64(num1)
	n2 := int64(num2)

	for i := n1; i <= n2; i++ {
		if (big.NewInt(i).ProbablyPrime(1)) {
			res = append(res, int(i))
		}
	}

	return res
}

func task4(num int) {
	for i := 0; i <= num; i++ {
		for j := 0; j <= num; j++ {
			if (i == 0 && j == 0) {
				fmt.Print("\t")
			} else if (j == 0) {
				fmt.Printf("%d\t", i)
			} else if (i == 0) {
				fmt.Printf("%d\t", j)
			} else {
				fmt.Printf("%d\t", i * j)
			}
		}
		fmt.Println()
	}
}