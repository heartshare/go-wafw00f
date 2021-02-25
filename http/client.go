package http

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

type Http interface {
	DoGet(url string, data map[string]string,
		headers map[string]string) (int, map[string]string, []byte)
	DoPost(url string, data map[string]string,
		headers map[string]string) (int, map[string]string, []byte)
	DoOther(url string, method string, data map[string]string,
		headers map[string]string) (int, map[string]string, []byte)
}

type Client struct {
}

func (h *Client) DoGet(url string, data map[string]string,
	headers map[string]string) (int, map[string]string, []byte) {
	return doRequest("GET", url, data, headers)
}

func (h *Client) DoPost(url string, data map[string]string,
	headers map[string]string) (int, map[string]string, []byte) {
	return doRequest("POST", url, data, headers)
}

func (h *Client) DoOther(url string, method string, data map[string]string,
	headers map[string]string) (int, map[string]string, []byte) {
	return doRequest(method, url, data, headers)
}

func doRequest(method string, url string, data map[string]string,
	headers map[string]string) (int, map[string]string, []byte) {
	client := http.Client{}
	var (
		req       *http.Request
		finalData []byte
	)
	if value, ok := headers["Content-Type"]; ok {
		if value == "application/json" {
			resolveJson(data)
		}
	} else {
		finalData = resolveForm(data)
	}
	if len(finalData) != 0 {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(finalData))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
	resp, err := client.Do(req)
	if err == nil {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		defer func() {
			if resp != nil {
				_ = resp.Body.Close()
			}
		}()
		respMap := make(map[string]string)
		for k, v := range resp.Header {
			respMap[k] = strings.Join(v, "")
		}
		return resp.StatusCode, respMap, body
	}
	return -1, nil, nil
}

func resolveForm(data map[string]string) []byte {
	var temp bytes.Buffer
	var finalData []byte
	for k, v := range data {
		temp.WriteString(k + "=" + v + "&")
	}
	finalData = []byte(strings.TrimRight(temp.String(), "&"))
	return finalData
}

func resolveJson(data map[string]string) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytesData, _ := json.Marshal(data)
	return bytesData
}
