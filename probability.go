package main

func P(col string, dataset []map[string]string) map[string]float64 {
	ps := map[string]float64{}
	tabs := Freq(col, dataset)
	denom := Count(dataset)[col]

	for k, numer := range tabs {
		ps[k] = float64(numer) / float64(denom)
	}

	return ps
}
