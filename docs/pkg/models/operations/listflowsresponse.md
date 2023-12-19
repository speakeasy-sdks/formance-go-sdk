# ListFlowsResponse


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ContentType`                                                                        | *string*                                                                             | :heavy_check_mark:                                                                   | HTTP response content type for this operation                                        |
| `Error`                                                                              | [*shared.Error](../../../pkg/models/shared/error.md)                                 | :heavy_minus_sign:                                                                   | General error                                                                        |
| `ListWorkflowsResponse`                                                              | [*shared.ListWorkflowsResponse](../../../pkg/models/shared/listworkflowsresponse.md) | :heavy_minus_sign:                                                                   | List of workflows                                                                    |
| `StatusCode`                                                                         | *int*                                                                                | :heavy_check_mark:                                                                   | HTTP response status code for this operation                                         |
| `RawResponse`                                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)                               | :heavy_check_mark:                                                                   | Raw HTTP response; suitable for custom response parsing                              |