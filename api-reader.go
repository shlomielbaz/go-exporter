package main

type apiReader struct{}

func (apiReader) Read() (p []byte) {

	ba := []byte("READ FROM API READER")

	// fmt.Println(string(ba))

	return ba
}
