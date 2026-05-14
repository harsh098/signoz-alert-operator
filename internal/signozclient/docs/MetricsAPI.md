# \MetricsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricAlerts**](MetricsAPI.md#GetMetricAlerts) | **Get** /api/v2/metrics/{metric_name}/alerts | Get metric alerts
[**GetMetricAttributes**](MetricsAPI.md#GetMetricAttributes) | **Get** /api/v2/metrics/{metric_name}/attributes | Get metric attributes
[**GetMetricDashboards**](MetricsAPI.md#GetMetricDashboards) | **Get** /api/v2/metrics/{metric_name}/dashboards | Get metric dashboards
[**GetMetricHighlights**](MetricsAPI.md#GetMetricHighlights) | **Get** /api/v2/metrics/{metric_name}/highlights | Get metric highlights
[**GetMetricMetadata**](MetricsAPI.md#GetMetricMetadata) | **Get** /api/v2/metrics/{metric_name}/metadata | Get metric metadata
[**GetMetricsOnboardingStatus**](MetricsAPI.md#GetMetricsOnboardingStatus) | **Get** /api/v2/metrics/onboarding | Check if non-SigNoz metrics have been received
[**GetMetricsStats**](MetricsAPI.md#GetMetricsStats) | **Post** /api/v2/metrics/stats | Get metrics statistics
[**GetMetricsTreemap**](MetricsAPI.md#GetMetricsTreemap) | **Post** /api/v2/metrics/treemap | Get metrics treemap
[**InspectMetrics**](MetricsAPI.md#InspectMetrics) | **Post** /api/v2/metrics/inspect | Inspect raw metric data points
[**ListMetrics**](MetricsAPI.md#ListMetrics) | **Get** /api/v2/metrics | List metric names
[**UpdateMetricMetadata**](MetricsAPI.md#UpdateMetricMetadata) | **Post** /api/v2/metrics/{metric_name}/metadata | Update metric metadata



## GetMetricAlerts

> GetMetricAlerts200Response GetMetricAlerts(ctx, metricName).Execute()

Get metric alerts



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
	metricName := "metricName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricAlerts(context.Background(), metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricAlerts`: GetMetricAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricAlerts200Response**](GetMetricAlerts200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricAttributes

> GetMetricAttributes200Response GetMetricAttributes(ctx, metricName).Start(start).End(end).Execute()

Get metric attributes



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
	metricName := "metricName_example" // string | 
	start := int32(56) // int32 |  (optional)
	end := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricAttributes(context.Background(), metricName).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricAttributes`: GetMetricAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | 
 **end** | **int32** |  | 

### Return type

[**GetMetricAttributes200Response**](GetMetricAttributes200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricDashboards

> GetMetricDashboards200Response GetMetricDashboards(ctx, metricName).Execute()

Get metric dashboards



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
	metricName := "metricName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricDashboards(context.Background(), metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricDashboards`: GetMetricDashboards200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricDashboards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricDashboards200Response**](GetMetricDashboards200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricHighlights

> GetMetricHighlights200Response GetMetricHighlights(ctx, metricName).Execute()

Get metric highlights



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
	metricName := "metricName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricHighlights(context.Background(), metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricHighlights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricHighlights`: GetMetricHighlights200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricHighlights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricHighlightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricHighlights200Response**](GetMetricHighlights200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricMetadata

> GetMetricMetadata200Response GetMetricMetadata(ctx, metricName).Execute()

Get metric metadata



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
	metricName := "metricName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricMetadata(context.Background(), metricName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricMetadata`: GetMetricMetadata200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMetricMetadata200Response**](GetMetricMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsOnboardingStatus

> GetMetricsOnboardingStatus200Response GetMetricsOnboardingStatus(ctx).Execute()

Check if non-SigNoz metrics have been received



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
	resp, r, err := apiClient.MetricsAPI.GetMetricsOnboardingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricsOnboardingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsOnboardingStatus`: GetMetricsOnboardingStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricsOnboardingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsOnboardingStatusRequest struct via the builder pattern


### Return type

[**GetMetricsOnboardingStatus200Response**](GetMetricsOnboardingStatus200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsStats

> GetMetricsStats200Response GetMetricsStats(ctx).MetricsexplorertypesStatsRequest(metricsexplorertypesStatsRequest).Execute()

Get metrics statistics



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
	metricsexplorertypesStatsRequest := *openapiclient.NewMetricsexplorertypesStatsRequest(int64(123), int32(123), int64(123)) // MetricsexplorertypesStatsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricsStats(context.Background()).MetricsexplorertypesStatsRequest(metricsexplorertypesStatsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsStats`: GetMetricsStats200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsexplorertypesStatsRequest** | [**MetricsexplorertypesStatsRequest**](MetricsexplorertypesStatsRequest.md) |  | 

### Return type

[**GetMetricsStats200Response**](GetMetricsStats200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsTreemap

> GetMetricsTreemap200Response GetMetricsTreemap(ctx).MetricsexplorertypesTreemapRequest(metricsexplorertypesTreemapRequest).Execute()

Get metrics treemap



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
	metricsexplorertypesTreemapRequest := *openapiclient.NewMetricsexplorertypesTreemapRequest(int64(123), int32(123), openapiclient.MetricsexplorertypesTreemapMode("timeseries"), int64(123)) // MetricsexplorertypesTreemapRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetMetricsTreemap(context.Background()).MetricsexplorertypesTreemapRequest(metricsexplorertypesTreemapRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetMetricsTreemap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricsTreemap`: GetMetricsTreemap200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetMetricsTreemap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsTreemapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsexplorertypesTreemapRequest** | [**MetricsexplorertypesTreemapRequest**](MetricsexplorertypesTreemapRequest.md) |  | 

### Return type

[**GetMetricsTreemap200Response**](GetMetricsTreemap200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InspectMetrics

> InspectMetrics200Response InspectMetrics(ctx).MetricsexplorertypesInspectMetricsRequest(metricsexplorertypesInspectMetricsRequest).Execute()

Inspect raw metric data points



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
	metricsexplorertypesInspectMetricsRequest := *openapiclient.NewMetricsexplorertypesInspectMetricsRequest(int64(123), "MetricName_example", int64(123)) // MetricsexplorertypesInspectMetricsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.InspectMetrics(context.Background()).MetricsexplorertypesInspectMetricsRequest(metricsexplorertypesInspectMetricsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.InspectMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InspectMetrics`: InspectMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.InspectMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInspectMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsexplorertypesInspectMetricsRequest** | [**MetricsexplorertypesInspectMetricsRequest**](MetricsexplorertypesInspectMetricsRequest.md) |  | 

### Return type

[**InspectMetrics200Response**](InspectMetrics200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetrics

> ListMetrics200Response ListMetrics(ctx).Start(start).End(end).Limit(limit).SearchText(searchText).Source(source).Execute()

List metric names



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
	start := int32(56) // int32 |  (optional)
	end := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	searchText := "searchText_example" // string |  (optional)
	source := "source_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.ListMetrics(context.Background()).Start(start).End(end).Limit(limit).SearchText(searchText).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.ListMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetrics`: ListMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.ListMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** |  | 
 **end** | **int32** |  | 
 **limit** | **int32** |  | 
 **searchText** | **string** |  | 
 **source** | **string** |  | 

### Return type

[**ListMetrics200Response**](ListMetrics200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricMetadata

> string UpdateMetricMetadata(ctx, metricName).MetricsexplorertypesUpdateMetricMetadataRequest(metricsexplorertypesUpdateMetricMetadataRequest).Execute()

Update metric metadata



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
	metricName := "metricName_example" // string | 
	metricsexplorertypesUpdateMetricMetadataRequest := *openapiclient.NewMetricsexplorertypesUpdateMetricMetadataRequest("Description_example", false, "MetricName_example", openapiclient.MetrictypesTemporality("delta"), openapiclient.MetrictypesType("gauge"), "Unit_example") // MetricsexplorertypesUpdateMetricMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.UpdateMetricMetadata(context.Background(), metricName).MetricsexplorertypesUpdateMetricMetadataRequest(metricsexplorertypesUpdateMetricMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.UpdateMetricMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMetricMetadata`: string
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.UpdateMetricMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsexplorertypesUpdateMetricMetadataRequest** | [**MetricsexplorertypesUpdateMetricMetadataRequest**](MetricsexplorertypesUpdateMetricMetadataRequest.md) |  | 

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

