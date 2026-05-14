# \QuerierAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryRangeV5**](QuerierAPI.md#QueryRangeV5) | **Post** /api/v5/query_range | Query range
[**ReplaceVariables**](QuerierAPI.md#ReplaceVariables) | **Post** /api/v5/substitute_vars | Replace variables



## QueryRangeV5

> GetPublicDashboardWidgetQueryRange200Response QueryRangeV5(ctx).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()

Query range



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
	querybuildertypesv5QueryRangeRequest := *openapiclient.NewQuerybuildertypesv5QueryRangeRequest() // Querybuildertypesv5QueryRangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuerierAPI.QueryRangeV5(context.Background()).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuerierAPI.QueryRangeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRangeV5`: GetPublicDashboardWidgetQueryRange200Response
	fmt.Fprintf(os.Stdout, "Response from `QuerierAPI.QueryRangeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRangeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querybuildertypesv5QueryRangeRequest** | [**Querybuildertypesv5QueryRangeRequest**](Querybuildertypesv5QueryRangeRequest.md) |  | 

### Return type

[**GetPublicDashboardWidgetQueryRange200Response**](GetPublicDashboardWidgetQueryRange200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceVariables

> ReplaceVariables200Response ReplaceVariables(ctx).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()

Replace variables



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
	querybuildertypesv5QueryRangeRequest := *openapiclient.NewQuerybuildertypesv5QueryRangeRequest() // Querybuildertypesv5QueryRangeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuerierAPI.ReplaceVariables(context.Background()).Querybuildertypesv5QueryRangeRequest(querybuildertypesv5QueryRangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuerierAPI.ReplaceVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceVariables`: ReplaceVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `QuerierAPI.ReplaceVariables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **querybuildertypesv5QueryRangeRequest** | [**Querybuildertypesv5QueryRangeRequest**](Querybuildertypesv5QueryRangeRequest.md) |  | 

### Return type

[**ReplaceVariables200Response**](ReplaceVariables200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

