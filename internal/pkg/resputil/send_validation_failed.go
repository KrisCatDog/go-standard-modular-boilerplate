package resputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
)

// validationErrorCode defines a custom type for internal error codes.
type validationErrorCode uint

// fieldError represents base JSON structure for each validated fields.
type fieldError struct {
	Code    validationErrorCode `json:"code"`
	Message string              `json:"message"`
}

// Validator internal errors list.
const (
	ErrRequired validationErrorCode = iota + 100 // Represents that the field is not filled
)

// String returns an actual text for the error code.
func (c validationErrorCode) String() string {
	return [...]string{
		"The %v field is required", // ErrRequired - required
	}[c-100]
}

// errorsWrapper maps validation errors with the validation rules as a key.
var errorsWrapper = map[string]validationErrorCode{
	"required": ErrRequired,
}

// SendValidationFailed returns formatted JSON response messages related to validation errors.
func SendValidationFailed(c *gin.Context, err error) {
	errs := make(map[string]interface{})

	for _, err := range err.(validator.ValidationErrors) {
		errw := errorsWrapper[err.Tag()]

		errs[err.Field()] = fieldError{
			Code:    errw,
			Message: fmt.Sprintf(errw.String(), err.Field()),
		}
	}

	c.JSON(http.StatusUnprocessableEntity, &baseErrorResponse{
		Code:    api.ErrCodeFailedValidation,
		Message: "The given data was invalid",
		Errors:  errs,
	})
}
