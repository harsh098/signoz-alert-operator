# \DowntimeschedulesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDowntimeSchedule**](DowntimeschedulesAPI.md#CreateDowntimeSchedule) | **Post** /api/v1/downtime_schedules | Create downtime schedule
[**DeleteDowntimeScheduleByID**](DowntimeschedulesAPI.md#DeleteDowntimeScheduleByID) | **Delete** /api/v1/downtime_schedules/{id} | Delete downtime schedule
[**GetDowntimeScheduleByID**](DowntimeschedulesAPI.md#GetDowntimeScheduleByID) | **Get** /api/v1/downtime_schedules/{id} | Get downtime schedule by ID
[**ListDowntimeSchedules**](DowntimeschedulesAPI.md#ListDowntimeSchedules) | **Get** /api/v1/downtime_schedules | List downtime schedules
[**UpdateDowntimeScheduleByID**](DowntimeschedulesAPI.md#UpdateDowntimeScheduleByID) | **Put** /api/v1/downtime_schedules/{id} | Update downtime schedule



## CreateDowntimeSchedule

> CreateDowntimeSchedule201Response CreateDowntimeSchedule(ctx).RuletypesPostablePlannedMaintenance(ruletypesPostablePlannedMaintenance).Execute()

Create downtime schedule



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
	ruletypesPostablePlannedMaintenance := *openapiclient.NewRuletypesPostablePlannedMaintenance("Name_example", *openapiclient.NewRuletypesSchedule("Timezone_example")) // RuletypesPostablePlannedMaintenance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DowntimeschedulesAPI.CreateDowntimeSchedule(context.Background()).RuletypesPostablePlannedMaintenance(ruletypesPostablePlannedMaintenance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimeschedulesAPI.CreateDowntimeSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDowntimeSchedule`: CreateDowntimeSchedule201Response
	fmt.Fprintf(os.Stdout, "Response from `DowntimeschedulesAPI.CreateDowntimeSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDowntimeScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruletypesPostablePlannedMaintenance** | [**RuletypesPostablePlannedMaintenance**](RuletypesPostablePlannedMaintenance.md) |  | 

### Return type

[**CreateDowntimeSchedule201Response**](CreateDowntimeSchedule201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDowntimeScheduleByID

> DeleteDowntimeScheduleByID(ctx, id).Execute()

Delete downtime schedule



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
	r, err := apiClient.DowntimeschedulesAPI.DeleteDowntimeScheduleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimeschedulesAPI.DeleteDowntimeScheduleByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDowntimeScheduleByIDRequest struct via the builder pattern


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


## GetDowntimeScheduleByID

> CreateDowntimeSchedule201Response GetDowntimeScheduleByID(ctx, id).Execute()

Get downtime schedule by ID



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
	resp, r, err := apiClient.DowntimeschedulesAPI.GetDowntimeScheduleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimeschedulesAPI.GetDowntimeScheduleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDowntimeScheduleByID`: CreateDowntimeSchedule201Response
	fmt.Fprintf(os.Stdout, "Response from `DowntimeschedulesAPI.GetDowntimeScheduleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDowntimeScheduleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDowntimeSchedule201Response**](CreateDowntimeSchedule201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDowntimeSchedules

> ListDowntimeSchedules200Response ListDowntimeSchedules(ctx).Active(active).Recurring(recurring).Execute()

List downtime schedules



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
	active := true // bool |  (optional)
	recurring := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DowntimeschedulesAPI.ListDowntimeSchedules(context.Background()).Active(active).Recurring(recurring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimeschedulesAPI.ListDowntimeSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDowntimeSchedules`: ListDowntimeSchedules200Response
	fmt.Fprintf(os.Stdout, "Response from `DowntimeschedulesAPI.ListDowntimeSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDowntimeSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** |  | 
 **recurring** | **bool** |  | 

### Return type

[**ListDowntimeSchedules200Response**](ListDowntimeSchedules200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDowntimeScheduleByID

> UpdateDowntimeScheduleByID(ctx, id).RuletypesPostablePlannedMaintenance(ruletypesPostablePlannedMaintenance).Execute()

Update downtime schedule



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
	ruletypesPostablePlannedMaintenance := *openapiclient.NewRuletypesPostablePlannedMaintenance("Name_example", *openapiclient.NewRuletypesSchedule("Timezone_example")) // RuletypesPostablePlannedMaintenance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DowntimeschedulesAPI.UpdateDowntimeScheduleByID(context.Background(), id).RuletypesPostablePlannedMaintenance(ruletypesPostablePlannedMaintenance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DowntimeschedulesAPI.UpdateDowntimeScheduleByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateDowntimeScheduleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruletypesPostablePlannedMaintenance** | [**RuletypesPostablePlannedMaintenance**](RuletypesPostablePlannedMaintenance.md) |  | 

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

