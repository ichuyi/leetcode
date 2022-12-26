package main

var letterMap = map[rune][]rune{
	'1': {},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	return recordCombinationsBack([]rune(digits), 0, nil, nil)
}

func recordCombinationsBack(input []rune, index int, tmp []rune, res []string) []string {
	if index == len(input) {
		if len(tmp) > 0 {
			res = append(res, string(tmp))
		}
		return res
	}
	for _, v := range letterMap[input[index]] {
		tmp = append(tmp, v)
		res = recordCombinationsBack(input, index+1, tmp, res)
	}
	return res
}
