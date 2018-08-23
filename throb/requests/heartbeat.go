package requests

type HeartbeatReq struct {
	Ms []MessageReq `json:"ms"`
	Ss []string     `json:"ss"`
}

type MessageReq struct {
	Name string `json:"n" valid:"Required"`
	GtId int64  `json:"gt" valid:"Numeric"`
	LtId int64  `json:"lt" valid:"Numeric"`
	Last bool   `json:"la"`
}
