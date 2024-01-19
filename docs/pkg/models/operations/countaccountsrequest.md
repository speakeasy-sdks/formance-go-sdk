# CountAccountsRequest


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Ledger`                                                                                              | *string*                                                                                              | :heavy_check_mark:                                                                                    | Name of the ledger.                                                                                   | ledger001                                                                                             |
| `Address`                                                                                             | **string*                                                                                             | :heavy_minus_sign:                                                                                    | Filter accounts by address pattern (regular expression placed between ^ and $).                       | users:.+                                                                                              |
| `Metadata`                                                                                            | [*operations.Metadata](../../../pkg/models/operations/metadata.md)                                    | :heavy_minus_sign:                                                                                    | Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below. |                                                                                                       |