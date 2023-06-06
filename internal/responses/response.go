package responses

type Response struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	ErrMsg     []string    `json:"errmsg"`
	Data       interface{} `json:"data"`
}
