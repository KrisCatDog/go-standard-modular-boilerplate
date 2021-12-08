package resputil

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
)

// baseSuccessResponse represents base JSON response structure for failed request.
type baseErrorResponse struct {
	Code    api.CommonErrorCode `json:"code"`
	Message string              `json:"message"`
	Errors  interface{}         `json:"errors"`
}

// SendError returns formatted JSON response for failed request.
func SendError(c *gin.Context, err error) {
	var httpStatus int
	var resp baseErrorResponse
	var internalErr *errorsutil.InternalError

	if errors.As(err, &internalErr) {
		switch internalErr.Code() {
		case api.ErrCodeNotFound:
			httpStatus = http.StatusNotFound
		case api.ErrCodeBadRequest:
			httpStatus = http.StatusBadRequest
		case api.ErrCodeInternalServer:
			fallthrough
		default:
			httpStatus = http.StatusInternalServerError
		}

		resp.Code = internalErr.Code()
		resp.Message = internalErr.Error()

	} else {
		httpStatus = http.StatusInternalServerError

		resp.Code = api.ErrCodeInternalServer
		resp.Message = "Internal server error"
	}

	c.JSON(httpStatus, &resp)
}
