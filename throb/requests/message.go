package requests

type CreateMessageReq struct {
	Index       string `json:"index" valid:"Required"`
	Body        string `json:"body" valid:"Required"`
	Uid         string `json:"uid"`
}

type DeleteMessageReq struct {
	Index string `json:"index" valid:"Required"`
	Id    int64  `json:"id" valid:"Required"`
}

type DeleteIndexReq struct {
	Index string `json:"index" valid:"Required"`
}
