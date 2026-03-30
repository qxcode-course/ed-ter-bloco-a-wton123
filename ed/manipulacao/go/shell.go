func getMen(vet []int) []int {
	var res []int
	for _, x := range vet {
		if x > 0 {
			res = append(res, x)
		}
	}
	return res
}

func getCalmWomen(vet []int) []int {
	var res []int
	for _, x := range vet {
		if x < 0 && x > -10 {
			res = append(res, x)
		}
	}
	return res
}

func sortVet(vet []int) []int {
	res := make([]int, len(vet))
	copy(res, vet)

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}

func sortStress(vet []int) []int {
	res := make([]int, len(vet))
	copy(res, vet)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if abs(res[i]) > abs(res[j]) {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}

func reverse(vet []int) []int {
	res := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		res[i] = vet[len(vet)-1-i]
	}
	return res
}

func unique(vet []int) []int {
	seen := make(map[int]bool)
	var res []int

	for _, x := range vet {
		if !seen[x] {
			res = append(res, x)
			seen[x] = true
		}
	}
	return res
}

func repeated(vet []int) []int {
	count := make(map[int]int)
	for _, x := range vet {
		count[x]++
	}

	var res []int

	for _, x := range vet {
		if count[x] > 1 {
			res = append(res, x)
		}
	}
	return res
}



