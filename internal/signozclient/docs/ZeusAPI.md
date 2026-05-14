# \ZeusAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHosts**](ZeusAPI.md#GetHosts) | **Get** /api/v2/zeus/hosts | Get host info from Zeus.
[**PutHost**](ZeusAPI.md#PutHost) | **Put** /api/v2/zeus/hosts | Put host in Zeus for a deployment.
[**PutProfile**](ZeusAPI.md#PutProfile) | **Put** /api/v2/zeus/profiles | Put profile in Zeus for a deployment.



## GetHosts

> GetHosts200Response GetHosts(ctx).Execute()

Get host info from Zeus.



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
	resp, r, err := apiClient.ZeusAPI.GetHosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZeusAPI.GetHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosts`: GetHosts200Response
	fmt.Fprintf(os.Stdout, "Response from `ZeusAPI.GetHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


### Return type

[**GetHosts200Response**](GetHosts200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutHost

> PutHost(ctx).ZeustypesPostableHost(zeustypesPostableHost).Execute()

Put host in Zeus for a deployment.



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
	zeustypesPostableHost := *openapiclient.NewZeustypesPostableHost("Name_example") // ZeustypesPostableHost |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZeusAPI.PutHost(context.Background()).ZeustypesPostableHost(zeustypesPostableHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZeusAPI.PutHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zeustypesPostableHost** | [**ZeustypesPostableHost**](ZeustypesPostableHost.md) |  | 

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


## PutProfile

> PutProfile(ctx).ZeustypesPostableProfile(zeustypesPostableProfile).Execute()

Put profile in Zeus for a deployment.



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
	zeustypesPostableProfile := *openapiclient.NewZeustypesPostableProfile("ExistingObservabilityTool_example", false, int64(123), int64(123), int64(123), []string{"ReasonsForInterestInSignoz_example"}, "TimelineForMigratingToSignoz_example", false, "WhereDidYouDiscoverSignoz_example") // ZeustypesPostableProfile |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZeusAPI.PutProfile(context.Background()).ZeustypesPostableProfile(zeustypesPostableProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZeusAPI.PutProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zeustypesPostableProfile** | [**ZeustypesPostableProfile**](ZeustypesPostableProfile.md) |  | 

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

