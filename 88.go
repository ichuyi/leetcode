package main

func merge88(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	for index := m + n - 1; index >= 0; index-- {
		var t int
		if index1 < 0 {
			t = nums2[index2]
			index2--
		} else if index2 < 0 {
			t = nums1[index1]
			index1--
		} else {
			t = nums1[index1]
			if nums2[index2] > t {
				t = nums2[index2]
				index2--
			} else {
				index1--
			}
		}
		nums1[index] = t
	}
}
