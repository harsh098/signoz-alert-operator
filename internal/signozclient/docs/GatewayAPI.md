# \GatewayAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIngestionKey**](GatewayAPI.md#CreateIngestionKey) | **Post** /api/v2/gateway/ingestion_keys | Create ingestion key for workspace
[**CreateIngestionKeyLimit**](GatewayAPI.md#CreateIngestionKeyLimit) | **Post** /api/v2/gateway/ingestion_keys/{keyId}/limits | Create limit for the ingestion key
[**DeleteIngestionKey**](GatewayAPI.md#DeleteIngestionKey) | **Delete** /api/v2/gateway/ingestion_keys/{keyId} | Delete ingestion key for workspace
[**DeleteIngestionKeyLimit**](GatewayAPI.md#DeleteIngestionKeyLimit) | **Delete** /api/v2/gateway/ingestion_keys/limits/{limitId} | Delete limit for the ingestion key
[**GetIngestionKeys**](GatewayAPI.md#GetIngestionKeys) | **Get** /api/v2/gateway/ingestion_keys | Get ingestion keys for workspace
[**SearchIngestionKeys**](GatewayAPI.md#SearchIngestionKeys) | **Get** /api/v2/gateway/ingestion_keys/search | Search ingestion keys for workspace
[**UpdateIngestionKey**](GatewayAPI.md#UpdateIngestionKey) | **Patch** /api/v2/gateway/ingestion_keys/{keyId} | Update ingestion key for workspace
[**UpdateIngestionKeyLimit**](GatewayAPI.md#UpdateIngestionKeyLimit) | **Patch** /api/v2/gateway/ingestion_keys/limits/{limitId} | Update limit for the ingestion key



## CreateIngestionKey

> CreateIngestionKey201Response CreateIngestionKey(ctx).GatewaytypesPostableIngestionKey(gatewaytypesPostableIngestionKey).Execute()

Create ingestion key for workspace



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
	gatewaytypesPostableIngestionKey := *openapiclient.NewGatewaytypesPostableIngestionKey("Name_example") // GatewaytypesPostableIngestionKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.CreateIngestionKey(context.Background()).GatewaytypesPostableIngestionKey(gatewaytypesPostableIngestionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.CreateIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionKey`: CreateIngestionKey201Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.CreateIngestionKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewaytypesPostableIngestionKey** | [**GatewaytypesPostableIngestionKey**](GatewaytypesPostableIngestionKey.md) |  | 

### Return type

[**CreateIngestionKey201Response**](CreateIngestionKey201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIngestionKeyLimit

> CreateIngestionKeyLimit201Response CreateIngestionKeyLimit(ctx, keyId).GatewaytypesPostableIngestionKeyLimit(gatewaytypesPostableIngestionKeyLimit).Execute()

Create limit for the ingestion key



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
	keyId := "keyId_example" // string | 
	gatewaytypesPostableIngestionKeyLimit := *openapiclient.NewGatewaytypesPostableIngestionKeyLimit() // GatewaytypesPostableIngestionKeyLimit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.CreateIngestionKeyLimit(context.Background(), keyId).GatewaytypesPostableIngestionKeyLimit(gatewaytypesPostableIngestionKeyLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.CreateIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionKeyLimit`: CreateIngestionKeyLimit201Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.CreateIngestionKeyLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewaytypesPostableIngestionKeyLimit** | [**GatewaytypesPostableIngestionKeyLimit**](GatewaytypesPostableIngestionKeyLimit.md) |  | 

### Return type

[**CreateIngestionKeyLimit201Response**](CreateIngestionKeyLimit201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestionKey

> DeleteIngestionKey(ctx, keyId).Execute()

Delete ingestion key for workspace



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
	keyId := "keyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPI.DeleteIngestionKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.DeleteIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionKeyRequest struct via the builder pattern


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


## DeleteIngestionKeyLimit

> DeleteIngestionKeyLimit(ctx, limitId).Execute()

Delete limit for the ingestion key



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
	limitId := "limitId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPI.DeleteIngestionKeyLimit(context.Background(), limitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.DeleteIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionKeyLimitRequest struct via the builder pattern


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


## GetIngestionKeys

> GetIngestionKeys200Response GetIngestionKeys(ctx).Page(page).PerPage(perPage).Execute()

Get ingestion keys for workspace



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
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GetIngestionKeys(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GetIngestionKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngestionKeys`: GetIngestionKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GetIngestionKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GetIngestionKeys200Response**](GetIngestionKeys200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIngestionKeys

> GetIngestionKeys200Response SearchIngestionKeys(ctx).Name(name).Page(page).PerPage(perPage).Execute()

Search ingestion keys for workspace



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
	name := "name_example" // string | 
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.SearchIngestionKeys(context.Background()).Name(name).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.SearchIngestionKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIngestionKeys`: GetIngestionKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.SearchIngestionKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIngestionKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 

### Return type

[**GetIngestionKeys200Response**](GetIngestionKeys200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIngestionKey

> UpdateIngestionKey(ctx, keyId).GatewaytypesPostableIngestionKey(gatewaytypesPostableIngestionKey).Execute()

Update ingestion key for workspace



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
	keyId := "keyId_example" // string | 
	gatewaytypesPostableIngestionKey := *openapiclient.NewGatewaytypesPostableIngestionKey("Name_example") // GatewaytypesPostableIngestionKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPI.UpdateIngestionKey(context.Background(), keyId).GatewaytypesPostableIngestionKey(gatewaytypesPostableIngestionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.UpdateIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewaytypesPostableIngestionKey** | [**GatewaytypesPostableIngestionKey**](GatewaytypesPostableIngestionKey.md) |  | 

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


## UpdateIngestionKeyLimit

> UpdateIngestionKeyLimit(ctx, limitId).GatewaytypesUpdatableIngestionKeyLimit(gatewaytypesUpdatableIngestionKeyLimit).Execute()

Update limit for the ingestion key



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
	limitId := "limitId_example" // string | 
	gatewaytypesUpdatableIngestionKeyLimit := *openapiclient.NewGatewaytypesUpdatableIngestionKeyLimit(*openapiclient.NewGatewaytypesLimitConfig()) // GatewaytypesUpdatableIngestionKeyLimit |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPI.UpdateIngestionKeyLimit(context.Background(), limitId).GatewaytypesUpdatableIngestionKeyLimit(gatewaytypesUpdatableIngestionKeyLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.UpdateIngestionKeyLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestionKeyLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewaytypesUpdatableIngestionKeyLimit** | [**GatewaytypesUpdatableIngestionKeyLimit**](GatewaytypesUpdatableIngestionKeyLimit.md) |  | 

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

