package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid
	ErrInvalidToken     = 30001 // token is invalid
	ErrInvalidOTP       = 30002
	ErrSendEmailOTP     = 30003
	// Register Code
	ErrCodeUserHasExists = 5001 // user has already registered
)

// message

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP error",
	ErrSendEmailOTP:      "Failed to send email OTP",
	ErrCodeUserHasExists: "user has already registered",
}
