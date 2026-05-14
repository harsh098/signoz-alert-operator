# \TracesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleExportRawDataPOST**](TracesAPI.md#HandleExportRawDataPOST) | **Post** /api/v1/export_raw_data | Export raw data



## HandleExportRawDataPOST

> string HandleExportRawDataPOST(ctx).Format(format).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()

Export raw data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harsh098/signoz-alert-operator/internal/signozclient"
)

func main() {
	format := "format_example" // string | The output format for the export. (optional) (default to "csv")
	querybuildertypesv5QueryRangeRequest := *openapiclient.NewQuerybuildertypesv5QueryRangeRequest() // Querybuildertypesv5QueryRangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.HandleExportRawDataPOST(context.Background()).Format(format).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.HandleExportRawDataPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleExportRawDataPOST`: string
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.HandleExportRawDataPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleExportRawDataPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | The output format for the export. | [default to &quot;csv&quot;]
 **querybuildertypesv5QueryRangeRequest** | [**Querybuildertypesv5QueryRangeRequest**](Querybuildertypesv5QueryRangeRequest.md) |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

