package main

import "fmt"

//Given a non-empty array of digits representing a non-negative integer, plus one to the integer. The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit. You may assume the integer does not contain any leading zero, except the number 0 itself.

//Input: [1,2,3]
//Output: [1,2,4]
//Explanation: The array represents the integer 123.

func plusOne(digits []int) []int {

	size := len(digits)
	carryOver := false

	for i := size - 1; i >= 0; i-- {
		if digits[i]+1 >= 10 {
			carryOver = true
			digits[i] = 0
		} else {
			carryOver = false
			digits[i] = digits[i] + 1
			break
		}
	}

	if carryOver {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {

	testCase1 := []int{8, 9}
	expectedAnswer1 := []int{9, 0}
	testHelper(testCase1, expectedAnswer1)

	testCase2 := []int{9, 9}
	expectedAnswer2 := []int{1, 0, 0}
	testHelper(testCase2, expectedAnswer2)

	testCase3 := []int{9}
	expectedAnswer3 := []int{1, 0}
	testHelper(testCase3, expectedAnswer3)
}

func testHelper(testCase []int, answer []int) {
	testCaseCopy := make([]int, len(testCase))
	copy(testCaseCopy, testCase)

	result := plusOne(testCase)
	success := true

	fmt.Print(testCaseCopy, " + 1 = ", result, " Test: ")
	for index, val := range answer {
		if val != result[index] {
			success = false
		}
	}
	if success {
		fmt.Println("SUCCEEDED")
	} else {
		fmt.Println("FAILED")
	}

}
