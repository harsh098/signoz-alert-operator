# \DashboardAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicDashboard**](DashboardAPI.md#CreatePublicDashboard) | **Post** /api/v1/dashboards/{id}/public | Create public dashboard
[**DeletePublicDashboard**](DashboardAPI.md#DeletePublicDashboard) | **Delete** /api/v1/dashboards/{id}/public | Delete public dashboard
[**GetPublicDashboard**](DashboardAPI.md#GetPublicDashboard) | **Get** /api/v1/dashboards/{id}/public | Get public dashboard
[**GetPublicDashboardData**](DashboardAPI.md#GetPublicDashboardData) | **Get** /api/v1/public/dashboards/{id} | Get public dashboard data
[**GetPublicDashboardWidgetQueryRange**](DashboardAPI.md#GetPublicDashboardWidgetQueryRange) | **Get** /api/v1/public/dashboards/{id}/widgets/{idx}/query_range | Get query range result
[**UpdatePublicDashboard**](DashboardAPI.md#UpdatePublicDashboard) | **Put** /api/v1/dashboards/{id}/public | Update public dashboard



## CreatePublicDashboard

> CreatePublicDashboard201Response CreatePublicDashboard(ctx, id).DashboardtypesPostablePublicDashboard(dashboardtypesPostablePublicDashboard).Execute()

Create public dashboard



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
	dashboardtypesPostablePublicDashboard := *openapiclient.NewDashboardtypesPostablePublicDashboard() // DashboardtypesPostablePublicDashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.CreatePublicDashboard(context.Background(), id).DashboardtypesPostablePublicDashboard(dashboardtypesPostablePublicDashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.CreatePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicDashboard`: CreatePublicDashboard201Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.CreatePublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardtypesPostablePublicDashboard** | [**DashboardtypesPostablePublicDashboard**](DashboardtypesPostablePublicDashboard.md) |  | 

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


## DeletePublicDashboard

> string DeletePublicDashboard(ctx, id).Execute()

Delete public dashboard



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
	resp, r, err := apiClient.DashboardAPI.DeletePublicDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.DeletePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePublicDashboard`: string
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.DeletePublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicDashboardRequest struct via the builder pattern


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


## GetPublicDashboard

> GetPublicDashboard200Response GetPublicDashboard(ctx, id).Execute()

Get public dashboard



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
	resp, r, err := apiClient.DashboardAPI.GetPublicDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboard`: GetPublicDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPublicDashboard200Response**](GetPublicDashboard200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDashboardData

> GetPublicDashboardData200Response GetPublicDashboardData(ctx, id).Execute()

Get public dashboard data



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
	resp, r, err := apiClient.DashboardAPI.GetPublicDashboardData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPublicDashboardData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboardData`: GetPublicDashboardData200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPublicDashboardData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPublicDashboardData200Response**](GetPublicDashboardData200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDashboardWidgetQueryRange

> GetPublicDashboardWidgetQueryRange200Response GetPublicDashboardWidgetQueryRange(ctx, id, idx).Execute()

Get query range result



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
	idx := "idx_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetPublicDashboardWidgetQueryRange(context.Background(), id, idx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPublicDashboardWidgetQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDashboardWidgetQueryRange`: GetPublicDashboardWidgetQueryRange200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPublicDashboardWidgetQueryRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**idx** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDashboardWidgetQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPublicDashboardWidgetQueryRange200Response**](GetPublicDashboardWidgetQueryRange200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicDashboard

> string UpdatePublicDashboard(ctx, id).DashboardtypesUpdatablePublicDashboard(dashboardtypesUpdatablePublicDashboard).Execute()

Update public dashboard



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
	dashboardtypesUpdatablePublicDashboard := *openapiclient.NewDashboardtypesUpdatablePublicDashboard() // DashboardtypesUpdatablePublicDashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.UpdatePublicDashboard(context.Background(), id).DashboardtypesUpdatablePublicDashboard(dashboardtypesUpdatablePublicDashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.UpdatePublicDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePublicDashboard`: string
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.UpdatePublicDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardtypesUpdatablePublicDashboard** | [**DashboardtypesUpdatablePublicDashboard**](DashboardtypesUpdatablePublicDashboard.md) |  | 

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

