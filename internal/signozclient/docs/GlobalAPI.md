# \GlobalAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalConfig**](GlobalAPI.md#GetGlobalConfig) | **Get** /api/v1/global/config | Get global config



## GetGlobalConfig

> GetGlobalConfig200Response GetGlobalConfig(ctx).Execute()

Get global config



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalAPI.GetGlobalConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalAPI.GetGlobalConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalConfig`: GetGlobalConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalAPI.GetGlobalConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalConfigRequest struct via the builder pattern


### Return type

[**GetGlobalConfig200Response**](GetGlobalConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

