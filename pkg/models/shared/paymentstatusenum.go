package shared

type PaymentStatusEnum string

const (
	PaymentStatusEnumPending    PaymentStatusEnum = "PENDING"
	PaymentStatusEnumActive     PaymentStatusEnum = "ACTIVE"
	PaymentStatusEnumTerminated PaymentStatusEnum = "TERMINATED"
	PaymentStatusEnumFailed     PaymentStatusEnum = "FAILED"
)
