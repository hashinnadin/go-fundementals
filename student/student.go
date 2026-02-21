package student

import "fmt"

func Dupilicatesstored(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
		} else {
			result = append(result, value)
		}
	}
	return result
}

func main() {

	num := Dupilicatesstored([]int{1, 2, 3, 4, 5, 6, 7, 7, 4, 3, 0, 2})
	fmt.Println(num)
}
