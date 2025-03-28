package problem11

import "fmt"

func maxArea(height []int) int {
	
	l, r := 0, len(height) - 1
	ma := 0

	for l  <  r {
		var h int 
		if height[l] > height[r] {
			h = height[r]
		} else {
			h  = height[l]
		}

		width := r - l
		area := width * h

		if area > ma {
			ma  = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		
	}
	
	return ma
}

func Implement() error {
 nums := []int{1,8,6,2,5,4,8,3,7}
 list := maxArea(nums)
 fmt.Println(":)",list)
 return nil
}
