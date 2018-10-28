package passiontec

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (cli *Client) CaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//初始化会话
	resp, err := cli.Get("http://irs.passiontec.cn")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	resp2, err := cli.Get("http://irs.passiontec.cn/irs-web/captcha")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp2.Body.Close()

	w.Header().Set("Content-Type", resp2.Header.Get("Content-Type"))

	_, err = io.Copy(w, resp2.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type LoginResponse struct {
	Code       string
	Success    bool
	ErrorCode  int
	ErrorMsg   string
	ErrorInfo  string
	TotalCount int
}

func (cli *Client) DoLogin(userName, employeeName, password, validateCode string) error {
	data := url.Values{
		"userName":     {userName},
		"employeeName": {employeeName},
		"password":     {password},
		"validateCode": {validateCode},
	}
	resp, err := cli.PostForm("http://irs.passiontec.cn/irs-web/employeeLogin", data)
	if err != nil {
		return fmt.Errorf("登陆请求错误: %s", err)
	}
	defer resp.Body.Close()

	var info LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return fmt.Errorf("登陆解析错误: %s", err)
	}

	if info.Code != "200" || info.ErrorInfo == "" {
		return fmt.Errorf("非法的登陆响应: %+s", info)
	}

	resp2, err := cli.Get(info.ErrorInfo)
	if err != nil {
		return fmt.Errorf("登陆跳转错误: %s", err)
	}
	defer resp2.Body.Close()

	cli.Logined = true

	return nil
}
