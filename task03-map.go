package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	n := make([]int, len(input))
	var i int
	for k, _ := range input {
		n[i] = k
		i++
	}
	sort.Ints(n)
	for i := 0; i < len(input); i++ {
		result = append(result, input[n[i]])
	}

	return
}
