package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

// listTodosResponse defines the payload of response body for the todos list.
type listTodosResponse struct {
	Todos []Todo `json:"todos"`
}

// list handle incoming GET request to send the list of Todos from the datastore.
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

	resputil.SendSuccess(c, http.StatusOK, "Successfully got todos list", &listTodosResponse{
		Todos: todos,
	})
}
