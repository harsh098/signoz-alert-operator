# \RulesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](RulesAPI.md#CreateRule) | **Post** /api/v2/rules | Create alert rule
[**DeleteRuleByID**](RulesAPI.md#DeleteRuleByID) | **Delete** /api/v2/rules/{id} | Delete alert rule
[**GetRuleByID**](RulesAPI.md#GetRuleByID) | **Get** /api/v2/rules/{id} | Get alert rule by ID
[**GetRuleHistoryFilterKeys**](RulesAPI.md#GetRuleHistoryFilterKeys) | **Get** /api/v2/rules/{id}/history/filter_keys | Get rule history filter keys
[**GetRuleHistoryFilterValues**](RulesAPI.md#GetRuleHistoryFilterValues) | **Get** /api/v2/rules/{id}/history/filter_values | Get rule history filter values
[**GetRuleHistoryOverallStatus**](RulesAPI.md#GetRuleHistoryOverallStatus) | **Get** /api/v2/rules/{id}/history/overall_status | Get rule overall status timeline
[**GetRuleHistoryStats**](RulesAPI.md#GetRuleHistoryStats) | **Get** /api/v2/rules/{id}/history/stats | Get rule history stats
[**GetRuleHistoryTimeline**](RulesAPI.md#GetRuleHistoryTimeline) | **Get** /api/v2/rules/{id}/history/timeline | Get rule history timeline
[**GetRuleHistoryTopContributors**](RulesAPI.md#GetRuleHistoryTopContributors) | **Get** /api/v2/rules/{id}/history/top_contributors | Get top contributors to rule firing
[**ListRules**](RulesAPI.md#ListRules) | **Get** /api/v2/rules | List alert rules
[**PatchRuleByID**](RulesAPI.md#PatchRuleByID) | **Patch** /api/v2/rules/{id} | Patch alert rule
[**TestRule**](RulesAPI.md#TestRule) | **Post** /api/v2/rules/test | Test alert rule
[**UpdateRuleByID**](RulesAPI.md#UpdateRuleByID) | **Put** /api/v2/rules/{id} | Update alert rule



## CreateRule

> CreateRule201Response CreateRule(ctx).RuletypesPostableRule(ruletypesPostableRule).Execute()

Create alert rule



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
	ruletypesPostableRule := *openapiclient.NewRuletypesPostableRule("Alert_example", openapiclient.RuletypesAlertType("METRIC_BASED_ALERT"), *openapiclient.NewRuletypesRuleCondition(*openapiclient.NewRuletypesAlertCompositeQuery(openapiclient.RuletypesPanelType("value"), []openapiclient.Querybuildertypesv5QueryEnvelope{openapiclient.Querybuildertypesv5QueryEnvelope{Querybuildertypesv5QueryEnvelopeBuilderLog: openapiclient.NewQuerybuildertypesv5QueryEnvelopeBuilderLog()}}, openapiclient.RuletypesQueryType("builder"))), openapiclient.RuletypesRuleType("threshold_rule")) // RuletypesPostableRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRule(context.Background()).RuletypesPostableRule(ruletypesPostableRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule`: CreateRule201Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruletypesPostableRule** | [**RuletypesPostableRule**](RuletypesPostableRule.md) |  | 

### Return type

[**CreateRule201Response**](CreateRule201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleByID

> DeleteRuleByID(ctx, id).Execute()

Delete alert rule



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
	r, err := apiClient.RulesAPI.DeleteRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRuleByIDRequest struct via the builder pattern


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


## GetRuleByID

> CreateRule201Response GetRuleByID(ctx, id).Execute()

Get alert rule by ID



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
	resp, r, err := apiClient.RulesAPI.GetRuleByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleByID`: CreateRule201Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateRule201Response**](CreateRule201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryFilterKeys

> GetFieldsKeys200Response GetRuleHistoryFilterKeys(ctx, id).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()

Get rule history filter keys



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
	signal := openapiclient.TelemetrytypesSignal("traces") // TelemetrytypesSignal |  (optional)
	source := openapiclient.TelemetrytypesSource("meter") // TelemetrytypesSource |  (optional)
	limit := int32(56) // int32 |  (optional)
	startUnixMilli := int64(789) // int64 |  (optional)
	endUnixMilli := int64(789) // int64 |  (optional)
	fieldContext := openapiclient.TelemetrytypesFieldContext("metric") // TelemetrytypesFieldContext |  (optional)
	fieldDataType := openapiclient.TelemetrytypesFieldDataType("string") // TelemetrytypesFieldDataType |  (optional)
	metricName := "metricName_example" // string |  (optional)
	metricNamespace := "metricNamespace_example" // string |  (optional)
	searchText := "searchText_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryFilterKeys(context.Background(), id).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryFilterKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryFilterKeys`: GetFieldsKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryFilterKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryFilterKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**TelemetrytypesSignal**](TelemetrytypesSignal.md) |  | 
 **source** | [**TelemetrytypesSource**](TelemetrytypesSource.md) |  | 
 **limit** | **int32** |  | 
 **startUnixMilli** | **int64** |  | 
 **endUnixMilli** | **int64** |  | 
 **fieldContext** | [**TelemetrytypesFieldContext**](TelemetrytypesFieldContext.md) |  | 
 **fieldDataType** | [**TelemetrytypesFieldDataType**](TelemetrytypesFieldDataType.md) |  | 
 **metricName** | **string** |  | 
 **metricNamespace** | **string** |  | 
 **searchText** | **string** |  | 

### Return type

[**GetFieldsKeys200Response**](GetFieldsKeys200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryFilterValues

> GetFieldsValues200Response GetRuleHistoryFilterValues(ctx, id).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()

Get rule history filter values



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
	signal := openapiclient.TelemetrytypesSignal("traces") // TelemetrytypesSignal |  (optional)
	source := openapiclient.TelemetrytypesSource("meter") // TelemetrytypesSource |  (optional)
	limit := int32(56) // int32 |  (optional)
	startUnixMilli := int64(789) // int64 |  (optional)
	endUnixMilli := int64(789) // int64 |  (optional)
	fieldContext := openapiclient.TelemetrytypesFieldContext("metric") // TelemetrytypesFieldContext |  (optional)
	fieldDataType := openapiclient.TelemetrytypesFieldDataType("string") // TelemetrytypesFieldDataType |  (optional)
	metricName := "metricName_example" // string |  (optional)
	metricNamespace := "metricNamespace_example" // string |  (optional)
	searchText := "searchText_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	existingQuery := "existingQuery_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryFilterValues(context.Background(), id).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryFilterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryFilterValues`: GetFieldsValues200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryFilterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryFilterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**TelemetrytypesSignal**](TelemetrytypesSignal.md) |  | 
 **source** | [**TelemetrytypesSource**](TelemetrytypesSource.md) |  | 
 **limit** | **int32** |  | 
 **startUnixMilli** | **int64** |  | 
 **endUnixMilli** | **int64** |  | 
 **fieldContext** | [**TelemetrytypesFieldContext**](TelemetrytypesFieldContext.md) |  | 
 **fieldDataType** | [**TelemetrytypesFieldDataType**](TelemetrytypesFieldDataType.md) |  | 
 **metricName** | **string** |  | 
 **metricNamespace** | **string** |  | 
 **searchText** | **string** |  | 
 **name** | **string** |  | 
 **existingQuery** | **string** |  | 

### Return type

[**GetFieldsValues200Response**](GetFieldsValues200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryOverallStatus

> GetRuleHistoryOverallStatus200Response GetRuleHistoryOverallStatus(ctx, id).Start(start).End(end).Execute()

Get rule overall status timeline



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
	start := int64(789) // int64 | 
	end := int64(789) // int64 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryOverallStatus(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryOverallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryOverallStatus`: GetRuleHistoryOverallStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryOverallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryOverallStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int64** |  | 
 **end** | **int64** |  | 


### Return type

[**GetRuleHistoryOverallStatus200Response**](GetRuleHistoryOverallStatus200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryStats

> GetRuleHistoryStats200Response GetRuleHistoryStats(ctx, id).Start(start).End(end).Execute()

Get rule history stats



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
	start := int64(789) // int64 | 
	end := int64(789) // int64 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryStats(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryStats`: GetRuleHistoryStats200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int64** |  | 
 **end** | **int64** |  | 


### Return type

[**GetRuleHistoryStats200Response**](GetRuleHistoryStats200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryTimeline

> GetRuleHistoryTimeline200Response GetRuleHistoryTimeline(ctx, id).Start(start).End(end).State(state).FilterExpression(filterExpression).Limit(limit).Order(order).Cursor(cursor).Execute()

Get rule history timeline



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
	start := int64(789) // int64 | 
	end := int64(789) // int64 | 
	id := "id_example" // string | 
	state := openapiclient.RuletypesAlertState("inactive") // RuletypesAlertState |  (optional)
	filterExpression := "filterExpression_example" // string |  (optional)
	limit := int64(789) // int64 |  (optional)
	order := openapiclient.Querybuildertypesv5OrderDirection("asc") // Querybuildertypesv5OrderDirection |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryTimeline(context.Background(), id).Start(start).End(end).State(state).FilterExpression(filterExpression).Limit(limit).Order(order).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryTimeline`: GetRuleHistoryTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int64** |  | 
 **end** | **int64** |  | 

 **state** | [**RuletypesAlertState**](RuletypesAlertState.md) |  | 
 **filterExpression** | **string** |  | 
 **limit** | **int64** |  | 
 **order** | [**Querybuildertypesv5OrderDirection**](Querybuildertypesv5OrderDirection.md) |  | 
 **cursor** | **string** |  | 

### Return type

[**GetRuleHistoryTimeline200Response**](GetRuleHistoryTimeline200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHistoryTopContributors

> GetRuleHistoryTopContributors200Response GetRuleHistoryTopContributors(ctx, id).Start(start).End(end).Execute()

Get top contributors to rule firing



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
	start := int64(789) // int64 | 
	end := int64(789) // int64 | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHistoryTopContributors(context.Background(), id).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHistoryTopContributors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHistoryTopContributors`: GetRuleHistoryTopContributors200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHistoryTopContributors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHistoryTopContributorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int64** |  | 
 **end** | **int64** |  | 


### Return type

[**GetRuleHistoryTopContributors200Response**](GetRuleHistoryTopContributors200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> ListRules200Response ListRules(ctx).Execute()

List alert rules



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
	resp, r, err := apiClient.RulesAPI.ListRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRules`: ListRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ListRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


### Return type

[**ListRules200Response**](ListRules200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRuleByID

> CreateRule201Response PatchRuleByID(ctx, id).RuletypesPostableRule(ruletypesPostableRule).Execute()

Patch alert rule



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
	ruletypesPostableRule := *openapiclient.NewRuletypesPostableRule("Alert_example", openapiclient.RuletypesAlertType("METRIC_BASED_ALERT"), *openapiclient.NewRuletypesRuleCondition(*openapiclient.NewRuletypesAlertCompositeQuery(openapiclient.RuletypesPanelType("value"), []openapiclient.Querybuildertypesv5QueryEnvelope{openapiclient.Querybuildertypesv5QueryEnvelope{Querybuildertypesv5QueryEnvelopeBuilderLog: openapiclient.NewQuerybuildertypesv5QueryEnvelopeBuilderLog()}}, openapiclient.RuletypesQueryType("builder"))), openapiclient.RuletypesRuleType("threshold_rule")) // RuletypesPostableRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.PatchRuleByID(context.Background(), id).RuletypesPostableRule(ruletypesPostableRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.PatchRuleByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRuleByID`: CreateRule201Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.PatchRuleByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruletypesPostableRule** | [**RuletypesPostableRule**](RuletypesPostableRule.md) |  | 

### Return type

[**CreateRule201Response**](CreateRule201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRule

> TestRule200Response TestRule(ctx).RuletypesPostableRule(ruletypesPostableRule).Execute()

Test alert rule



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
	ruletypesPostableRule := *openapiclient.NewRuletypesPostableRule("Alert_example", openapiclient.RuletypesAlertType("METRIC_BASED_ALERT"), *openapiclient.NewRuletypesRuleCondition(*openapiclient.NewRuletypesAlertCompositeQuery(openapiclient.RuletypesPanelType("value"), []openapiclient.Querybuildertypesv5QueryEnvelope{openapiclient.Querybuildertypesv5QueryEnvelope{Querybuildertypesv5QueryEnvelopeBuilderLog: openapiclient.NewQuerybuildertypesv5QueryEnvelopeBuilderLog()}}, openapiclient.RuletypesQueryType("builder"))), openapiclient.RuletypesRuleType("threshold_rule")) // RuletypesPostableRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.TestRule(context.Background()).RuletypesPostableRule(ruletypesPostableRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.TestRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRule`: TestRule200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.TestRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruletypesPostableRule** | [**RuletypesPostableRule**](RuletypesPostableRule.md) |  | 

### Return type

[**TestRule200Response**](TestRule200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleByID

> UpdateRuleByID(ctx, id).RuletypesPostableRule(ruletypesPostableRule).Execute()

Update alert rule



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
	ruletypesPostableRule := *openapiclient.NewRuletypesPostableRule("Alert_example", openapiclient.RuletypesAlertType("METRIC_BASED_ALERT"), *openapiclient.NewRuletypesRuleCondition(*openapiclient.NewRuletypesAlertCompositeQuery(openapiclient.RuletypesPanelType("value"), []openapiclient.Querybuildertypesv5QueryEnvelope{openapiclient.Querybuildertypesv5QueryEnvelope{Querybuildertypesv5QueryEnvelopeBuilderLog: openapiclient.NewQuerybuildertypesv5QueryEnvelopeBuilderLog()}}, openapiclient.RuletypesQueryType("builder"))), openapiclient.RuletypesRuleType("threshold_rule")) // RuletypesPostableRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.UpdateRuleByID(context.Background(), id).RuletypesPostableRule(ruletypesPostableRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateRuleByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruletypesPostableRule** | [**RuletypesPostableRule**](RuletypesPostableRule.md) |  | 

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

