# \ServiceaccountAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceaccountAPI.md#CreateServiceAccount) | **Post** /api/v1/service_accounts | Create service account
[**CreateServiceAccountKey**](ServiceaccountAPI.md#CreateServiceAccountKey) | **Post** /api/v1/service_accounts/{id}/keys | Create a service account key
[**CreateServiceAccountRole**](ServiceaccountAPI.md#CreateServiceAccountRole) | **Post** /api/v1/service_accounts/{id}/roles | Create service account role
[**DeleteServiceAccount**](ServiceaccountAPI.md#DeleteServiceAccount) | **Delete** /api/v1/service_accounts/{id} | Deletes a service account
[**DeleteServiceAccountRole**](ServiceaccountAPI.md#DeleteServiceAccountRole) | **Delete** /api/v1/service_accounts/{id}/roles/{rid} | Delete service account role
[**GetMyServiceAccount**](ServiceaccountAPI.md#GetMyServiceAccount) | **Get** /api/v1/service_accounts/me | Gets my service account
[**GetServiceAccount**](ServiceaccountAPI.md#GetServiceAccount) | **Get** /api/v1/service_accounts/{id} | Gets a service account
[**GetServiceAccountRoles**](ServiceaccountAPI.md#GetServiceAccountRoles) | **Get** /api/v1/service_accounts/{id}/roles | Gets service account roles
[**ListServiceAccountKeys**](ServiceaccountAPI.md#ListServiceAccountKeys) | **Get** /api/v1/service_accounts/{id}/keys | List service account keys
[**ListServiceAccounts**](ServiceaccountAPI.md#ListServiceAccounts) | **Get** /api/v1/service_accounts | List service accounts
[**RevokeServiceAccountKey**](ServiceaccountAPI.md#RevokeServiceAccountKey) | **Delete** /api/v1/service_accounts/{id}/keys/{fid} | Revoke a service account key
[**UpdateMyServiceAccount**](ServiceaccountAPI.md#UpdateMyServiceAccount) | **Put** /api/v1/service_accounts/me | Updates my service account
[**UpdateServiceAccount**](ServiceaccountAPI.md#UpdateServiceAccount) | **Put** /api/v1/service_accounts/{id} | Updates a service account
[**UpdateServiceAccountKey**](ServiceaccountAPI.md#UpdateServiceAccountKey) | **Put** /api/v1/service_accounts/{id}/keys/{fid} | Updates a service account key



## CreateServiceAccount

> CreatePublicDashboard201Response CreateServiceAccount(ctx).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()

Create service account



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
	serviceaccounttypesPostableServiceAccount := *openapiclient.NewServiceaccounttypesPostableServiceAccount("Name_example") // ServiceaccounttypesPostableServiceAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.CreateServiceAccount(context.Background()).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccount`: CreatePublicDashboard201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceaccounttypesPostableServiceAccount** | [**ServiceaccounttypesPostableServiceAccount**](ServiceaccounttypesPostableServiceAccount.md) |  | 

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


## CreateServiceAccountKey

> CreateServiceAccountKey201Response CreateServiceAccountKey(ctx, id).ServiceaccounttypesPostableFactorAPIKey(serviceaccounttypesPostableFactorAPIKey).Execute()

Create a service account key



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
	serviceaccounttypesPostableFactorAPIKey := *openapiclient.NewServiceaccounttypesPostableFactorAPIKey(int32(123), "Name_example") // ServiceaccounttypesPostableFactorAPIKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.CreateServiceAccountKey(context.Background(), id).ServiceaccounttypesPostableFactorAPIKey(serviceaccounttypesPostableFactorAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.CreateServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountKey`: CreateServiceAccountKey201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.CreateServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceaccounttypesPostableFactorAPIKey** | [**ServiceaccounttypesPostableFactorAPIKey**](ServiceaccounttypesPostableFactorAPIKey.md) |  | 

### Return type

[**CreateServiceAccountKey201Response**](CreateServiceAccountKey201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccountRole

> CreatePublicDashboard201Response CreateServiceAccountRole(ctx, id).ServiceaccounttypesPostableServiceAccountRole(serviceaccounttypesPostableServiceAccountRole).Execute()

Create service account role



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
	serviceaccounttypesPostableServiceAccountRole := *openapiclient.NewServiceaccounttypesPostableServiceAccountRole("Id_example") // ServiceaccounttypesPostableServiceAccountRole |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.CreateServiceAccountRole(context.Background(), id).ServiceaccounttypesPostableServiceAccountRole(serviceaccounttypesPostableServiceAccountRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.CreateServiceAccountRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccountRole`: CreatePublicDashboard201Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.CreateServiceAccountRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceaccounttypesPostableServiceAccountRole** | [**ServiceaccounttypesPostableServiceAccountRole**](ServiceaccounttypesPostableServiceAccountRole.md) |  | 

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


## DeleteServiceAccount

> string DeleteServiceAccount(ctx, id).Execute()

Deletes a service account



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
	resp, r, err := apiClient.ServiceaccountAPI.DeleteServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.DeleteServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceAccount`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.DeleteServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccountRole

> string DeleteServiceAccountRole(ctx, id, rid).Execute()

Delete service account role



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
	rid := "rid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.DeleteServiceAccountRole(context.Background(), id, rid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.DeleteServiceAccountRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceAccountRole`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.DeleteServiceAccountRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyServiceAccount

> GetServiceAccount200Response GetMyServiceAccount(ctx).Execute()

Gets my service account



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
	resp, r, err := apiClient.ServiceaccountAPI.GetMyServiceAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.GetMyServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyServiceAccount`: GetServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.GetMyServiceAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyServiceAccountRequest struct via the builder pattern


### Return type

[**GetServiceAccount200Response**](GetServiceAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccount

> GetServiceAccount200Response GetServiceAccount(ctx, id).Execute()

Gets a service account



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
	resp, r, err := apiClient.ServiceaccountAPI.GetServiceAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.GetServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccount`: GetServiceAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.GetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceAccount200Response**](GetServiceAccount200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountRoles

> GetServiceAccountRoles200Response GetServiceAccountRoles(ctx, id).Execute()

Gets service account roles



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
	resp, r, err := apiClient.ServiceaccountAPI.GetServiceAccountRoles(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.GetServiceAccountRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountRoles`: GetServiceAccountRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.GetServiceAccountRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceAccountRoles200Response**](GetServiceAccountRoles200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccountKeys

> ListServiceAccountKeys200Response ListServiceAccountKeys(ctx, id).Execute()

List service account keys



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
	resp, r, err := apiClient.ServiceaccountAPI.ListServiceAccountKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.ListServiceAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccountKeys`: ListServiceAccountKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.ListServiceAccountKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListServiceAccountKeys200Response**](ListServiceAccountKeys200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> ListServiceAccounts200Response ListServiceAccounts(ctx).Execute()

List service accounts



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
	resp, r, err := apiClient.ServiceaccountAPI.ListServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccounts`: ListServiceAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


### Return type

[**ListServiceAccounts200Response**](ListServiceAccounts200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeServiceAccountKey

> string RevokeServiceAccountKey(ctx, id, fid).Execute()

Revoke a service account key



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
	fid := "fid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.RevokeServiceAccountKey(context.Background(), id, fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.RevokeServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeServiceAccountKey`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.RevokeServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMyServiceAccount

> string UpdateMyServiceAccount(ctx).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()

Updates my service account



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
	serviceaccounttypesPostableServiceAccount := *openapiclient.NewServiceaccounttypesPostableServiceAccount("Name_example") // ServiceaccounttypesPostableServiceAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.UpdateMyServiceAccount(context.Background()).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.UpdateMyServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMyServiceAccount`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.UpdateMyServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMyServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceaccounttypesPostableServiceAccount** | [**ServiceaccounttypesPostableServiceAccount**](ServiceaccounttypesPostableServiceAccount.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccount

> string UpdateServiceAccount(ctx, id).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()

Updates a service account



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
	serviceaccounttypesPostableServiceAccount := *openapiclient.NewServiceaccounttypesPostableServiceAccount("Name_example") // ServiceaccounttypesPostableServiceAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.UpdateServiceAccount(context.Background(), id).ServiceaccounttypesPostableServiceAccount(serviceaccounttypesPostableServiceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.UpdateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceAccount`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.UpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceaccounttypesPostableServiceAccount** | [**ServiceaccounttypesPostableServiceAccount**](ServiceaccounttypesPostableServiceAccount.md) |  | 

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


## UpdateServiceAccountKey

> string UpdateServiceAccountKey(ctx, id, fid).ServiceaccounttypesUpdatableFactorAPIKey(serviceaccounttypesUpdatableFactorAPIKey).Execute()

Updates a service account key



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
	fid := "fid_example" // string | 
	serviceaccounttypesUpdatableFactorAPIKey := *openapiclient.NewServiceaccounttypesUpdatableFactorAPIKey(int32(123), "Name_example") // ServiceaccounttypesUpdatableFactorAPIKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceaccountAPI.UpdateServiceAccountKey(context.Background(), id, fid).ServiceaccounttypesUpdatableFactorAPIKey(serviceaccounttypesUpdatableFactorAPIKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountAPI.UpdateServiceAccountKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceAccountKey`: string
	fmt.Fprintf(os.Stdout, "Response from `ServiceaccountAPI.UpdateServiceAccountKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceaccounttypesUpdatableFactorAPIKey** | [**ServiceaccounttypesUpdatableFactorAPIKey**](ServiceaccounttypesUpdatableFactorAPIKey.md) |  | 

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

