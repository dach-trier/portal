package repo

type Cursor[T any] struct {
	After T
	Limit int
}
