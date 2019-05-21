package AliYunTextScan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) AliTextScan() (*AliResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, host+c.Path+"?clientInfo="+url.QueryEscape(c.ClientParams), strings.NewReader(c.Body))
	if err != nil {
		return nil, err
	}

	addRequestHeader(c.Body, req, c.ClientParams, c.Path, c.AccessKeyId, c.AccessKeySecret)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp AliResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
