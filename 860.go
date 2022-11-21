package main

func lemonadeChange(bills []int) bool {
	r := make(map[int]int)
	for _, v := range bills {
		switch v {
		case 5:
			r[v]++
		case 10:
			if r[5] > 0 {
				r[5]--
				r[10]++
			} else {
				return false
			}
		case 20:
			if r[10] > 0 && r[5] > 0 {
				r[10]--
				r[5]--
				r[20]++
			} else if r[5] > 2 {
				r[5] -= 3
				r[20]++
			} else {
				return false
			}
		}
	}
	return true
}
