package dsa



// GreaterElementsToRight function finds the next greater element to the right for each element in the array
func GreaterElementsToRight(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	stack := []int{}

	// Traverse the array from right to left
	for i := n - 1; i >= 0; i-- {
		// Pop elements from the stack while they are less than or equal to the current element
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		// If stack is empty, there is no greater element to the right
		if len(stack) == 0 {
			result[i] = -1
		} else {
			// The top element of the stack is the next greater element
			result[i] = stack[len(stack)-1]
		}

		// Push the current element onto the stack
		stack = append(stack, arr[i])
	}

	return result
}