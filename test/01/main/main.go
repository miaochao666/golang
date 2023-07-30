package main

import "fmt"

func main() {
	// 初始化0个元素，容量为10的切片
	arr1 := make([]any, 0, 10)
	fmt.Println(arr1)
	arr2 := []int{1, 2, 3, 4}
	//arr2 = remove1(arr2, 1)
	arr2 = remove2(arr2, 1)
	fmt.Println(arr2)
}

func remove1[T AnyType](arr []T, index int) []T {
	if index < 0 || len(arr) <= index {
		return arr
	}
	newarr1 := arr[:index]
	newarr2 := arr[index+1 : len(arr)]
	newarr3 := append(newarr1, newarr2...)
	return newarr3
}

func remove2[T AnyType](arr []T, index int) []T {
	result := make([]T, 0)
	for i, val := range arr {
		if i != index {
			result = append(result, val)
		}
	}
	return result
}

type AnyType interface {
	int | string
}
