package postgresql

import (
	"context"
)

// Delete deletes the existing Todo record by its ID, returns updated ID.
func (t *Todo) Delete(ctx context.Context, id int64) (int64, error) {
	deletedID, err := t.q.DeleteTodo(ctx, id)
	if err != nil {
		return deletedID, err
	}

	return deletedID, nil
}
