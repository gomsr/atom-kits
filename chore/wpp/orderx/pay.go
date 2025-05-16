package orderx

import "strings"

// 支付方式：0->未支付；1->支付宝；2->微信; 3->paypal
const (
	PayTypeUnpaid = 0
	Unpaid        = "unpaid"
	PayTypeAli    = 1
	Alipay        = "alipay"
	PayTypeWx     = 2
	Wechatpay     = "wechatpay"
	PayTypePaypal = 3
	Paypal        = "paypal"
)

func PayType2Pb(payType string) int64 {
	if strings.EqualFold(payType, Unpaid) {
		return PayTypeUnpaid
	}
	if strings.EqualFold(payType, Alipay) {
		return PayTypeAli
	}
	if strings.EqualFold(payType, Wechatpay) {
		return PayTypeWx
	}
	if strings.EqualFold(payType, Paypal) {
		return PayTypePaypal
	}
	return -1
}

func PayType2Web(payType int64) string {
	if payType == PayTypeUnpaid {
		return Unpaid
	}
	if payType == PayTypeAli {
		return Alipay
	}
	if payType == PayTypeWx {
		return Wechatpay
	}

	if payType == PayTypePaypal {
		return Paypal
	}

	return ""
}

// 支付状态: 0->待付款；1->已付款；2->已关闭；3->退款异常；4->已退款；5->o同意支付
const (
	TradeNotPay    = 0
	TradeSuccess   = 1
	TradeClosed    = 2
	TradeRefundErr = 3
	TradeRefundSuc = 4
	TradeApprove   = 5

	TradeNotPayS    = "notpay"
	TradeSuccessS   = "success"
	TradeClosedS    = "closed"
	TradeRefundErrS = "refunderr"
	TradeRefundSucS = "refundsucc"
	TradeApproveS   = "approve"
)

func PayStatus2Pb(payStatus string) int64 {
	if strings.EqualFold(payStatus, TradeNotPayS) {
		return TradeNotPay
	}
	if strings.EqualFold(payStatus, TradeSuccessS) {
		return TradeSuccess
	}
	if strings.EqualFold(payStatus, TradeClosedS) {
		return TradeClosed
	}
	if strings.EqualFold(payStatus, TradeRefundErrS) {
		return TradeRefundErr
	}
	if strings.EqualFold(payStatus, TradeRefundSucS) {
		return TradeRefundSuc
	}
	if strings.EqualFold(payStatus, TradeApproveS) {
		return TradeApprove
	}

	return -1
}

func PayStatus2Web(payStatus int64) string {
	if payStatus == TradeNotPay {
		return TradeNotPayS
	}
	if payStatus == TradeSuccess {
		return TradeSuccessS
	}
	if payStatus == TradeClosed {
		return TradeClosedS
	}
	if payStatus == TradeRefundErr {
		return TradeRefundErrS
	}
	if payStatus == TradeRefundSuc {
		return TradeRefundSucS
	}
	if payStatus == TradeApprove {
		return TradeApproveS
	}
	return ""
}
