package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfo struct {
	AgencyId         int
	AgencyName       string
	BelongId         int
	BelongName       string
	HotelId          int
	Ibid             int
	Mobile           string
	NativeForAndroid bool
	RealName         string
	SubName          string
	Type             int
	UserId           int
	Username         string
}

func (cli *Client) GetCurrentUserInfo() (*UserInfo, error) {
	if !cli.Logined {
		return nil, fmt.Errorf("未登录")
	}

	resp, err := cli.Get("https://e.passiontec.cn/hq-agency/user/current")
	if err != nil {
		return nil, fmt.Errorf("请求执行错误: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		cli.Logined = false
		return nil, fmt.Errorf("未登录")
	}

	var info struct {
		BaseResponse
		Data UserInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, fmt.Errorf("请求读取错误: %s", err)
	}

	return &info.Data, nil
}
