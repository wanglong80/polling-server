package requests

type GetMessageListReq struct {
	Index  string `json:"index" valid:"Required"`
	Uid    int64  `json:"uid"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

type CreateMessageReq struct {
	Index       string `json:"index" valid:"Required"`
	Type        int64  `json:"type" valid:"Required"`
	Uid         int64  `json:"uid" valid:"Required"`
	Body        string `json:"body" valid:"Required"`
	Persistence bool   `json:"persistence" valid:"Required"`
}

type DeleteMessageReq struct {
	Index string `json:"index" valid:"Required"`
	Id    int64  `json:"id" valid:"Required"`
}

type DeleteIndexReq struct {
	Index string `json:"index" valid:"Required"`
}
