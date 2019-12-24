package main

func twoSum(nums []int, target int) []int {
	//result stores the numbers in nums as keys
	//and the index of the numbers as values
	result := make(map[int]int)

	for i, num := range nums {
		//checks to see if the key result[target - num] exists
		//if not then add it to the map
		//if it does grab the value of that key (idx) and then grab current index
		if idx, ok := result[target-num]; ok {
			return []int{idx, i}
		}
		result[num] = i
	}
	return []int{0, 0}

}
