# \FieldsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFieldsKeys**](FieldsAPI.md#GetFieldsKeys) | **Get** /api/v1/fields/keys | Get field keys
[**GetFieldsValues**](FieldsAPI.md#GetFieldsValues) | **Get** /api/v1/fields/values | Get field values



## GetFieldsKeys

> GetFieldsKeys200Response GetFieldsKeys(ctx).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()

Get field keys



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
	resp, r, err := apiClient.FieldsAPI.GetFieldsKeys(context.Background()).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.GetFieldsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldsKeys`: GetFieldsKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.GetFieldsKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsKeysRequest struct via the builder pattern


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


## GetFieldsValues

> GetFieldsValues200Response GetFieldsValues(ctx).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()

Get field values



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
	resp, r, err := apiClient.FieldsAPI.GetFieldsValues(context.Background()).Signal(signal).Source(source).Limit(limit).StartUnixMilli(startUnixMilli).EndUnixMilli(endUnixMilli).FieldContext(fieldContext).FieldDataType(fieldDataType).MetricName(metricName).MetricNamespace(metricNamespace).SearchText(searchText).Name(name).ExistingQuery(existingQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.GetFieldsValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldsValues`: GetFieldsValues200Response
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.GetFieldsValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsValuesRequest struct via the builder pattern


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

