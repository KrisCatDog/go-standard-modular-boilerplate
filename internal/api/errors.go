package api

type GeneralErrorCode uint

const CodeSuccess = 0

const (
	ErrCodeInternalServer   GeneralErrorCode = iota + 1 // Represents errors that occur in the internal codebase
	ErrCodeInternalDatabase                             // Represents errors that occur in the database
	ErrCodeNotFound                                     // Represents errors that occur because data was not found
	ErrCodeBadRequest                                   // Represents errors that occur because the request does not match the expected format
	ErrCodeFailedValidation                             // Represents errors that occur because the request does not meet the validation requirements
)
