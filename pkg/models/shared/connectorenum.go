package shared

type ConnectorEnum string

const (
	ConnectorEnumStripe        ConnectorEnum = "STRIPE"
	ConnectorEnumDummyPay      ConnectorEnum = "DUMMY-PAY"
	ConnectorEnumWise          ConnectorEnum = "WISE"
	ConnectorEnumModulr        ConnectorEnum = "MODULR"
	ConnectorEnumCurrencyCloud ConnectorEnum = "CURRENCY-CLOUD"
	ConnectorEnumBankingCircle ConnectorEnum = "BANKING-CIRCLE"
)
