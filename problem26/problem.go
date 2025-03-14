package problem26

import "fmt"

func removeDuplicates(nums []int) int {
	found := make(map[int]int)
	idx := 0
  	for index, value := range nums {
		if _, ok := found[value]; !ok {
			found[value] = index
			nums[idx] = value
			idx++
		} 
	}	
	return len(found)	
}

func Implement() error {
 nums := []int{0,0,1,1,1,2,2,3,3,4}
 list := removeDuplicates(nums)
 fmt.Println(":)",list)
 return nil
}
