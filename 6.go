package main

func convert(s string, numRows int) string {
	lines := make([][]rune, len(s))
	line := 0
	step := 1
	if numRows == 1 {
		step = 0
	}
	for _, v := range s {
		lines[line] = append(lines[line], v)
		line += step
		if line >= numRows || line < 0 {
			step = -step
			line += step * 2
		}
	}
	result := ""
	for _, v := range lines {
		result = result + string(v)
	}
	return result
}
func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	lines := make([][]rune, numRows)
	r := numRows*2 - 2
	m := numRows - 1
	for i, v := range s {
		line := i % r
		if line > m {
			line = m - (line - m)
		}
		lines[line] = append(lines[line], v)
	}
	ans := make([]rune, 0, len(s))
	for _, v := range lines {
		ans = append(ans, v...)
	}
	return string(ans)
}
