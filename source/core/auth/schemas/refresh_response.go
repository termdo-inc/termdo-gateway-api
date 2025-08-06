package schemas

type RefreshResponse = BaseResponse[struct {
	AccountID int `json:"accountId"`
}]
