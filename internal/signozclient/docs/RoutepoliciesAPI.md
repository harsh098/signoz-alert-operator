# \RoutepoliciesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoutePolicy**](RoutepoliciesAPI.md#CreateRoutePolicy) | **Post** /api/v1/route_policies | Create route policy
[**DeleteRoutePolicyByID**](RoutepoliciesAPI.md#DeleteRoutePolicyByID) | **Delete** /api/v1/route_policies/{id} | Delete route policy
[**GetAllRoutePolicies**](RoutepoliciesAPI.md#GetAllRoutePolicies) | **Get** /api/v1/route_policies | List route policies
[**GetRoutePolicyByID**](RoutepoliciesAPI.md#GetRoutePolicyByID) | **Get** /api/v1/route_policies/{id} | Get route policy by ID
[**UpdateRoutePolicy**](RoutepoliciesAPI.md#UpdateRoutePolicy) | **Put** /api/v1/route_policies/{id} | Update route policy



## CreateRoutePolicy

> CreateRoutePolicy201Response CreateRoutePolicy(ctx).AlertmanagertypesPostableRoutePolicy(alertmanagertypesPostableRoutePolicy).Execute()

Create route policy



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
	alertmanagertypesPostableRoutePolicy := *openapiclient.NewAlertmanagertypesPostableRoutePolicy([]string{"Channels_example"}, "Expression_example", "Name_example") // AlertmanagertypesPostableRoutePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutepoliciesAPI.CreateRoutePolicy(context.Background()).AlertmanagertypesPostableRoutePolicy(alertmanagertypesPostableRoutePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutepoliciesAPI.CreateRoutePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoutePolicy`: CreateRoutePolicy201Response
	fmt.Fprintf(os.Stdout, "Response from `RoutepoliciesAPI.CreateRoutePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertmanagertypesPostableRoutePolicy** | [**AlertmanagertypesPostableRoutePolicy**](AlertmanagertypesPostableRoutePolicy.md) |  | 

### Return type

[**CreateRoutePolicy201Response**](CreateRoutePolicy201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutePolicyByID

> DeleteRoutePolicyByID(ctx, id).Execute()

Delete route policy



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutepoliciesAPI.DeleteRoutePolicyByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutepoliciesAPI.DeleteRoutePolicyByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutePolicyByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoutePolicies

> GetAllRoutePolicies200Response GetAllRoutePolicies(ctx).Execute()

List route policies



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
	resp, r, err := apiClient.RoutepoliciesAPI.GetAllRoutePolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutepoliciesAPI.GetAllRoutePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRoutePolicies`: GetAllRoutePolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutepoliciesAPI.GetAllRoutePolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRoutePoliciesRequest struct via the builder pattern


### Return type

[**GetAllRoutePolicies200Response**](GetAllRoutePolicies200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutePolicyByID

> CreateRoutePolicy201Response GetRoutePolicyByID(ctx, id).Execute()

Get route policy by ID



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutepoliciesAPI.GetRoutePolicyByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutepoliciesAPI.GetRoutePolicyByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoutePolicyByID`: CreateRoutePolicy201Response
	fmt.Fprintf(os.Stdout, "Response from `RoutepoliciesAPI.GetRoutePolicyByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutePolicyByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRoutePolicy201Response**](CreateRoutePolicy201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutePolicy

> CreateRoutePolicy201Response UpdateRoutePolicy(ctx, id).AlertmanagertypesPostableRoutePolicy(alertmanagertypesPostableRoutePolicy).Execute()

Update route policy



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
	id := "id_example" // string | 
	alertmanagertypesPostableRoutePolicy := *openapiclient.NewAlertmanagertypesPostableRoutePolicy([]string{"Channels_example"}, "Expression_example", "Name_example") // AlertmanagertypesPostableRoutePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutepoliciesAPI.UpdateRoutePolicy(context.Background(), id).AlertmanagertypesPostableRoutePolicy(alertmanagertypesPostableRoutePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutepoliciesAPI.UpdateRoutePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoutePolicy`: CreateRoutePolicy201Response
	fmt.Fprintf(os.Stdout, "Response from `RoutepoliciesAPI.UpdateRoutePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertmanagertypesPostableRoutePolicy** | [**AlertmanagertypesPostableRoutePolicy**](AlertmanagertypesPostableRoutePolicy.md) |  | 

### Return type

[**CreateRoutePolicy201Response**](CreateRoutePolicy201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

