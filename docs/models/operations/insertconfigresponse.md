# InsertConfigResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ConfigResponse`                                                | [*shared.ConfigResponse](../../models/shared/configresponse.md) | :heavy_minus_sign:                                              | Config created successfully.                                    |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |
| `InsertConfig400TextPlainString`                                | **string*                                                       | :heavy_minus_sign:                                              | Bad Request                                                     |