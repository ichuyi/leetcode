package main

func permute(nums []int) [][]int {
	return doPermute(nums, 0)
}
func doPermute(nums []int, i int) [][]int {
	if i == len(nums)-1 {
		return [][]int{{nums[i]}}
	}
	r := doPermute(nums, i+1)
	s := make([][]int, 0)
	for _, v := range r {
		for j := 0; j <= len(v); j++ {
			var t []int
			t = append(t, v[:j]...)
			t = append(t, nums[i])
			t = append(t, v[j:]...)
			s = append(s, t)
		}
	}
	return s
}
