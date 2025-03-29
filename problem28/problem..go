package problem28

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func Implement(){ 
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Printf("T1: %d",strStr(haystack,needle))
}	
 
