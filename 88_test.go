package main

import (
	"fmt"
	"testing"
)

func TestMerge88(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge88(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
