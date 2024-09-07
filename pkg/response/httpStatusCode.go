package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid
	ErrInvalidToken     = 30001 // token is invalid

	// Register Code
	ErrCodeUserHasExists = 5001 // user has already registered
)

// message

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "user has already registered",
}
