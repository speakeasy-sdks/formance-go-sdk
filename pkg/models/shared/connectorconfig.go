// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/utils"
)

type ConnectorConfigType string

const (
	ConnectorConfigTypeStripeConfig        ConnectorConfigType = "StripeConfig"
	ConnectorConfigTypeDummyPayConfig      ConnectorConfigType = "DummyPayConfig"
	ConnectorConfigTypeWiseConfig          ConnectorConfigType = "WiseConfig"
	ConnectorConfigTypeModulrConfig        ConnectorConfigType = "ModulrConfig"
	ConnectorConfigTypeCurrencyCloudConfig ConnectorConfigType = "CurrencyCloudConfig"
	ConnectorConfigTypeBankingCircleConfig ConnectorConfigType = "BankingCircleConfig"
)

type ConnectorConfig struct {
	StripeConfig        *StripeConfig
	DummyPayConfig      *DummyPayConfig
	WiseConfig          *WiseConfig
	ModulrConfig        *ModulrConfig
	CurrencyCloudConfig *CurrencyCloudConfig
	BankingCircleConfig *BankingCircleConfig

	Type ConnectorConfigType
}

func CreateConnectorConfigStripeConfig(stripeConfig StripeConfig) ConnectorConfig {
	typ := ConnectorConfigTypeStripeConfig

	return ConnectorConfig{
		StripeConfig: &stripeConfig,
		Type:         typ,
	}
}

func CreateConnectorConfigDummyPayConfig(dummyPayConfig DummyPayConfig) ConnectorConfig {
	typ := ConnectorConfigTypeDummyPayConfig

	return ConnectorConfig{
		DummyPayConfig: &dummyPayConfig,
		Type:           typ,
	}
}

func CreateConnectorConfigWiseConfig(wiseConfig WiseConfig) ConnectorConfig {
	typ := ConnectorConfigTypeWiseConfig

	return ConnectorConfig{
		WiseConfig: &wiseConfig,
		Type:       typ,
	}
}

func CreateConnectorConfigModulrConfig(modulrConfig ModulrConfig) ConnectorConfig {
	typ := ConnectorConfigTypeModulrConfig

	return ConnectorConfig{
		ModulrConfig: &modulrConfig,
		Type:         typ,
	}
}

func CreateConnectorConfigCurrencyCloudConfig(currencyCloudConfig CurrencyCloudConfig) ConnectorConfig {
	typ := ConnectorConfigTypeCurrencyCloudConfig

	return ConnectorConfig{
		CurrencyCloudConfig: &currencyCloudConfig,
		Type:                typ,
	}
}

func CreateConnectorConfigBankingCircleConfig(bankingCircleConfig BankingCircleConfig) ConnectorConfig {
	typ := ConnectorConfigTypeBankingCircleConfig

	return ConnectorConfig{
		BankingCircleConfig: &bankingCircleConfig,
		Type:                typ,
	}
}

func (u *ConnectorConfig) UnmarshalJSON(data []byte) error {

	wiseConfig := new(WiseConfig)
	if err := utils.UnmarshalJSON(data, &wiseConfig, "", true, true); err == nil {
		u.WiseConfig = wiseConfig
		u.Type = ConnectorConfigTypeWiseConfig
		return nil
	}

	stripeConfig := new(StripeConfig)
	if err := utils.UnmarshalJSON(data, &stripeConfig, "", true, true); err == nil {
		u.StripeConfig = stripeConfig
		u.Type = ConnectorConfigTypeStripeConfig
		return nil
	}

	dummyPayConfig := new(DummyPayConfig)
	if err := utils.UnmarshalJSON(data, &dummyPayConfig, "", true, true); err == nil {
		u.DummyPayConfig = dummyPayConfig
		u.Type = ConnectorConfigTypeDummyPayConfig
		return nil
	}

	modulrConfig := new(ModulrConfig)
	if err := utils.UnmarshalJSON(data, &modulrConfig, "", true, true); err == nil {
		u.ModulrConfig = modulrConfig
		u.Type = ConnectorConfigTypeModulrConfig
		return nil
	}

	currencyCloudConfig := new(CurrencyCloudConfig)
	if err := utils.UnmarshalJSON(data, &currencyCloudConfig, "", true, true); err == nil {
		u.CurrencyCloudConfig = currencyCloudConfig
		u.Type = ConnectorConfigTypeCurrencyCloudConfig
		return nil
	}

	bankingCircleConfig := new(BankingCircleConfig)
	if err := utils.UnmarshalJSON(data, &bankingCircleConfig, "", true, true); err == nil {
		u.BankingCircleConfig = bankingCircleConfig
		u.Type = ConnectorConfigTypeBankingCircleConfig
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ConnectorConfig) MarshalJSON() ([]byte, error) {
	if u.StripeConfig != nil {
		return utils.MarshalJSON(u.StripeConfig, "", true)
	}

	if u.DummyPayConfig != nil {
		return utils.MarshalJSON(u.DummyPayConfig, "", true)
	}

	if u.WiseConfig != nil {
		return utils.MarshalJSON(u.WiseConfig, "", true)
	}

	if u.ModulrConfig != nil {
		return utils.MarshalJSON(u.ModulrConfig, "", true)
	}

	if u.CurrencyCloudConfig != nil {
		return utils.MarshalJSON(u.CurrencyCloudConfig, "", true)
	}

	if u.BankingCircleConfig != nil {
		return utils.MarshalJSON(u.BankingCircleConfig, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
