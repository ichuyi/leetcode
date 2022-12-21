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