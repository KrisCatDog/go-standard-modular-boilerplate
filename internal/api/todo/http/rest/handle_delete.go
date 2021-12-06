package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/resputil"
)

type DeleteTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type DeleteTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) delete(c *gin.Context) {
	var reqURI DeleteTodoRequest
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	deletedID, err := h.todoSvc.Delete(c, reqURI.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendSuccess(c, http.StatusOK, "Todo successfully deleted", &DeleteTodoResponse{
		Todo: Todo{
			ID: deletedID,
		},
	})
}
