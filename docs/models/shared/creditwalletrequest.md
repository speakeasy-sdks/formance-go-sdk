# CreditWalletRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Amount`                                           | [shared.Monetary](../../models/shared/monetary.md) | :heavy_check_mark:                                 | N/A                                                |
| `Balance`                                          | **string*                                          | :heavy_minus_sign:                                 | The balance to credit                              |
| `Metadata`                                         | map[string]*interface{}*                           | :heavy_minus_sign:                                 | Metadata associated with the wallet.               |
| `Reference`                                        | **string*                                          | :heavy_minus_sign:                                 | N/A                                                |
| `Sources`                                          | [][shared.Subject](../../models/shared/subject.md) | :heavy_check_mark:                                 | N/A                                                |