package problem27

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index++
		}
	}
	return index	 
}

func Implement(){ 
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	length := removeElement(nums,val)
	fmt.Printf("Test: %d", length)
}
