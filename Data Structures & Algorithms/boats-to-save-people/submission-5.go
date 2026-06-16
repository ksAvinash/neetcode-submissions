func numRescueBoats(people []int, limit int) int {
	pos := map[int][]int{}
	for i, p := range people {
		val, ex := pos[p]
		if ex {
			pos[p] = append(val, i)
		} else {
			pos[p] = []int{i}
		}
	}
	// fmt.Println(pos)

	boats := 0
	for len(pos) > 0 {
		for weight, indices := range pos {
			// fmt.Println("->> checking", weight, indices)
			if weight >= limit {
				boats += len(indices)
				delete(pos, weight)
			} else {
				pos[weight] = pos[weight][:len(pos[weight])-1] // remove current person
				if len(pos[weight]) == 0 {
					delete(pos, weight)
				}
				rem := limit - weight
				for rem > 0 {
					// fmt.Println("rem", rem)
					_, ex := pos[rem]
					if ex {
						// fmt.Println("found rem", rem, pos[rem])
						pos[rem] = pos[rem][:len(pos[rem])-1] // remove matching rem person
						if len(pos[rem]) == 0 {
							delete(pos, rem)
						}
						break
					}
					rem--
				}
				boats++
			}
			// fmt.Println("boat full", boats, pos)
		}
	}
	// fmt.Println("final", boats, pos)
	return boats
}
