package main

// comment
type IWriter interface {
	Write(p []byte) (n int, err error)
}
