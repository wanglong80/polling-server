package requests

type CreateStorageReq struct {
	Key  string                 `json:"key" valid:"Required"`
	Data map[string]interface{} `json:"data" valid:"Required"`
}
