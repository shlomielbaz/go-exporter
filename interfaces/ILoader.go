package main

// comment on exported type ILoader should be of the form "ILoader ..." (with optional leading article)
type ILoader interface {
	Load(r IReader)
}
