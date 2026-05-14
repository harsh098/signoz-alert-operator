# \PreferencesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgPreference**](PreferencesAPI.md#GetOrgPreference) | **Get** /api/v1/org/preferences/{name} | Get org preference
[**GetUserPreference**](PreferencesAPI.md#GetUserPreference) | **Get** /api/v1/user/preferences/{name} | Get user preference
[**ListOrgPreferences**](PreferencesAPI.md#ListOrgPreferences) | **Get** /api/v1/org/preferences | List org preferences
[**ListUserPreferences**](PreferencesAPI.md#ListUserPreferences) | **Get** /api/v1/user/preferences | List user preferences
[**UpdateOrgPreference**](PreferencesAPI.md#UpdateOrgPreference) | **Put** /api/v1/org/preferences/{name} | Update org preference
[**UpdateUserPreference**](PreferencesAPI.md#UpdateUserPreference) | **Put** /api/v1/user/preferences/{name} | Update user preference



## GetOrgPreference

> GetOrgPreference200Response GetOrgPreference(ctx, name).Execute()

Get org preference



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreferencesAPI.GetOrgPreference(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.GetOrgPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgPreference`: GetOrgPreference200Response
	fmt.Fprintf(os.Stdout, "Response from `PreferencesAPI.GetOrgPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrgPreference200Response**](GetOrgPreference200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPreference

> GetOrgPreference200Response GetUserPreference(ctx, name).Execute()

Get user preference



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreferencesAPI.GetUserPreference(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.GetUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPreference`: GetOrgPreference200Response
	fmt.Fprintf(os.Stdout, "Response from `PreferencesAPI.GetUserPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrgPreference200Response**](GetOrgPreference200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPreferences

> ListOrgPreferences200Response ListOrgPreferences(ctx).Execute()

List org preferences



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
	resp, r, err := apiClient.PreferencesAPI.ListOrgPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.ListOrgPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPreferences`: ListOrgPreferences200Response
	fmt.Fprintf(os.Stdout, "Response from `PreferencesAPI.ListOrgPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPreferencesRequest struct via the builder pattern


### Return type

[**ListOrgPreferences200Response**](ListOrgPreferences200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserPreferences

> ListOrgPreferences200Response ListUserPreferences(ctx).Execute()

List user preferences



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
	resp, r, err := apiClient.PreferencesAPI.ListUserPreferences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.ListUserPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserPreferences`: ListOrgPreferences200Response
	fmt.Fprintf(os.Stdout, "Response from `PreferencesAPI.ListUserPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserPreferencesRequest struct via the builder pattern


### Return type

[**ListOrgPreferences200Response**](ListOrgPreferences200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgPreference

> UpdateOrgPreference(ctx, name).PreferencetypesUpdatablePreference(preferencetypesUpdatablePreference).Execute()

Update org preference



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
	preferencetypesUpdatablePreference := *openapiclient.NewPreferencetypesUpdatablePreference() // PreferencetypesUpdatablePreference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreferencesAPI.UpdateOrgPreference(context.Background(), name).PreferencetypesUpdatablePreference(preferencetypesUpdatablePreference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.UpdateOrgPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preferencetypesUpdatablePreference** | [**PreferencetypesUpdatablePreference**](PreferencetypesUpdatablePreference.md) |  | 

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


## UpdateUserPreference

> UpdateUserPreference(ctx, name).PreferencetypesUpdatablePreference(preferencetypesUpdatablePreference).Execute()

Update user preference



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
	preferencetypesUpdatablePreference := *openapiclient.NewPreferencetypesUpdatablePreference() // PreferencetypesUpdatablePreference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PreferencesAPI.UpdateUserPreference(context.Background(), name).PreferencetypesUpdatablePreference(preferencetypesUpdatablePreference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesAPI.UpdateUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preferencetypesUpdatablePreference** | [**PreferencetypesUpdatablePreference**](PreferencetypesUpdatablePreference.md) |  | 

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

