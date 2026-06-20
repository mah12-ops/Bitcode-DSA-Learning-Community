// package main

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

// //   Complexity :  Brute Force (O(n²))

// func twoSumOptimal(nums []int, target int) []int {
// 	seen := make(map[int]int)

// 	for i, num := range nums {
// 		complement := target - num

// 		if idx, exists := seen[complement]; exists {
// 			return []int{idx, i}
// 		}

// 		seen[num] = i
// 	}

// 	return nil
// }

// // optimal solution using HashMap (O(n))
