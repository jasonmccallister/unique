package main

import (
	"flag"
	"fmt"
	"os"

	unique "github.com/jasonmccallister/unique/pkg"
)

var (
	fileArg      string
	columnArg    int
	hasHeaderArg bool
	countOnlyArg bool
)

func main() {
	flag.StringVar(&fileArg, "file", "", "Name of the file to process (relative to the current directory")
	flag.IntVar(&columnArg, "column", 0, "Which column should be used for unique values (defaults to the first column which is 0)")
	flag.BoolVar(&hasHeaderArg, "header", true, "If the file has a header, setting this will bypass the first row of the CSV")

	flag.Parse()

	if fileArg == "" {
		fmt.Println("\nError: You must pass a file to process")
		fmt.Println("")
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(1)
	}

	file, err := os.Open(fileArg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count, err := unique.CSV(file, unique.Options{
		HasHeaders: hasHeaderArg,
		Column:     columnArg,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("there are a total of %v unique values for column %v", count, columnArg))
}
