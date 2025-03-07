package tunnel

type Response struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	UserID      string `json:"userId"`
	AccessToken string `json:"accessToken"`
	Nickname    string `json:"nickname"`
	RealName    string `json:"realname"`
	OuterAudit  int    `json:"outerAudit"`
}
