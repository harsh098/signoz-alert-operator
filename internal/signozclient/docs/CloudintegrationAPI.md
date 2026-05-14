# \CloudintegrationAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCheckIn**](CloudintegrationAPI.md#AgentCheckIn) | **Post** /api/v1/cloud_integrations/{cloud_provider}/accounts/check_in | Agent check-in
[**AgentCheckInDeprecated**](CloudintegrationAPI.md#AgentCheckInDeprecated) | **Post** /api/v1/cloud-integrations/{cloud_provider}/agent-check-in | Agent check-in
[**CreateAccount**](CloudintegrationAPI.md#CreateAccount) | **Post** /api/v1/cloud_integrations/{cloud_provider}/accounts | Create account
[**DisconnectAccount**](CloudintegrationAPI.md#DisconnectAccount) | **Delete** /api/v1/cloud_integrations/{cloud_provider}/accounts/{id} | Disconnect account
[**GetAccount**](CloudintegrationAPI.md#GetAccount) | **Get** /api/v1/cloud_integrations/{cloud_provider}/accounts/{id} | Get account
[**GetConnectionCredentials**](CloudintegrationAPI.md#GetConnectionCredentials) | **Get** /api/v1/cloud_integrations/{cloud_provider}/credentials | Get connection credentials
[**GetService**](CloudintegrationAPI.md#GetService) | **Get** /api/v1/cloud_integrations/{cloud_provider}/services/{service_id} | Get service
[**ListAccounts**](CloudintegrationAPI.md#ListAccounts) | **Get** /api/v1/cloud_integrations/{cloud_provider}/accounts | List accounts
[**ListServicesMetadata**](CloudintegrationAPI.md#ListServicesMetadata) | **Get** /api/v1/cloud_integrations/{cloud_provider}/services | List services metadata
[**UpdateAccount**](CloudintegrationAPI.md#UpdateAccount) | **Put** /api/v1/cloud_integrations/{cloud_provider}/accounts/{id} | Update account
[**UpdateService**](CloudintegrationAPI.md#UpdateService) | **Put** /api/v1/cloud_integrations/{cloud_provider}/accounts/{id}/services/{service_id} | Update service



## AgentCheckIn

> AgentCheckInDeprecated200Response AgentCheckIn(ctx, cloudProvider).CloudintegrationtypesPostableAgentCheckIn(cloudintegrationtypesPostableAgentCheckIn).Execute()

Agent check-in



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
	cloudProvider := "cloudProvider_example" // string | 
	cloudintegrationtypesPostableAgentCheckIn := *openapiclient.NewCloudintegrationtypesPostableAgentCheckIn(map[string]interface{}{"key": interface{}(123)}) // CloudintegrationtypesPostableAgentCheckIn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.AgentCheckIn(context.Background(), cloudProvider).CloudintegrationtypesPostableAgentCheckIn(cloudintegrationtypesPostableAgentCheckIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.AgentCheckIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentCheckIn`: AgentCheckInDeprecated200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.AgentCheckIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudintegrationtypesPostableAgentCheckIn** | [**CloudintegrationtypesPostableAgentCheckIn**](CloudintegrationtypesPostableAgentCheckIn.md) |  | 

### Return type

[**AgentCheckInDeprecated200Response**](AgentCheckInDeprecated200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentCheckInDeprecated

> AgentCheckInDeprecated200Response AgentCheckInDeprecated(ctx, cloudProvider).CloudintegrationtypesPostableAgentCheckIn(cloudintegrationtypesPostableAgentCheckIn).Execute()

Agent check-in



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
	cloudProvider := "cloudProvider_example" // string | 
	cloudintegrationtypesPostableAgentCheckIn := *openapiclient.NewCloudintegrationtypesPostableAgentCheckIn(map[string]interface{}{"key": interface{}(123)}) // CloudintegrationtypesPostableAgentCheckIn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.AgentCheckInDeprecated(context.Background(), cloudProvider).CloudintegrationtypesPostableAgentCheckIn(cloudintegrationtypesPostableAgentCheckIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.AgentCheckInDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentCheckInDeprecated`: AgentCheckInDeprecated200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.AgentCheckInDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCheckInDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudintegrationtypesPostableAgentCheckIn** | [**CloudintegrationtypesPostableAgentCheckIn**](CloudintegrationtypesPostableAgentCheckIn.md) |  | 

### Return type

[**AgentCheckInDeprecated200Response**](AgentCheckInDeprecated200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> CreateAccount201Response CreateAccount(ctx, cloudProvider).CloudintegrationtypesPostableAccount(cloudintegrationtypesPostableAccount).Execute()

Create account



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
	cloudProvider := "cloudProvider_example" // string | 
	cloudintegrationtypesPostableAccount := *openapiclient.NewCloudintegrationtypesPostableAccount(*openapiclient.NewCloudintegrationtypesPostableAccountConfig(), *openapiclient.NewCloudintegrationtypesCredentials("IngestionKey_example", "IngestionUrl_example", "SigNozApiKey_example", "SigNozApiUrl_example")) // CloudintegrationtypesPostableAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.CreateAccount(context.Background(), cloudProvider).CloudintegrationtypesPostableAccount(cloudintegrationtypesPostableAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: CreateAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudintegrationtypesPostableAccount** | [**CloudintegrationtypesPostableAccount**](CloudintegrationtypesPostableAccount.md) |  | 

### Return type

[**CreateAccount201Response**](CreateAccount201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectAccount

> DisconnectAccount(ctx, cloudProvider, id).Execute()

Disconnect account



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
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudintegrationAPI.DisconnectAccount(context.Background(), cloudProvider, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.DisconnectAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectAccountRequest struct via the builder pattern


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


## GetAccount

> GetAccount200Response GetAccount(ctx, cloudProvider, id).Execute()

Get account



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
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.GetAccount(context.Background(), cloudProvider, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: GetAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAccount200Response**](GetAccount200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionCredentials

> GetConnectionCredentials200Response GetConnectionCredentials(ctx, cloudProvider).Execute()

Get connection credentials



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
	cloudProvider := "cloudProvider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.GetConnectionCredentials(context.Background(), cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.GetConnectionCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionCredentials`: GetConnectionCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.GetConnectionCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConnectionCredentials200Response**](GetConnectionCredentials200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> GetService200Response GetService(ctx, cloudProvider, serviceId).CloudIntegrationId(cloudIntegrationId).Execute()

Get service



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
	cloudProvider := "cloudProvider_example" // string | 
	serviceId := "serviceId_example" // string | 
	cloudIntegrationId := "cloudIntegrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.GetService(context.Background(), cloudProvider, serviceId).CloudIntegrationId(cloudIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.GetService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetService`: GetService200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudIntegrationId** | **string** |  | 

### Return type

[**GetService200Response**](GetService200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> ListAccounts200Response ListAccounts(ctx, cloudProvider).Execute()

List accounts



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
	cloudProvider := "cloudProvider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.ListAccounts(context.Background(), cloudProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: ListAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAccounts200Response**](ListAccounts200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServicesMetadata

> ListServicesMetadata200Response ListServicesMetadata(ctx, cloudProvider).CloudIntegrationId(cloudIntegrationId).Execute()

List services metadata



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
	cloudProvider := "cloudProvider_example" // string | 
	cloudIntegrationId := "cloudIntegrationId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudintegrationAPI.ListServicesMetadata(context.Background(), cloudProvider).CloudIntegrationId(cloudIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.ListServicesMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServicesMetadata`: ListServicesMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `CloudintegrationAPI.ListServicesMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudIntegrationId** | **string** |  | 

### Return type

[**ListServicesMetadata200Response**](ListServicesMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> UpdateAccount(ctx, cloudProvider, id).CloudintegrationtypesUpdatableAccount(cloudintegrationtypesUpdatableAccount).Execute()

Update account



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
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 
	cloudintegrationtypesUpdatableAccount := *openapiclient.NewCloudintegrationtypesUpdatableAccount(*openapiclient.NewCloudintegrationtypesUpdatableAccountConfig()) // CloudintegrationtypesUpdatableAccount |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudintegrationAPI.UpdateAccount(context.Background(), cloudProvider, id).CloudintegrationtypesUpdatableAccount(cloudintegrationtypesUpdatableAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudintegrationtypesUpdatableAccount** | [**CloudintegrationtypesUpdatableAccount**](CloudintegrationtypesUpdatableAccount.md) |  | 

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


## UpdateService

> UpdateService(ctx, cloudProvider, id, serviceId).CloudintegrationtypesUpdatableService(cloudintegrationtypesUpdatableService).Execute()

Update service



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
	cloudProvider := "cloudProvider_example" // string | 
	id := "id_example" // string | 
	serviceId := "serviceId_example" // string | 
	cloudintegrationtypesUpdatableService := *openapiclient.NewCloudintegrationtypesUpdatableService(*openapiclient.NewCloudintegrationtypesServiceConfig()) // CloudintegrationtypesUpdatableService |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudintegrationAPI.UpdateService(context.Background(), cloudProvider, id, serviceId).CloudintegrationtypesUpdatableService(cloudintegrationtypesUpdatableService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudintegrationAPI.UpdateService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** |  | 
**id** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cloudintegrationtypesUpdatableService** | [**CloudintegrationtypesUpdatableService**](CloudintegrationtypesUpdatableService.md) |  | 

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

