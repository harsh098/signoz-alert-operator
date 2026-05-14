# \SpanmapperAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpanMapper**](SpanmapperAPI.md#CreateSpanMapper) | **Post** /api/v1/span_mapper_groups/{groupId}/span_mappers | Create a span mapper
[**CreateSpanMapperGroup**](SpanmapperAPI.md#CreateSpanMapperGroup) | **Post** /api/v1/span_mapper_groups | Create a span attribute mapping group
[**DeleteSpanMapper**](SpanmapperAPI.md#DeleteSpanMapper) | **Delete** /api/v1/span_mapper_groups/{groupId}/span_mappers/{mapperId} | Delete a span mapper
[**DeleteSpanMapperGroup**](SpanmapperAPI.md#DeleteSpanMapperGroup) | **Delete** /api/v1/span_mapper_groups/{groupId} | Delete a span attribute mapping group
[**ListSpanMapperGroups**](SpanmapperAPI.md#ListSpanMapperGroups) | **Get** /api/v1/span_mapper_groups | List span attribute mapping groups
[**ListSpanMappers**](SpanmapperAPI.md#ListSpanMappers) | **Get** /api/v1/span_mapper_groups/{groupId}/span_mappers | List span mappers for a group
[**UpdateSpanMapper**](SpanmapperAPI.md#UpdateSpanMapper) | **Patch** /api/v1/span_mapper_groups/{groupId}/span_mappers/{mapperId} | Update a span mapper
[**UpdateSpanMapperGroup**](SpanmapperAPI.md#UpdateSpanMapperGroup) | **Patch** /api/v1/span_mapper_groups/{groupId} | Update a span attribute mapping group



## CreateSpanMapper

> CreateSpanMapper201Response CreateSpanMapper(ctx, groupId).SpantypesPostableSpanMapper(spantypesPostableSpanMapper).Execute()

Create a span mapper



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
	groupId := "groupId_example" // string | 
	spantypesPostableSpanMapper := *openapiclient.NewSpantypesPostableSpanMapper(*openapiclient.NewSpantypesSpanMapperConfig([]openapiclient.SpantypesSpanMapperSource{*openapiclient.NewSpantypesSpanMapperSource(openapiclient.SpantypesFieldContext("attribute"), "Key_example", openapiclient.SpantypesSpanMapperOperation("move"), int32(123))}), openapiclient.SpantypesFieldContext("attribute"), "Name_example") // SpantypesPostableSpanMapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanmapperAPI.CreateSpanMapper(context.Background(), groupId).SpantypesPostableSpanMapper(spantypesPostableSpanMapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.CreateSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpanMapper`: CreateSpanMapper201Response
	fmt.Fprintf(os.Stdout, "Response from `SpanmapperAPI.CreateSpanMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spantypesPostableSpanMapper** | [**SpantypesPostableSpanMapper**](SpantypesPostableSpanMapper.md) |  | 

### Return type

[**CreateSpanMapper201Response**](CreateSpanMapper201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpanMapperGroup

> CreateSpanMapperGroup201Response CreateSpanMapperGroup(ctx).SpantypesPostableSpanMapperGroup(spantypesPostableSpanMapperGroup).Execute()

Create a span attribute mapping group



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
	spantypesPostableSpanMapperGroup := *openapiclient.NewSpantypesPostableSpanMapperGroup(map[string]interface{}(123), *openapiclient.NewSpantypesSpanMapperGroupCondition([]string{"Attributes_example"}, []string{"Resource_example"}), "Name_example") // SpantypesPostableSpanMapperGroup |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanmapperAPI.CreateSpanMapperGroup(context.Background()).SpantypesPostableSpanMapperGroup(spantypesPostableSpanMapperGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.CreateSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpanMapperGroup`: CreateSpanMapperGroup201Response
	fmt.Fprintf(os.Stdout, "Response from `SpanmapperAPI.CreateSpanMapperGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanMapperGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spantypesPostableSpanMapperGroup** | [**SpantypesPostableSpanMapperGroup**](SpantypesPostableSpanMapperGroup.md) |  | 

### Return type

[**CreateSpanMapperGroup201Response**](CreateSpanMapperGroup201Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSpanMapper

> DeleteSpanMapper(ctx, groupId, mapperId).Execute()

Delete a span mapper



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
	groupId := "groupId_example" // string | 
	mapperId := "mapperId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpanmapperAPI.DeleteSpanMapper(context.Background(), groupId, mapperId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.DeleteSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**mapperId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpanMapperRequest struct via the builder pattern


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


## DeleteSpanMapperGroup

> DeleteSpanMapperGroup(ctx, groupId).Execute()

Delete a span attribute mapping group



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpanmapperAPI.DeleteSpanMapperGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.DeleteSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpanMapperGroupRequest struct via the builder pattern


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


## ListSpanMapperGroups

> ListSpanMapperGroups200Response ListSpanMapperGroups(ctx).Category(category).Enabled(enabled).Execute()

List span attribute mapping groups



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
	category := map[string]interface{}{"key": map[string]interface{}(123)} // map[string]interface{} |  (optional)
	enabled := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanmapperAPI.ListSpanMapperGroups(context.Background()).Category(category).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.ListSpanMapperGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpanMapperGroups`: ListSpanMapperGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `SpanmapperAPI.ListSpanMapperGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpanMapperGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **enabled** | **bool** |  | 

### Return type

[**ListSpanMapperGroups200Response**](ListSpanMapperGroups200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpanMappers

> ListSpanMapperGroups200Response ListSpanMappers(ctx, groupId).Execute()

List span mappers for a group



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpanmapperAPI.ListSpanMappers(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.ListSpanMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpanMappers`: ListSpanMapperGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `SpanmapperAPI.ListSpanMappers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpanMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSpanMapperGroups200Response**](ListSpanMapperGroups200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpanMapper

> UpdateSpanMapper(ctx, groupId, mapperId).SpantypesUpdatableSpanMapper(spantypesUpdatableSpanMapper).Execute()

Update a span mapper



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
	groupId := "groupId_example" // string | 
	mapperId := "mapperId_example" // string | 
	spantypesUpdatableSpanMapper := *openapiclient.NewSpantypesUpdatableSpanMapper() // SpantypesUpdatableSpanMapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpanmapperAPI.UpdateSpanMapper(context.Background(), groupId, mapperId).SpantypesUpdatableSpanMapper(spantypesUpdatableSpanMapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.UpdateSpanMapper``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**mapperId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpanMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spantypesUpdatableSpanMapper** | [**SpantypesUpdatableSpanMapper**](SpantypesUpdatableSpanMapper.md) |  | 

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


## UpdateSpanMapperGroup

> UpdateSpanMapperGroup(ctx, groupId).SpantypesUpdatableSpanMapperGroup(spantypesUpdatableSpanMapperGroup).Execute()

Update a span attribute mapping group



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
	groupId := "groupId_example" // string | 
	spantypesUpdatableSpanMapperGroup := *openapiclient.NewSpantypesUpdatableSpanMapperGroup() // SpantypesUpdatableSpanMapperGroup |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpanmapperAPI.UpdateSpanMapperGroup(context.Background(), groupId).SpantypesUpdatableSpanMapperGroup(spantypesUpdatableSpanMapperGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpanmapperAPI.UpdateSpanMapperGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpanMapperGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spantypesUpdatableSpanMapperGroup** | [**SpantypesUpdatableSpanMapperGroup**](SpantypesUpdatableSpanMapperGroup.md) |  | 

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

