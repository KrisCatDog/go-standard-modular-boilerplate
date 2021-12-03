package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

type UpdateTodoRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateTodoRequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type UpdateTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) update(c *gin.Context) {
	var reqURI UpdateTodoRequestURI
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	var reqBody UpdateTodoRequestBody
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

	resputil.SendJSON(c, http.StatusOK, "Todo successfully updated", &UpdateTodoResponse{
		Todo: Todo{
			ID: updatedID,
		},
	})
}
