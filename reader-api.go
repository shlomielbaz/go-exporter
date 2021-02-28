package main

import "fmt"

// APIReader is ...
type APIReader struct {
	url string
}

func newAPIReader(url string) APIReader {
	ar := APIReader{
		url: url,
	}
	return ar
}

func (ar APIReader) Read() []byte {
	fmt.Println(ar.url)
	ba := []byte("READ FROM API READER")

	// fmt.Println(string(ba))

	return ba
}
