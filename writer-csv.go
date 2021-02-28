package main

import "fmt"

// CSVWriter is ...
type CSVWriter struct {
}

func (ar CSVWriter) Write(buff []byte) {

	fmt.Println(string(buff))
}
