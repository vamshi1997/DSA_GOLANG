package main

import "fmt"

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	var result []int

	for index, value := range nums {

		if temp, ok := record[target-value]; ok {

			result = append(result, index)
			result = append(result, temp)
			break
		}

		record[value] = index
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}
