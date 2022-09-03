package main

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	stack := make([]int, 0)
	offset := -1
	pushedOffset, poppedOffset := 0, 0
	for pushedOffset, poppedOffset = 0, 0; poppedOffset < len(popped); {
		if offset < 0 {
			stack = append(stack, pushed[pushedOffset])
			offset++
			pushedOffset++
			continue
		}
		if stack[offset] == popped[poppedOffset] {
			stack = stack[0:offset]
			offset--
			poppedOffset++
			continue
		}
		if pushedOffset >= len(pushed) {
			return false
		}
		stack = append(stack, pushed[pushedOffset])
		offset++
		pushedOffset++
	}
	return poppedOffset == len(popped) && pushedOffset == len(pushed) && offset == -1
}
