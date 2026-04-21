package query

type Cursor[T any] struct {
	After T
	Limit int
}
