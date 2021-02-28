package main

// CommonLoader is ...
type CommonLoader struct {
	buffer []byte
}

func (c CommonLoader) load(r IReader) []byte {

	c.buffer = r.Read()

	return c.buffer

}

func (c CommonLoader) read(r IReader) []byte {
	return r.Read()
}

func (c CommonLoader) write(w IWriter, buf []byte) {
	w.Write(buf)
}
