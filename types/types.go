package types

type ContextKey int

const (
	TodoIDKey ContextKey = iota
	UserKey   ContextKey = iota
)
