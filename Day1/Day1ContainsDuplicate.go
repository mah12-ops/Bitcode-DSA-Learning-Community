package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}

		seen[num] = struct{}{}
	}

	return false
}

/*


A HashSet stores only unique values and provides O(1) average-time lookup.

Algorithm:

Traverse the array.
Check if the current number already exists in the HashSet.
If it exists, a duplicate is found → return true.
Otherwise, add the number to the HashSet.
If the traversal finishes, return false.

Time Complexity: O(n)
Space Complexity: O(n)




*/
