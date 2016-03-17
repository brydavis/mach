package main

import "strings"

func hashify(dataset [][]string) []map[string]string {
	headers := []string{}
	for _, d := range dataset[0] {
		headers = append(headers, strings.TrimSpace(d))
	}

	hash := []map[string]string{}

	for _, row := range dataset[1:] {
		m := map[string]string{}
		for n, val := range row {
			m[headers[n]] = strings.TrimSpace(val)
		}
		hash = append(hash, m)
	}

	return hash
}
