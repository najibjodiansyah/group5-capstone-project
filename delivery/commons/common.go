package commons

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponseLogin struct {
	Code            int         `json:"code"`
	Status          string      `json:"status"`
	Message         string      `json:"message"`
	Data            interface{} `json:"data"`
	CurrentUserName interface{} `json:"current_user_name"`
}

//NewInternalServerErrorResponse default internal server error response
func SuccessOperationDefault(status, message string) DefaultResponse {
	return DefaultResponse{
		200,
		status,
		message,
	}
}

func SuccessOperation(status, message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		200,
		status,
		message,
		data,
	}
}

func SuccessOperationLogin(status, message string, user, currentUserName interface{}) SuccessResponseLogin {
	return SuccessResponseLogin{
		200,
		status,
		message,
		user,
		currentUserName,
	}
}

//NewInternalServerErrorResponse default internal server error response
func InternalServerError(status, message string) DefaultResponse {
	return DefaultResponse{
		500,
		status,
		message,
	}
}

//NewNotFoundResponse default not found error response
func NotFound(status, message string) DefaultResponse {
	return DefaultResponse{
		404,
		status,
		message,
	}
}

//NewBadRequestResponse default not found error response
func BadRequest(status, message string) DefaultResponse {
	return DefaultResponse{
		400,
		status,
		message,
	}
}

//ForbiddedRequest default not found error response
func ForbiddedRequest(status, message string) DefaultResponse {
	return DefaultResponse{
		403,
		status,
		message,
	}
}

//ForbiddedRequest default not found error response
func UnauthorizedRequest(status, message string) DefaultResponse {
	return DefaultResponse{
		401,
		status,
		message,
	}
}

//NewConflictResponse default not found error response
func Conflict(status, message string) DefaultResponse {
	return DefaultResponse{
		409,
		status,
		message,
	}
}
