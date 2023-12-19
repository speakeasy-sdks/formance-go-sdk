# TaskStripe


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ConnectorID`                                                                     | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               |
| `CreatedAt`                                                                       | [time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_check_mark:                                                                | N/A                                                                               |
| `Descriptor`                                                                      | [shared.TaskStripeDescriptor](../../../pkg/models/shared/taskstripedescriptor.md) | :heavy_check_mark:                                                                | N/A                                                                               |
| `Error`                                                                           | **string*                                                                         | :heavy_minus_sign:                                                                | N/A                                                                               |
| `ID`                                                                              | *string*                                                                          | :heavy_check_mark:                                                                | N/A                                                                               |
| `State`                                                                           | [shared.TaskStripeState](../../../pkg/models/shared/taskstripestate.md)           | :heavy_check_mark:                                                                | N/A                                                                               |
| `Status`                                                                          | [shared.PaymentStatus](../../../pkg/models/shared/paymentstatus.md)               | :heavy_check_mark:                                                                | N/A                                                                               |
| `UpdatedAt`                                                                       | [time.Time](https://pkg.go.dev/time#Time)                                         | :heavy_check_mark:                                                                | N/A                                                                               |