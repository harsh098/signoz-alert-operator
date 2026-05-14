# \LogsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleExportRawDataPOST**](LogsAPI.md#HandleExportRawDataPOST) | **Post** /api/v1/export_raw_data | Export raw data
[**HandlePromoteAndIndexPaths**](LogsAPI.md#HandlePromoteAndIndexPaths) | **Post** /api/v1/logs/promote_paths | Promote and index paths
[**ListPromotedAndIndexedPaths**](LogsAPI.md#ListPromotedAndIndexedPaths) | **Get** /api/v1/logs/promote_paths | Promote and index paths



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	format := "format_example" // string | The output format for the export. (optional) (default to "csv")
	querybuildertypesv5QueryRangeRequest := *openapiclient.NewQuerybuildertypesv5QueryRangeRequest() // Querybuildertypesv5QueryRangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.HandleExportRawDataPOST(context.Background()).Format(format).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.HandleExportRawDataPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandleExportRawDataPOST`: string
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.HandleExportRawDataPOST`: %v\n", resp)
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


## HandlePromoteAndIndexPaths

> HandlePromoteAndIndexPaths(ctx).PromotetypesPromotePath(promotetypesPromotePath).Execute()

Promote and index paths



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	promotetypesPromotePath := []openapiclient.PromotetypesPromotePath{*openapiclient.NewPromotetypesPromotePath()} // []PromotetypesPromotePath |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogsAPI.HandlePromoteAndIndexPaths(context.Background()).PromotetypesPromotePath(promotetypesPromotePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.HandlePromoteAndIndexPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandlePromoteAndIndexPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promotetypesPromotePath** | [**[]PromotetypesPromotePath**](PromotetypesPromotePath.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPromotedAndIndexedPaths

> ListPromotedAndIndexedPaths200Response ListPromotedAndIndexedPaths(ctx).Execute()

Promote and index paths



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.ListPromotedAndIndexedPaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ListPromotedAndIndexedPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPromotedAndIndexedPaths`: ListPromotedAndIndexedPaths200Response
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.ListPromotedAndIndexedPaths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPromotedAndIndexedPathsRequest struct via the builder pattern


### Return type

[**ListPromotedAndIndexedPaths200Response**](ListPromotedAndIndexedPaths200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

