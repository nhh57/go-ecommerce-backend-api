package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid
	ErrInvalidToken     = 30001 // token is invalid
	ErrInvalidOTP       = 30002
	ErrSendEmailOTP     = 30003
	// Register Code
	ErrCodeUserHasExists = 5001 // user has already registered

	// Err Login
	ErrCodeOTPNotExists = 60009
)

// message

var msg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP error",
	ErrSendEmailOTP:      "Failed to send email OTP",
	ErrCodeUserHasExists: "User has already registered",
	ErrCodeOTPNotExists:  "OTP exists but not registered",
}
