package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Sheet is ...
type Sheet struct {
	rows     *[]Row
	filename string
}

func newSheet(filename string) Sheet {
	s := Sheet{
		rows:     &[]Row{},
		filename: filename,
	}
	return s
}

func (s Sheet) addRow(r Row) {
	*s.rows = append(*s.rows, r)
}

// lazy loading rows
func (s Sheet) load() {

	csvfile, err := os.Open(s.filename)

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		// Add record to CSV object
		s.addRow(Row{
			id:     record[0],
			source: record[1],
			target: record[2],
		})
	}
}

func (s Sheet) print() {
	for i, row := range *s.rows {
		fmt.Printf("Line # %d - %s, %s, %s\n", i, row.id, row.source, row.target)
	}
}
