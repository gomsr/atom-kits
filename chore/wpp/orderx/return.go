package orderx

import "strings"

const (
	ReturnUnhandled  = 0           // 未处理
	ReturnUnhandledS = "unhandled" // 未处理
	ReturnReturning  = 1           // 退货中
	ReturnReturningS = "returning" // 退货中
	ReturnReturned   = 2           // 已收货完成
	ReturnReturnedS  = "returned"  // 已收货完成
	ReturnRejected   = 3           // 已拒绝
	ReturnRejectedS  = "rejected"  // 已拒绝

	ReturnRefunding  = 4           // 退款ing
	ReturnRefundingS = "refunding" // 退款ing
	ReturnRefunded   = 5           // 已退款
	ReturnRefundedS  = "refunded"  // 已退款
	ReturnRefundErr  = 6           // 退款异常
	ReturnRefundErrS = "refunderr" // 退款异常
	ReturnClosed     = 7           // 已取消
	ReturnClosedS    = "closed"    // 已取消
)

func ReturnStatus2Pb(returnStatus string) int64 {
	if strings.EqualFold(returnStatus, ReturnUnhandledS) {
		return ReturnUnhandled
	}
	if strings.EqualFold(returnStatus, ReturnReturningS) {
		return ReturnReturning
	}
	if strings.EqualFold(returnStatus, ReturnReturnedS) {
		return ReturnReturned
	}
	if strings.EqualFold(returnStatus, ReturnRejectedS) {
		return ReturnRejected
	}
	if strings.EqualFold(returnStatus, ReturnRefundingS) {
		return ReturnRefunding
	}
	if strings.EqualFold(returnStatus, ReturnRefundedS) {
		return ReturnRefunded
	}
	if strings.EqualFold(returnStatus, ReturnRefundErrS) {
		return ReturnRefundErr
	}
	if strings.EqualFold(returnStatus, ReturnClosedS) {
		return ReturnClosed
	}
	return -1
}

func ReturnStatus2Web(returnStatus int64) string {
	if returnStatus == ReturnUnhandled {
		return ReturnUnhandledS
	}
	if returnStatus == ReturnReturning {
		return ReturnReturningS
	}
	if returnStatus == ReturnReturned {
		return ReturnReturnedS
	}
	if returnStatus == RefundAbnormal {
		return RefundAbnormalS
	}
	if returnStatus == ReturnRejected {
		return ReturnRejectedS
	}
	if returnStatus == ReturnRefunding {
		return ReturnRefundingS
	}
	if returnStatus == ReturnRefunded {
		return ReturnRefundedS
	}
	if returnStatus == ReturnRefundErr {
		return ReturnRefundErrS
	}
	if returnStatus == ReturnClosed {
		return ReturnClosedS
	}

	return ""
}
