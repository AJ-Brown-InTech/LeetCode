package twosum

import (
	"math/rand"
	"time"
	"errors"
)

// twoSum finds two numbers in nums that sum to target and returns their indices.
func twoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, value := range nums {
		comp := target - value
		if prevIndex, ok := result[comp]; ok {
			return []int{prevIndex, index}
		}
		result[value] = index
	}
	return nil
}

// generateRandomSlice creates a random slice of numbers between 0-9.
func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	slice := make([]int, size)

	for i := range slice {
		slice[i] = rand.Intn(10) // Random number between 0-9
	}

	return slice
}

// Setup generates a random slice and finds a two-sum solution for target = 7.
func Setup() ([]int, error) {
	randomNumberSlice := generateRandomSlice(10)
	result := twoSum(randomNumberSlice, 7)
	
	if result == nil {
		return nil, errors.New("no valid two-sum pair found")
	}

	return result, nil
}
