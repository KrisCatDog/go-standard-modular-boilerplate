package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

// updateTodoRequestURI defines the payload of request URI to update a Todo by ID.
type updateTodoRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

// createTodoRequest defines the payload of request body to update a Todo.
type updateTodoRequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

// updateTodoResponse defines the payload of response body after updating Todo.
type updateTodoResponse struct {
	Todo Todo `json:"todo"`
}

// update handle incoming PUT request to update an existing Todo record.
func (h *TodoHandler) update(c *gin.Context) {
	var reqURI updateTodoRequestURI
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	var reqBody updateTodoRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Invalid type of request body", api.ErrCodeBadRequest))

		return
	}

	updatedID, err := h.todoSvc.Update(c, reqURI.ID, todo.UpdateParams{
		Task:   reqBody.Task,
		IsDone: reqBody.IsDone,
	})
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendSuccess(c, http.StatusOK, "Todo successfully updated", &updateTodoResponse{
		Todo: Todo{
			ID: updatedID,
		},
	})
}
