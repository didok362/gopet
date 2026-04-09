package domain

type Nulladble[T any] struct {
	Value *T
	Set   bool
}
