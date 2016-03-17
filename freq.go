package main

func Freq(col string, dataset []map[string]string) map[string]int {
	tabs := map[string]int{}
	for _, row := range dataset {
		tabs[row[col]] += 1
	}
	return tabs
}
