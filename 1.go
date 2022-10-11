package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if vv, ok := m[target-v]; ok && i != vv {
			return []int{i, vv}
		}
	}
	return nil
}
