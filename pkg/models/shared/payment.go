package shared

import (
	"time"
)

type PaymentSchemeEnum string

const (
	PaymentSchemeEnumVisa       PaymentSchemeEnum = "visa"
	PaymentSchemeEnumMastercard PaymentSchemeEnum = "mastercard"
	PaymentSchemeEnumAmex       PaymentSchemeEnum = "amex"
	PaymentSchemeEnumDiners     PaymentSchemeEnum = "diners"
	PaymentSchemeEnumDiscover   PaymentSchemeEnum = "discover"
	PaymentSchemeEnumJcb        PaymentSchemeEnum = "jcb"
	PaymentSchemeEnumUnionpay   PaymentSchemeEnum = "unionpay"
	PaymentSchemeEnumSepaDebit  PaymentSchemeEnum = "sepa debit"
	PaymentSchemeEnumSepaCredit PaymentSchemeEnum = "sepa credit"
	PaymentSchemeEnumSepa       PaymentSchemeEnum = "sepa"
	PaymentSchemeEnumApplePay   PaymentSchemeEnum = "apple pay"
	PaymentSchemeEnumGooglePay  PaymentSchemeEnum = "google pay"
	PaymentSchemeEnumA2a        PaymentSchemeEnum = "a2a"
	PaymentSchemeEnumAchDebit   PaymentSchemeEnum = "ach debit"
	PaymentSchemeEnumAch        PaymentSchemeEnum = "ach"
	PaymentSchemeEnumRtp        PaymentSchemeEnum = "rtp"
	PaymentSchemeEnumUnknown    PaymentSchemeEnum = "unknown"
	PaymentSchemeEnumOther      PaymentSchemeEnum = "other"
)

type PaymentTypeEnum string

const (
	PaymentTypeEnumPayIn    PaymentTypeEnum = "PAY-IN"
	PaymentTypeEnumPayout   PaymentTypeEnum = "PAYOUT"
	PaymentTypeEnumTransfer PaymentTypeEnum = "TRANSFER"
	PaymentTypeEnumOther    PaymentTypeEnum = "OTHER"
)

type Payment struct {
	AccountID     string                 `json:"accountID"`
	Adjustments   []PaymentAdjustment    `json:"adjustments"`
	Asset         string                 `json:"asset"`
	CreatedAt     time.Time              `json:"createdAt"`
	ID            string                 `json:"id"`
	InitialAmount int64                  `json:"initialAmount"`
	Metadata      []PaymentMetadata      `json:"metadata"`
	Provider      ConnectorEnum          `json:"provider"`
	Raw           map[string]interface{} `json:"raw"`
	Reference     string                 `json:"reference"`
	Scheme        PaymentSchemeEnum      `json:"scheme"`
	Status        PaymentStatusEnum      `json:"status"`
	Type          PaymentTypeEnum        `json:"type"`
}
