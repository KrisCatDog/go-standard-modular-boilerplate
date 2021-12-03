package rest

import (
	"net/http"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
	"github.com/gin-gonic/gin"
)

type FindTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type FindTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) find(c *gin.Context) {
	var req FindTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	singleTodo, err := h.todoSvc.Find(c, req.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendJSON(c, http.StatusOK, "Successfully got todo details", &FindTodoResponse{
		Todo: Todo{
			ID:     singleTodo.ID,
			Task:   singleTodo.Task,
			IsDone: singleTodo.IsDone,
		},
	})
}
