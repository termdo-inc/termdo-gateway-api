package schemas

type RefreshResponse struct {
	AccountID int    `json:"accountId"`
	Token     string `json:"token"`
}
