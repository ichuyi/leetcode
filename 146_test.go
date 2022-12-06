package main

import (
	"fmt"
	"testing"
)

func TestLRUCache1(t *testing.T) {
	c := ConstructorLRU(2)
	c.Put(1, 0)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}
func TestLRUCache2(t *testing.T) {
	c := ConstructorLRU(1)
	c.Put(2, 1)
	fmt.Println(c.Get(2))
	c.Put(3, 2)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
}
func TestLRUCache3(t *testing.T) {
	c := ConstructorLRU(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
func TestLRUCache4(t *testing.T) {
	c := ConstructorLRU(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}
