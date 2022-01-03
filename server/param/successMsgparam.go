package param

type SuccessMsgParam struct {
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Token  string      `json:"token"`
	RToken string      `json:"r_token"`
}
