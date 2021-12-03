package rest

import (
	"net/http"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
	"github.com/gin-gonic/gin"
)

type ListTodosResponse struct {
	Todos []Todo `json:"todos"`
}

func (h *TodoHandler) list(c *gin.Context) {
	items, err := h.todoSvc.List(c)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	todos := make([]Todo, len(items))

	for i, item := range items {
		todos[i].ID = item.ID
		todos[i].Task = item.Task
		todos[i].IsDone = item.IsDone
	}

	resputil.SendJSON(c, http.StatusOK, "Successfully got todos list", &ListTodosResponse{
		Todos: todos,
	})
}