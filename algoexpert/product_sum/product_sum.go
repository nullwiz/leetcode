package main

import "fmt"

type SpecialArray []interface{}

func ProductSum(array SpecialArray, depth int) int {
	product := 0

	for _, elem := range array {
		if subList, ok := elem.(SpecialArray); ok {
			fmt.Printf("recursively calling Product sum with depth:%d, product:%d\n", depth, product)
			product += ProductSum(subList, depth+1)
		} else {
			if number, ok := elem.(int); ok {
				fmt.Printf("elem: %d, product:%d, depth:%d\n", elem, product, depth)
				product += number
			}
		}
	}
	fmt.Printf("depth:%d , curr product:%d\n", depth, product)
	return depth * product
}

func main() {
	arr := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	fmt.Println(ProductSum(arr, 1))
}
