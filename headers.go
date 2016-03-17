package main

func headers(dataset []map[string]string) (h []string) {
	m := map[string]bool{}

	for _, row := range dataset {
		for k, _ := range row {
			m[k] = true
		}
	}

	for k, _ := range m {
		h = append(h, k)
	}

	return
}
