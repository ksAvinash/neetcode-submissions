import (
	"maps"
	"slices"
)

func majorityElement(nums []int) []int {
	k := len(nums)/3
	counts := map[int]int{}
	res := map[int]bool{}
	
	for _, v := range nums {
		counts[v]++
		if counts[v] > k {
			res[v] = true
		}
	}
	return slices.Collect(maps.Keys(res))
}
