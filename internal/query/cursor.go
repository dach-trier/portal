package query

type Cursor[T any] struct {
	Limit int

	// optional
	After *T
}
