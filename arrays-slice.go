package main

import "fmt"

func main() {
	var numbers = []int{1, 1, 2, 3, 5, 8, 9}
	numbers2 := [7]int{2, 2, 3, 3, 4, 5, 6}
	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(numbers[2])

	for idx := 0; idx < len(numbers); idx++ {
		fmt.Println(numbers[idx])
	}
	fmt.Println(len(numbers))

	// Slice
	fmt.Println(numbers[2:])
	fmt.Println(numbers[:len(numbers)])
	fmt.Println(numbers[:])

	var names = [5]string {
		"Ana",
		"Jose",
		"Maria",
	}
	var p1 []string = names[0:2]
	fmt.Println(p1)
	p1[0] = "Rogerio"
	fmt.Println(p1)
	fmt.Println("Original:", names)

	fmt.Printf("\n%T\n", names)
	fmt.Printf("\ncap=%d, len=%d\n", cap(p1), len(p1))

	nums := make([]int, 0, 5)
	fmt.Println(nums[0:2])
	fmt.Println(nums)
	var nums2 = append(nums, 3)
	fmt.Println("Nums:", nums)
	fmt.Println("Nums2:", nums2)

}