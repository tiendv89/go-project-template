package utils

type IError interface {
	RespCode() uint

	RespMsg() string
}

type Error struct {
	Code    uint
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) RespCode() uint {
	return e.Code
}

func (e Error) RespMsg() string {
	return e.Message
}
