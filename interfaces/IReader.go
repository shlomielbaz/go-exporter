package interfaces

// comment on exported type ILoader should be of the form "ILoader ..." (with optional leading article)
type IReader interface {
	Read(p []byte) (n int, err error)
}
