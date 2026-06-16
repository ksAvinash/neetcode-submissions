import "slices"

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	boats := 0
	i, j := 0, len(people)-1
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		boats++
	}
	return boats
}
