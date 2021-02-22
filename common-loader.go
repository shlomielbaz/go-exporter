package main

import "fmt"

type commonLoader struct{}

func (commonLoader) load(r IReader) {

	buf := r.Read()

	fmt.Println(string(buf))

}
