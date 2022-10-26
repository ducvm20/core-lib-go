package main

type LocalErr struct {
	Code   int         `json:"statusCode"`
	Msg    string      `json:"message"`
	Detail interface{} `json:"data,omitempty"`
}
type PubErr struct {
	Code   int         `json:"status"`
	Msg    string      `json:"msg"`
	Detail interface{} `json:"detail,omitempty"`
}

func (e LocalErr) Error() string {
	return e.Msg
}

const (
	DB_FAIL        = 300
	INVALID_PARAMS = 400
	LIMIT_ACCESS   = 400
	REQUIRED_TOKEN = 400
	EXPIRED_TOKEN  = 400
	INVALID_TOKEN  = 400
	BUSY           = 300
	UNKNOWN        = 301
)
const (
	MESS_DB_FAIL        = "Kết nối hệ thống lỗi, vui lòng thử lại sau ít phút"
	MESS_INVALID_PARAMS = "Sai thông tin đầu vào, vui lòng kiểm tra lại thông tin"
	MESS_LIMIT_ACCESS   = "Bạn truy cập quá nhanh."
	MESS_REQUIRED_TOKEN = "Token không tồn tại"
	MESS_EXPIRED_TOKEN  = "Token hết hạn"
	MESS_INVALID_TOKEN  = "Token không hợp lệ"
	MESS_BUSY           = "Hệ thống đang bận, vui lòng thử lại sau ít phút"
	MESS_UNKNOWN        = "Có lỗi trong quá trình xử lý, vui lòng thử lại sau ít phút"
)

var ErrDbFail *LocalErr = &LocalErr{Code: DB_FAIL, Msg: MESS_DB_FAIL}
var ErrParam *LocalErr = &LocalErr{Code: INVALID_PARAMS, Msg: MESS_INVALID_PARAMS}
var ErrRateLimit *LocalErr = &LocalErr{Code: LIMIT_ACCESS, Msg: MESS_LIMIT_ACCESS}
var ErrTokenRequired *LocalErr = &LocalErr{Code: REQUIRED_TOKEN, Msg: MESS_REQUIRED_TOKEN}
var ErrTokenExpire *LocalErr = &LocalErr{Code: EXPIRED_TOKEN, Msg: MESS_EXPIRED_TOKEN}
var ErrInvalidToken *LocalErr = &LocalErr{Code: INVALID_TOKEN, Msg: MESS_INVALID_TOKEN}
var ErrBusy *LocalErr = &LocalErr{Code: BUSY, Msg: MESS_BUSY}
var ErrorAccess *LocalErr = &LocalErr{Code: UNKNOWN, Msg: MESS_UNKNOWN}
