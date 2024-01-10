// author: jiazujiang
// date: 2023-06-30
// desc: 通用http请求
package cl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var (
	cl *http.Client
	tr *http.Transport
)

func init() {
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}
	cl = &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second}
}

// 通用返回结构体
type (
	ClRespModel struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

func Get(url string, resp interface{}) error {
	clReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	clResp, err := cl.Do(clReq)
	if err != nil {
		return err
	}
	defer clResp.Body.Close()
	clBody, err := io.ReadAll(clResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(clBody, &resp)
	if err != nil {
		return err
	}
	return nil
}

func PostJson(url string, body []byte, resp interface{}) (err error) {
	clReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	clReq.Header.Set("Content-Type", "application/json")
	clResp, err := cl.Do(clReq)
	if err != nil {
		return err
	}
	defer clResp.Body.Close()
	clBody, err := io.ReadAll(clResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(clBody, &resp)
	if err != nil {
		return err
	}
	return nil
}

func PostJsonStruct(url string, body interface{}, resp interface{}) (err error) {
	bodyByte, err := json.Marshal(&body)
	if err != nil {
		return err
	}
	clReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return err
	}
	clReq.Header.Set("Content-Type", "application/json")
	clResp, err := cl.Do(clReq)
	if err != nil {
		return err
	}
	defer clResp.Body.Close()
	clBody, err := io.ReadAll(clResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(clBody, &resp)
	if err != nil {
		return err
	}
	return nil
}

func PostJsonStructWithHeader(url string, body interface{}, header map[string]string, resp interface{}) (err error) {
	bodyByte, err := json.Marshal(&body)
	if err != nil {
		return err
	}
	clReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return err
	}
	clReq.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		clReq.Header.Set(k, v)
	}
	clResp, err := cl.Do(clReq)
	if err != nil {
		return err
	}
	defer clResp.Body.Close()
	clBody, err := io.ReadAll(clResp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(clBody, &resp)
	if err != nil {
		return err
	}
	return nil
}
