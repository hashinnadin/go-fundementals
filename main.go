package main

import "fmt"

func Dupilicatesstored(arr []int) ([]int, error) {

	if len(arr) <= 0 {
		return arr, fmt.Errorf("Not value in Your Array")
	}
	seen := make(map[int]bool)
	result := []int{}
	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
		} else {
			result = append(result, value)
		}
	}
	return result, nil
}

func main() {

	arr, err := Dupilicatesstored([]int{1, 2, 3, 2, 4, 1, 5, 8, 5, 7})
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Result :", arr)

	}
}
