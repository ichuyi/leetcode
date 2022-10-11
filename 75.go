package main

func sortColors(nums []int) {
	red, white, blue := -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
			white++
			blue++
			insert(nums, red, i)
		} else if nums[i] == 1 {
			white++
			blue++
			insert(nums, white, i)
		} else {
			blue++
			insert(nums, blue, i)
		}
	}
}
func insert(input []int, m, n int) {
	value := input[n]
	for i := n; i > m; i-- {
		input[i] = input[i-1]
	}
	input[m] = value
}
