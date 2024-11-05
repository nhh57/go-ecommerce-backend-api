package response

const (
	ErrCodeSuccess      = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid

	ErrInvalidToken = 30001 // token is invalid
	ErrInvalidOTP   = 30002
	ErrSendEmailOTP = 30003
	// ErrCodeAuthFailed User Authentication
	ErrCodeAuthFailed = 40005
	// ErrCodeUserHasExists Register Code
	ErrCodeUserHasExists = 50001 // user has already registered

	// ErrCodeOTPNotExists Err Login
	ErrCodeOTPNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008
)

// message
var msg = map[int]string{
	ErrCodeSuccess:          "Success",
	ErrCodeParamInvalid:     "Email is invalid",
	ErrInvalidToken:         "Token is invalid",
	ErrInvalidOTP:           "OTP error",
	ErrSendEmailOTP:         "Failed to send email OTP",
	ErrCodeUserHasExists:    "User has already registered",
	ErrCodeOTPNotExists:     "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "Authentication failed",
}
