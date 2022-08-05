package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var output []int

	for _, current := range nums {
		complement := target - current
		for _, num := range output {
			if num == complement {
				output = []int{complement, current}
				return output
			}
		}
		output = append(output, current)
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
