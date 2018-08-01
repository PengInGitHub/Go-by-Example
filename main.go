package main

import (
	"flag"
	"fmt"
)

func main() {
	problems := Parse()
	for _, p := range problems {
		fmt.Printf("%+v\n", p)
	}
}

func playFlag() string {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question:answer")
	startDate := flag.String("start", "2018-07-20", "start date for analysis")
	testID := flag.Int("test", 3, "ID of test model")
	sslDisabled := flag.Bool("ssl", true, "if SSL is disabled")

	var flagVar int
	flag.IntVar(&flagVar, "referenceID", 1, "ID of reference ID")
	flag.Parse()
	return fmt.Sprintf("csv has name: %s, analysis start date: %s, ID of test: %d, SSL disabled: %t, reference: %d\n", *csvFileName, *startDate, *testID, *sslDisabled, flagVar)
}
