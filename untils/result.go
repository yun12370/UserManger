package untils

type Result[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Success[T any](message string, data T) Result[T] {
	return Result[T]{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func Fail[T any](code int, message string) Result[T] {
	return Result[T]{
		Code:    code,
		Message: message,
	}
}
