package schemas

type HttpStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ServerError struct {
	Name       string  `json:"name"`
	Message    string  `json:"message"`
	StackTrace *string `json:"stackTrace"`
}

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponse[T any] struct {
	HttpStatus   HttpStatus    `json:"httpStatus"`
	ServerError  *ServerError  `json:"serverError"`
	ClientErrors []ClientError `json:"clientErrors"`
	Data         T             `json:"data"`
	Token        *string       `json:"token"`
}
