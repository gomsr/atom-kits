package orderx

import "strings"

const (
	RefundProcessing = 0 // PROCESSING
	RefundSuccess    = 1 //"SUCCESS"
	RefundClosed     = 2 //"CLOSED"
	RefundAbnormal   = 3 //"ABNORMAL"
	RefundInit       = 4 // INIT

	RefundProcessingS = "processing" // PROCESSING
	RefundSuccessS    = "success"    //"SUCCESS"
	RefundClosedS     = "closed"     //"CLOSED"
	RefundAbnormalS   = "abnormal"   //"ABNORMAL"
	RefundInitS       = "init"
)

func RefundStatus2Pb(payStatus string) int64 {
	if strings.EqualFold(payStatus, RefundProcessingS) {
		return RefundProcessing
	}
	if strings.EqualFold(payStatus, RefundSuccessS) {
		return RefundSuccess
	}
	if strings.EqualFold(payStatus, RefundClosedS) {
		return RefundClosed
	}
	if strings.EqualFold(payStatus, RefundAbnormalS) {
		return RefundAbnormal
	}
	if strings.EqualFold(payStatus, RefundInitS) {
		return RefundInit
	}

	return -1
}

func RefundStatus2Web(payStatus int64) string {
	if payStatus == RefundProcessing {
		return RefundProcessingS
	}
	if payStatus == RefundSuccess {
		return RefundSuccessS
	}
	if payStatus == RefundClosed {
		return RefundClosedS
	}
	if payStatus == RefundAbnormal {
		return RefundAbnormalS
	}
	if payStatus == RefundInit {
		return RefundInitS
	}

	return ""
}
