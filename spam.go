package main

import "fmt"

func spamScan1(dataset interface{}) {
	switch t := dataset.(type) {
	case [][]string:
		h := hashify(t)
		spamScan1(h)
	case []map[string]string:
		fmt.Println("\nspamScan1 -------------")
		for _, row := range t {
			if row["contains_images"] == "true" {
				if row["suspicious_words"] == "true" {
					fmt.Printf("\t%s is SPAM\n", row["id"])
				} else {
					fmt.Printf("\t%s is HAM\n", row["id"])
				}
			} else {
				if row["unknown_sender"] == "true" {
					fmt.Printf("\t%s is SPAM\n", row["id"])
				} else {
					fmt.Printf("\t%s is HAM\n", row["id"])
				}
			}
		}
	}
}

func spamScan2(dataset interface{}) {
	switch t := dataset.(type) {
	case [][]string:
		h := hashify(t)
		spamScan2(h)
	case []map[string]string:
		fmt.Println("\nspamScan2 -------------")
		for _, row := range t {
			if row["suspicious_words"] == "true" {
				fmt.Printf("\t%s is SPAM\n", row["id"])
			} else {
				fmt.Printf("\t%s is HAM\n", row["id"])
			}
		}
	}
}
