package passiontec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PayReport struct {
	CashMoney              int //现金
	CashProportion         string
	CashOpValue            int
	WechatMoney            int //微信
	WechatProportion       string
	WechatOpValue          int
	AlipayMoney            int //支付宝
	AlipayProportion       string
	AlipayOpValue          int
	VipMainMoney           int //会员消费
	VipMainProportion      string
	VipMainOpValue         int
	OncreditMoney          int //信用卡
	OncreditProportion     string
	OncreditOpValue        int
	BankCardMoney          int //银行卡
	BankCardMoneyDetail    string
	BankCardProportion     string
	BankCardOpValue        int
	ThirdpayMoney          int //第三方支付
	ThirdpayMoneyDetail    string
	ThirdpayProportion     string
	ThirdpayOpValue        int
	CouponpayMoney         int //优惠券
	CouponpayMoneyDetail   string
	CouponpayProportion    string
	CouponpayOpValue       int
	GroupbuypayMoney       int //团购
	GroupbuypayMoneyDetail string
	GroupbuypayProportion  string
	GroupbuypayOpValue     int
	MoneyTotal             int //总计
}

func (report PayReport) MoneyByPayName() map[string]float32 {
	return map[string]float32{
		"现金":    float32(report.CashMoney) / 100,
		"微信":    float32(report.WechatMoney) / 100,
		"支付宝":   float32(report.AlipayMoney) / 100,
		"会员消费":  float32(report.VipMainMoney) / 100,
		"信用卡":   float32(report.OncreditMoney) / 100,
		"银行卡":   float32(report.BankCardMoney) / 100,
		"第三方支付": float32(report.ThirdpayMoney) / 100,
		"优惠券":   float32(report.CouponpayMoney) / 100,
		"团购":    float32(report.GroupbuypayMoney) / 100,
	}
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
