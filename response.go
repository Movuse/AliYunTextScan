package AliYunTextScan

type AliResponse struct {
	Code      int        `json:"code"`
	Msg       string     `json:"msg"`
	RequestID string     `json:"requestId"`
	Data      []RespData `json:"data"`
}

type RespData struct {
	Code    int          `json:"code"`
	Content string       `json:"content"`
	DataID  string       `json:"dataId"`
	Msg     string       `json:"msg"`
	Results []RespDetail `json:"results"`
}

type RespDetail struct {
	Label      string  `json:"label"`
	Rate       float64 `json:"rate"`
	Scene      string  `json:"scene"`
	Suggestion string  `json:"suggestion"`
	Details    []struct {
		Contexts []struct {
			Context string `json:"context"`
		} `json:"contexts"`
		Label string `json:"label"`
	} `json:"details"`
}
