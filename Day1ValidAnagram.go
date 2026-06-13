package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		count[ch]--

		if count[ch] < 0 {
			return false
		}
	}

	return true
}

/*
Two strings are anagrams if they contain the same characters with the same frequencies.

Approach:

Count the frequency of each character in the first string.
Decrease the frequency using the second string.
If all frequencies become zero, the strings are anagrams.

Time Complexity: O(n)
Space Complexity: O(n)

*/
