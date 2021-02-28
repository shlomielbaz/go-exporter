package main

import (
	"io/ioutil"
	"log"
)

// FileReader is ...
type FileReader struct {
	path string
}

func newFileReader(path string) FileReader {
	fr := FileReader{
		path: path,
	}
	return fr
}

func (fr FileReader) Read() []byte {
	// fmt.Println(fr.path)
	// ba := []byte("READ FROM FILE READER")

	// // fmt.Println(string(ba))

	// csvfile, err := os.Open(fr.path)

	// if err != nil {
	// 	log.Fatalln("Couldn't open the csv file", err)
	// }

	ba, err := ioutil.ReadFile(fr.path) // just pass the file name
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// if err != nil {
	//     fmt.Print(err)
	// }

	// fmt.Println(b) // print the content as 'bytes'

	// str := string(b) // convert content to a 'string'

	// fmt.Println(str) // print the content as a 'string'

	// // Parse the file
	// r := csv.NewReader(bufio.NewReader(csvfile))

	// // Iterate through the records
	// for {
	// 	// Read each record from csv
	// 	record, err := r.Read()

	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Add record to CSV object
	// 	s.addRow(Row{
	// 		id:     record[0],
	// 		source: record[1],
	// 		target: record[2],
	// 	})
	// }

	return ba
}
