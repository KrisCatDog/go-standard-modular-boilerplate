package todo

// CreateParams defines the fields used for creating new Todo records.
type CreateParams struct {
	Task   string
	IsDone bool
}

// CreateParams defines the fields used for updating Todo records.
type UpdateParams struct {
	Task   string
	IsDone bool
}
