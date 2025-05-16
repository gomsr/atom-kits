package couponx

// 优惠券类型； 0->全场赠券；1->会员赠券; 2->标签赠券; 3->购物赠券；4->注册赠券；5->生日赠券;6->网红码;
const (
	CouponTypeAll       = 0
	CouponTypeMember    = 1
	CouponTypeMemberTag = 2
	CouponTypeShopping  = 3
	CouponTypeRegister  = 4
	CouponTypeBirthday  = 5
	CouponTypeKolCode   = 6
)

// 发放人群规则: 0->全场赠券；1->会员赠券; 3->指定用户标签
const (
	CouponCrowdAll       = 0
	CouponCrowdMember    = 1
	CouponCrowdMemberTag = 2
)

// 优惠券分类: 0->Admin赠券; 1->购物赠券；2->注册赠券；3->生日赠券;4->网红码;

// 使用平台：0->全部；1->移动；2->PC；3->平板
const (
	PlatformAll    = 0
	PlatformMobile = 1
	PlatformPC     = 2
	PlatformTablet = 3
)

// 优惠券领取: 0->无需领取无需登录(逻辑券); 1->用户手动领取；2->发放到用户账户;
const (
	CouponReceiveTypeLogic  = 0 // 无需领取无需登录
	CouponReceiveTypeMember = 1 // 用户手动领取
	CouponReceiveTypeAdmin  = 2 // 发放到用户账户
)

// 优惠券通知: 0->不通知; 1->邮件通知; 2->短信通知
const (
	CouponNoticeTypeNo    = 0 // 不通知
	CouponNoticeTypeEmail = 1 // 邮件通知
	CouponNoticeTypeSms   = 2 // 短信通知
)

// 优惠券使用规则: 0->全场通用；1->指定分类；2->指定商品;
const (
	UseTypeAll             = 0
	UseTypeProductCategory = 1
	UseTypeProduct         = 2
)

// 优惠券类型: 0->满减；1->折扣;
const (
	CouponDisTypeFullReduce = 0
	CouponDisTypePercent    = 1
)

const (
	CouponUseStatusUnused  = 0
	CouponUseStatusUsed    = 1
	CouponUseStatusExpired = 2
)

func IsLogic(typo, receiveType int64) bool {
	return CouponTypeKolCode == typo ||
		(CouponReceiveTypeLogic == typo && receiveType == CouponReceiveTypeLogic)
}
