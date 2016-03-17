package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
)

func main() {

	data, _ := ioutil.ReadFile("data/spam.csv")

	r := csv.NewReader(bytes.NewReader(data))
	result, _ := r.ReadAll()

	// set of probabilities
	// ex. you draw 12 cards
	// from a deck and organize
	// them by suits
	// ps := []float64{
	// 	(5. / 12.), // clubs
	// 	(3. / 12.), // hearts
	// 	(3. / 12.), // spades
	// 	(1. / 12.), // diamonds
	// }

	D := hashify(result)

	// frequency table for specified column
	fmt.Println(Freq("class", D))

	fmt.Println(P("class", D))

	ps := P("class", D)

	// measure dataset heterogeniety
	fmt.Println(H(ps))

	// rem("class", D)

	fmt.Println(headers(D))
}
