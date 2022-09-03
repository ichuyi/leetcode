package main

import "testing"

func TestMyQueue(t *testing.T) {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	m.Push(3)
	m.Push(4)
	m.Pop()
	m.Push(5)
	m.Pop()
	m.Pop()
	m.Pop()
	m.Pop()
}
