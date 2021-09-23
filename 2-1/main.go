package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	df, err := os.Create("sample.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer func() {
		if err := df.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	cw := csv.NewWriter(df)
	cw.Write([]string{"aaaaaaa", "nnnnnnn", "bbbbbbbbbbbb"})
	cw.Flush()
}
