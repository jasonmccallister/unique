package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	fileArg      string
	columnArg    int
	hasHeaderArg bool
	countOnlyArg bool
)

func main() {
	flag.StringVar(&fileArg, "file", "", "Name of the file to process (relative to the current directory")
	flag.IntVar(&columnArg, "column", 0, "Which column should be used for unique values")
	flag.BoolVar(&hasHeaderArg, "header", true, "If the file has a header, setting this will bypass the first row of the CSV")
	flag.Parse()

	if fileArg == "" {
		fmt.Println("\nError: You must pass a file to process")
		fmt.Println("")
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(1)
	}

	data := make(map[string]int)
	csvFile, err := open(fileArg)
	handleError(err)

	rdr := csv.NewReader(csvFile)
	for {
		line, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(line)
			log.Println(err.Error())
			continue
		}

		if line != nil {
			col := line[0]
			if d, ok := data[col]; ok {
				data[col] = d + 1
			} else {
				data[col] = 1
			}
		}
	}

	log.Println(fmt.Sprintf("there are a total of %v unique values for column %v", len(data), columnArg))
}

func open(file string) (*bufio.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(f), nil
}

func handleError(e error) {
	if e != nil {
		log.Println(e.Error())
		os.Exit(1)
	}
}
