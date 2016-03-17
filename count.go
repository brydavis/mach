package main

func Count(dataset []map[string]string) map[string]int {
	cnts := map[string]int{}
	for _, row := range dataset {
		for k, _ := range row {
			cnts[k] += 1
		}
	}
	return cnts
}
