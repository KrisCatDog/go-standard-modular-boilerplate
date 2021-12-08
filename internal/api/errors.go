package api

// CommonErrorCode represents a type of internal common errors code.
type CommonErrorCode uint

// CodeSuccess means request executed successfully.
const CodeSuccess = 0

const (
	_ CommonErrorCode = iota // Skip the first value for success code

	// ErrCodeInternalServer means an error occured in internal codebase.
	ErrCodeInternalServer

	// ErrCodeInternalDatabase means an error occured in database.
	ErrCodeInternalDatabase

	// ErrCodeNotFound means an error occured because data was not found.
	ErrCodeNotFound

	// ErrCodeBadRequest means an error occured because the request does not match with expected format.
	ErrCodeBadRequest

	// ErrCodeFailedValidation means an error occured because the request does not meet the validation rules.
	ErrCodeFailedValidation
)
