package AliYunTextScan

import (
	"encoding/json"
	"mashuai/aliyundun/uuid"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Path            string //request path
	ClientParams    string
	Body            string //request body json
}

type ClientInfo struct {
	SdkVersion  string `json:"sdkVersion"`
	CfgVersion  string `json:"cfgVersion"`
	UserType    string `json:"userType"`
	UserID      string `json:"userId"`
	UserNick    string `json:"userNick"`
	Avatar      string `json:"avatar"`
	Imei        string `json:"imei"`
	Imsi        string `json:"imsi"`
	Umid        string `json:"umid"`
	IP          string `json:"ip"`
	Os          string `json:"os"`
	Channel     string `json:"channel"`
	HostAppName string `json:"hostAppName"`
	HostPackage string `json:"hostPackage"`
	HostVersion string `json:"hostVersion"`
}

type TextScanBody struct {
	Scenes []string       `json:"scenes"`
	Tasks  []TextScanTask `json:"tasks"`
}

type TextScanTask struct {
	DataID  string `json:"dataId"`
	Content string `json:"content"`
}

func NewTextScan(accessKey, secret, path string, clientInfo *ClientInfo, contents []string) (*Client, error) {
	client := &Client{}
	client.AccessKeyId = accessKey
	client.AccessKeySecret = secret
	client.Path = path

	var tasks []TextScanTask
	for _, content := range contents {
		var task TextScanTask
		task.DataID = uuid.Rand().Hex()
		task.Content = content
		tasks = append(tasks, task)
	}

	body := TextScanBody{
		Scenes: []string{"antispam"},
		Tasks:  tasks,
	}

	bodyByte, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	infoByte, err := json.Marshal(clientInfo)
	if err != nil {
		return nil, err
	}

	client.ClientParams = string(infoByte)
	client.Body = string(bodyByte)
	return client, nil
}
