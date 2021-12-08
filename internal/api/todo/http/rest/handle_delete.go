package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

// deleteTodoRequest defines the payload of request URI to delete a todo.
type deleteTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// deleteTodoResponse defines the payload of response body after deleting Todo.
type deleteTodoResponse struct {
	Todo Todo `json:"todo"`
}

// delete handle incoming DELETE request to delete a Todo from the datastore.
func (h *TodoHandler) delete(c *gin.Context) {
	var reqURI deleteTodoRequest
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	deletedID, err := h.todoSvc.Delete(c, reqURI.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendSuccess(c, http.StatusOK, "Todo successfully deleted", &deleteTodoResponse{
		Todo: Todo{
			ID: deletedID,
		},
	})
}
