package errcode

var (
	Success                   = NewError(0, "success")
	ServerErr                 = NewError(500101, "internal server error")
	InvalidParams             = NewError(400101, "invalid params")
	NotFound                  = NewError(404101, "not found")
	UnauthorizedAuthNotExist  = NewError(401101, "unauthorized auth not exist")
	UnauthorizedTokenErr      = NewError(401102, "token error")
	UnauthorizedTokenTimeout  = NewError(401103, "token timeout")
	UnauthorizedTokenGenerate = NewError(401104, "token generate error")
	ToomanyRequests           = NewError(429101, "too many requests")
)
