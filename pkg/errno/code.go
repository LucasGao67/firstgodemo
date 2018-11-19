package errno

// 用来统一存自定义的错误码，

var (
	// common errors
	OK                  = &Errno{0, "OK"}
	InternalServerError = &Errno{10001, "Internal server error."}
	ErrBind             = &Errno{10002, "Error occurred while binding the request body to the struct"}

	// user errors
	ErrUserNotFound = &Errno{20102, "The user was not found."}
)
