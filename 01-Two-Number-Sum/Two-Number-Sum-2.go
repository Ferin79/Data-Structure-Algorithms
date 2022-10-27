package main

func TwoNumberSum(array []int, target int) []int {
	
	m := map[int]bool{}
	
	for _, num := range array{
		match := target - num
		
		if _, ok := m[match]; ok {
			return []int{match, num}
		}
		m[num] = true
	}
	
	return []int{}
}