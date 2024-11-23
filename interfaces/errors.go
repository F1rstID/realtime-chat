package interfaces

// Error codes
const (
	// Success codes (2xxx)
	SuccessCode = 2000
	CreatedCode = 2001

	// Client error codes (4xxx)
	InvalidRequestError     = 4000
	UnauthorizedError       = 4001
	ForbiddenError          = 4002
	NotFoundError           = 4003
	EmailExistsError        = 4004
	NicknameExistsError     = 4005
	InvalidCredentialsError = 4006
	ValidationError         = 4007
	InvalidTokenError       = 4008

	// Server error codes (5xxx)
	InternalServerError = 5000
	DatabaseError       = 5001
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var errorMap = map[int]AppError{
	// Success
	SuccessCode: {
		Code:    SuccessCode,
		Message: "Success",
		Status:  200,
	},
	CreatedCode: {
		Code:    CreatedCode,
		Message: "Created successfully",
		Status:  201,
	},

	// Client errors
	InvalidRequestError: {
		Code:    InvalidRequestError,
		Message: "Invalid request",
		Status:  400,
	},
	UnauthorizedError: {
		Code:    UnauthorizedError,
		Message: "Unauthorized",
		Status:  401,
	},
	ForbiddenError: {
		Code:    ForbiddenError,
		Message: "Forbidden",
		Status:  403,
	},
	NotFoundError: {
		Code:    NotFoundError,
		Message: "Resource not found",
		Status:  404,
	},
	EmailExistsError: {
		Code:    EmailExistsError,
		Message: "Email already exists",
		Status:  400,
	},
	NicknameExistsError: {
		Code:    NicknameExistsError,
		Message: "Nickname already exists",
		Status:  400,
	},
	InvalidCredentialsError: {
		Code:    InvalidCredentialsError,
		Message: "Invalid email or password",
		Status:  401,
	},
	ValidationError: {
		Code:    ValidationError,
		Message: "Validation failed",
		Status:  400,
	},
	InvalidTokenError: {
		Code:    InvalidTokenError,
		Message: "Invalid or expired token",
		Status:  401,
	},

	// Server errors
	InternalServerError: {
		Code:    InternalServerError,
		Message: "Internal server error",
		Status:  500,
	},
	DatabaseError: {
		Code:    DatabaseError,
		Message: "Database error",
		Status:  500,
	},
}
