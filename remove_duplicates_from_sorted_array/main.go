package main

import "fmt"

// Two pointer approach, one for the index/value in question (currentPointer), and another to iterate through the list. Iterate through the list until index1 value != index2. When index1 value == index2 value, increment index1 by 1, and replace with index2 value. This will replace the duplicate value.

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i int

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {

	testCase := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expectedArray := []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}
	expectedLength := 5

	result := removeDuplicates(testCase)

	success := true
	if result != expectedLength {
		success = false
	} else {
		for index, val := range testCase {
			if val != expectedArray[index] {
				success = false
			}
		}
	}

	if success {
		fmt.Println("SUCCEEDED")
	} else {
		fmt.Println("FAILED")
	}

}
