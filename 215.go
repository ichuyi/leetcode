package main

func findKthLargest(nums []int, k int) int {
	nums, leftIndex, midIndex := splitPart(nums, getPivot(nums))
	if len(nums)-midIndex < k && len(nums)-leftIndex >= k {
		return nums[leftIndex]
	}
	if len(nums)-midIndex >= k {
		return findKthLargest(nums[midIndex:], k)
	}
	return findKthLargest(nums[:leftIndex], leftIndex+k-len(nums))

}
func getPivot(nums []int) int {
	return nums[len(nums)/2]
}
func splitPart(nums []int, pivot int) ([]int, int, int) {
	left, mid, right := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			mid = append(mid, v)
		} else {
			right = append(right, v)
		}
	}
	leftIndex := len(left)
	midIndex := leftIndex + len(mid)
	left = append(left, mid...)
	left = append(left, right...)
	return left, leftIndex, midIndex
}

// insertNums TODO
// func splitPart(nums []int, pivot int) (int, int) {
// 	leftIndex := 0
// 	midIndex := 0
// 	for i, v := range nums {
// 		if v > pivot {
// 			continue
// 		}
// 		if v == pivot {
// 			insertNums(nums, midIndex, i)
// 			midIndex++
// 		} else {
// 			insertNums(nums, leftIndex, i)
// 			leftIndex++
// 			midIndex++
// 		}
// 	}
// 	return leftIndex, midIndex
// }
func insertNums(nums []int, i, j int) {
	t := nums[j]
	for j > i {
		nums[j] = nums[j-1]
		j--
	}
	nums[i] = t
}
