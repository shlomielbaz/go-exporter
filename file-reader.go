package main

type fileReader struct{}

func (fileReader) Read() []byte {
	ba := []byte("READ FROM FILE READER")

	// fmt.Println(string(ba))

	return ba
}
