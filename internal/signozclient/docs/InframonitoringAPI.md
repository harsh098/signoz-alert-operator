# \InframonitoringAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHosts**](InframonitoringAPI.md#ListHosts) | **Post** /api/v2/infra_monitoring/hosts | List Hosts for Infra Monitoring
[**ListPods**](InframonitoringAPI.md#ListPods) | **Post** /api/v2/infra_monitoring/pods | List Pods for Infra Monitoring



## ListHosts

> ListHosts200Response ListHosts(ctx).InframonitoringtypesPostableHosts(inframonitoringtypesPostableHosts).Execute()

List Hosts for Infra Monitoring



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
	inframonitoringtypesPostableHosts := *openapiclient.NewInframonitoringtypesPostableHosts(int64(123), int32(123), int64(123)) // InframonitoringtypesPostableHosts |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InframonitoringAPI.ListHosts(context.Background()).InframonitoringtypesPostableHosts(inframonitoringtypesPostableHosts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InframonitoringAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: ListHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `InframonitoringAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inframonitoringtypesPostableHosts** | [**InframonitoringtypesPostableHosts**](InframonitoringtypesPostableHosts.md) |  | 

### Return type

[**ListHosts200Response**](ListHosts200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPods

> ListPods200Response ListPods(ctx).InframonitoringtypesPostablePods(inframonitoringtypesPostablePods).Execute()

List Pods for Infra Monitoring



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
	inframonitoringtypesPostablePods := *openapiclient.NewInframonitoringtypesPostablePods(int64(123), int32(123), int64(123)) // InframonitoringtypesPostablePods |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InframonitoringAPI.ListPods(context.Background()).InframonitoringtypesPostablePods(inframonitoringtypesPostablePods).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InframonitoringAPI.ListPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPods`: ListPods200Response
	fmt.Fprintf(os.Stdout, "Response from `InframonitoringAPI.ListPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inframonitoringtypesPostablePods** | [**InframonitoringtypesPostablePods**](InframonitoringtypesPostablePods.md) |  | 

### Return type

[**ListPods200Response**](ListPods200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

