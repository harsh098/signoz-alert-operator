# \TracedetailAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWaterfall**](TracedetailAPI.md#GetWaterfall) | **Post** /api/v3/traces/{traceID}/waterfall | Get waterfall view for a trace



## GetWaterfall

> GetWaterfall200Response GetWaterfall(ctx, traceID).TracedetailtypesPostableWaterfall(tracedetailtypesPostableWaterfall).Execute()

Get waterfall view for a trace



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
	traceID := "traceID_example" // string | 
	tracedetailtypesPostableWaterfall := *openapiclient.NewTracedetailtypesPostableWaterfall() // TracedetailtypesPostableWaterfall |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracedetailAPI.GetWaterfall(context.Background(), traceID).TracedetailtypesPostableWaterfall(tracedetailtypesPostableWaterfall).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracedetailAPI.GetWaterfall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWaterfall`: GetWaterfall200Response
	fmt.Fprintf(os.Stdout, "Response from `TracedetailAPI.GetWaterfall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWaterfallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tracedetailtypesPostableWaterfall** | [**TracedetailtypesPostableWaterfall**](TracedetailtypesPostableWaterfall.md) |  | 

### Return type

[**GetWaterfall200Response**](GetWaterfall200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

