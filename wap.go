package alipay

import (
	"net/http"
	"net/url"
	"strings"
)

// https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (this *AliPay) TradeWapPay(param AliPayTradeWapPay) (url *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}
	var buf = strings.NewReader(p.Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	rep, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	if err != nil {
		return nil, err
	}
	url = rep.Request.URL
	return url, err
}

func (this *AliPay) TradeWapPayv2(param AliPayTradeWapPay) (results string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}
