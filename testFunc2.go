package main

import ."fmt"

type testInt func(int) bool //声明了一个函数类型

func isOld(i int) bool {
	if i % 2 == 0 {
		return true
	}
	return false
}
func isNew(i int) bool {
	if i % 2 == 0 {
		return false
	}
	return true
}
func filter(s []int, f testInt) []int {
	var result []int
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	slice := []int{1,2,3,4,5,6,7,8,9}
	Println("slice=", slice)
	odd := filter(slice, isOld)
	Println("Odd=", odd)
	new := filter(slice, isNew)
	Println("new=", new)
}