package responses

type Response struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	ErrMsg     []string    `json:"errmsg"`
	Data       interface{} `json:"data"`
}
