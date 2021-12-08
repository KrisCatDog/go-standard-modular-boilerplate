package todo

// Todo defines an entity for the service layer.
type Todo struct {
	ID     int64
	Task   string
	IsDone bool
}
