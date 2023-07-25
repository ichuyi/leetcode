package main

func generateParenthesis(n int) []string {
	return generateParenthesisBack(nil, 0, n, nil)
}
func generateParenthesisBack(stack []rune, offset, push int, res []string) []string {
	if push == 0 && offset == 0 {
		res = append(res, string(stack))
		return res
	}
	if push > 0 {
		stack = append(stack, '(')
		res = generateParenthesisBack(stack, offset+1, push-1, res)
		stack = stack[:len(stack)-1]
	}
	if offset > 0 {
		stack = append(stack, ')')
		res = generateParenthesisBack(stack, offset-1, push, res)
		stack = stack[:len(stack)-1]
	}
	return res
}

func generateParenthesis2(n int) []string {
	return doGenerateParenthesis2(nil, 0, n, nil)
}
func doGenerateParenthesis2(stack []rune, offset, push int, res []string) []string {
	if push == 0 && offset == 0 {
		tmp := string(stack)
		res = append(res, tmp)
		return res
	}
	if push > 0 {
		stack = append(stack, '(')
		res = doGenerateParenthesis2(stack, offset+1, push-1, res)
		stack = stack[:len(stack)-1]
	}
	if offset > 0 {
		stack = append(stack, ')')
		res = doGenerateParenthesis2(stack, offset-1, push, res)
	}
	return res
}
