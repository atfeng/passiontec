package passiontec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PayReport struct {
	CashMoney              int
	CashProportion         string
	CashOpValue            int
	WechatMoney            int
	WechatProportion       string
	WechatOpValue          int
	AlipayMoney            int
	AlipayProportion       string
	AlipayOpValue          int
	VipMainMoney           int
	VipMainProportion      string
	VipMainOpValue         int
	OncreditMoney          int
	OncreditProportion     string
	OncreditOpValue        int
	BankCardMoney          int
	BankCardMoneyDetail    string
	BankCardProportion     string
	BankCardOpValue        int
	ThirdpayMoney          int
	ThirdpayMoneyDetail    string
	ThirdpayProportion     string
	ThirdpayOpValue        int
	CouponpayMoney         int
	CouponpayMoneyDetail   string
	CouponpayProportion    string
	CouponpayOpValue       int
	GroupbuypayMoney       int
	GroupbuypayMoneyDetail string
	GroupbuypayProportion  string
	GroupbuypayOpValue     int
	MoneyTotal             int
}

func (cli *Client) GetPayReportDetail(start, end time.Time) (*PayReport, error) {
	data := map[string]interface{}{
		"startTime":  start.Unix() * 1000,
		"endTime":    end.Unix() * 1000,
		"casherName": "",
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("参数错误: %s", err)
	}

	req, err := http.NewRequest("POST", "https://e.passiontec.cn/irs-api/zeus-web/pay-report/detail", bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("创建请求错误: %s", err)
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求执行错误: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求错误[%d]", resp.StatusCode)
	}

	var info struct {
		BaseResponse
		Data PayReport
	}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, fmt.Errorf("请求读取错误: %s", err)
	}

	return &info.Data, nil
}
