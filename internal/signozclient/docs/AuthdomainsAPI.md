# \AuthdomainsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthDomain**](AuthdomainsAPI.md#CreateAuthDomain) | **Post** /api/v1/domains | Create auth domain
[**DeleteAuthDomain**](AuthdomainsAPI.md#DeleteAuthDomain) | **Delete** /api/v1/domains/{id} | Delete auth domain
[**GetAuthDomain**](AuthdomainsAPI.md#GetAuthDomain) | **Get** /api/v1/domains/{id} | Get auth domain by ID
[**ListAuthDomains**](AuthdomainsAPI.md#ListAuthDomains) | **Get** /api/v1/domains | List all auth domains
[**UpdateAuthDomain**](AuthdomainsAPI.md#UpdateAuthDomain) | **Put** /api/v1/domains/{id} | Update auth domain



## CreateAuthDomain

> CreatePublicDashboard201Response CreateAuthDomain(ctx).AuthtypesPostableAuthDomain(authtypesPostableAuthDomain).Execute()

Create auth domain



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
	authtypesPostableAuthDomain := *openapiclient.NewAuthtypesPostableAuthDomain() // AuthtypesPostableAuthDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthdomainsAPI.CreateAuthDomain(context.Background()).AuthtypesPostableAuthDomain(authtypesPostableAuthDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthdomainsAPI.CreateAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthDomain`: CreatePublicDashboard201Response
	fmt.Fprintf(os.Stdout, "Response from `AuthdomainsAPI.CreateAuthDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authtypesPostableAuthDomain** | [**AuthtypesPostableAuthDomain**](AuthtypesPostableAuthDomain.md) |  | 

### Return type

[**CreatePublicDashboard201Response**](CreatePublicDashboard201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthDomain

> DeleteAuthDomain(ctx, id).Execute()

Delete auth domain



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthdomainsAPI.DeleteAuthDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthdomainsAPI.DeleteAuthDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthDomainRequest struct via the builder pattern


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


## GetAuthDomain

> GetAuthDomain200Response GetAuthDomain(ctx, id).Execute()

Get auth domain by ID



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthdomainsAPI.GetAuthDomain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthdomainsAPI.GetAuthDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthDomain`: GetAuthDomain200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthdomainsAPI.GetAuthDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAuthDomain200Response**](GetAuthDomain200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthDomains

> ListAuthDomains200Response ListAuthDomains(ctx).Execute()

List all auth domains



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
	resp, r, err := apiClient.AuthdomainsAPI.ListAuthDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthdomainsAPI.ListAuthDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuthDomains`: ListAuthDomains200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthdomainsAPI.ListAuthDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthDomainsRequest struct via the builder pattern


### Return type

[**ListAuthDomains200Response**](ListAuthDomains200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthDomain

> UpdateAuthDomain(ctx, id).AuthtypesUpdatableAuthDomain(authtypesUpdatableAuthDomain).Execute()

Update auth domain



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
	id := "id_example" // string | 
	authtypesUpdatableAuthDomain := *openapiclient.NewAuthtypesUpdatableAuthDomain() // AuthtypesUpdatableAuthDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthdomainsAPI.UpdateAuthDomain(context.Background(), id).AuthtypesUpdatableAuthDomain(authtypesUpdatableAuthDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthdomainsAPI.UpdateAuthDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateAuthDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authtypesUpdatableAuthDomain** | [**AuthtypesUpdatableAuthDomain**](AuthtypesUpdatableAuthDomain.md) |  | 

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

