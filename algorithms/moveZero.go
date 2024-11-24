package main

import "fmt"

func main() {
	arr := []int{4, 0, 0, 7, 8, 0, 9}
	move_zeroes(arr)

	fmt.Println(arr)
}

//{4, 0, 0, 7, 8, 0, 9}
//{4, 7, 8, 0, 0, 0, 9}

func move_zeroes(nums []int) {
	zero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}
