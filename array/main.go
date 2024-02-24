package main

import "fmt"

func main() {
	// ประกาศตัวแปร Array กำหนดค่าเริ่มต้น 4 ตำแหน่ง
	var fourNum [4]int
	fourNum[0] = 1
	fourNum[1] = 2

	fmt.Println(fourNum) // [1 2 0 0]
	// ประกาศตัวแปร Array ค่าว่าง
	var nums []int
	nums = make([]int, 4)
	nums[0] = 1
	nums[3] = 2
	fmt.Println(nums) //[1 0 0 2]

	var fiveNum [5]int
	fiveNum[2] = 20
	fiveNum[4] = 11
	var num1 []int
	num1 = fiveNum[2:5] // การ pointer splice ค่า array
	fmt.Println(num1)   // [20 0 11]

	var s []int // การ append ค่าให้กับ array
	s = append(s, 1, 2, 3, 4, 5)
	s = append(s, 99)
	s = append(s, 7)
	fmt.Println(s) // [1 2 3 4 5 99 7]

	fmt.Println(len(s))
	fmt.Println(cap(s))

	var s1 []int
	if s1 == nil {
		fmt.Println("it's nil")
	}

	a := [...]int{1, 2, 7, 3, 9, 11}
	s1 = a[:]
	fmt.Printf("%d %d %p %p\n", len(s1), cap(s1), s1, &a)
	s1 = append(s1, 12, 15)
	fmt.Printf("%d %d %p %p\n", len(s1), cap(s1), s1, &a)

	// Exercise
	s2 := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed

	s2 = s2[0:3]

	fmt.Println(s2)

	s3 := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	s3 = s3[4:7]
	fmt.Println(s3)

}
