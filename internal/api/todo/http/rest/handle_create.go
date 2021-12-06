package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

type CreateTodoRequest struct {
	Task   string `json:"task" validate:"required"`
	IsDone bool   `json:"is_done" validate:"required"`
}

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) create(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Invalid type of request body", api.ErrCodeBadRequest))

		return
	}

	if err := h.validate.Struct(&req); err != nil {
		resputil.SendValidationFailed(c, err)

		return
	}

	newTodo, err := h.todoSvc.Create(c, todo.CreateParams{
		Task:   req.Task,
		IsDone: false,
	})
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendSuccess(c, http.StatusCreated, "Todo successfully created", &CreateTodoResponse{
		Todo: Todo{
			ID:     newTodo.ID,
			Task:   newTodo.Task,
			IsDone: newTodo.IsDone,
		},
	})
}
