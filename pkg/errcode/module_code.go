package errcode

var (
	ErrorGetTagFail     = NewError(10010001, "get tag fail")
	ErrorGetTagListFail = NewError(10010002, "get tag list fail")
	ErrorCreateTagFail  = NewError(10010003, "create tag fail")
	ErrorUpdateTagFail  = NewError(10010004, "update tag fail")
	ErrorDeleteTagFail  = NewError(10010005, "delete tag fail")
	ErrorCountTagFail   = NewError(10010006, "count tag fail")
)
