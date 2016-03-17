package main

import (
	"fmt"
	"math"
)

// Shannon's entropy
func H(ps interface{}) float64 {
	switch t := ps.(type) {
	case map[string]float64:
		total := 0.0
		for _, p := range t {
			p_log2 := math.Log2(p)
			total += p * p_log2
		}
		return -1.0 * total
	case []float64:
		total := 0.0
		for _, p := range t {
			p_log2 := math.Log2(p)
			total += p * p_log2
		}
		return -1.0 * total
	default:
		return 0.0
	}
}

func rem(target string, dataset []map[string]string) {

	// for each feature that are not the target feature

	for _, hs := range headers(dataset) {
		if hs != target {
			for _, p := range P(hs, dataset) {
				h := H(ps)

			}

		}
	}

	//		ps:=P(feature)
	//		h:= H(ps)
	//		ps * h
	//		add total
	//		return -1.0 * total

	fmt.Println(s)
}

func IG() {

}

// s := map[string]map[string]int{}
// // for each feature that are not the target feature
// for _, row := range dataset {
// 	for k, v := range row {
// 		if _, in := s[k]; in {
// 			s[k][v]++
// 		} else {
// 			s[k] = map[string]int{v: 1}
// 		}
// 	}
// }
