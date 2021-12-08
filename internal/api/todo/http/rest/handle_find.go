package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

// findTodoRequest defines the payload of request URI to find a todo by ID.
type findTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// findTodoResponse defines the payload of response body for the requested Todo details.
type findTodoResponse struct {
	Todo Todo `json:"todo"`
}

// find handle incoming GET request to find a Todo by ID from the datastore.
func (h *TodoHandler) find(c *gin.Context) {
	var req findTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	singleTodo, err := h.todoSvc.Find(c, req.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendSuccess(c, http.StatusOK, "Successfully got todo details", &findTodoResponse{
		Todo: Todo{
			ID:     singleTodo.ID,
			Task:   singleTodo.Task,
			IsDone: singleTodo.IsDone,
		},
	})
}
