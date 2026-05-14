# \LlmpricingrulesAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateLLMPricingRules**](LlmpricingrulesAPI.md#CreateOrUpdateLLMPricingRules) | **Put** /api/v1/llm_pricing_rules | Create or update pricing rules
[**DeleteLLMPricingRule**](LlmpricingrulesAPI.md#DeleteLLMPricingRule) | **Delete** /api/v1/llm_pricing_rules/{id} | Delete a pricing rule
[**GetLLMPricingRule**](LlmpricingrulesAPI.md#GetLLMPricingRule) | **Get** /api/v1/llm_pricing_rules/{id} | Get a pricing rule
[**ListLLMPricingRules**](LlmpricingrulesAPI.md#ListLLMPricingRules) | **Get** /api/v1/llm_pricing_rules | List pricing rules



## CreateOrUpdateLLMPricingRules

> CreateOrUpdateLLMPricingRules(ctx).LlmpricingruletypesUpdatableLLMPricingRules(llmpricingruletypesUpdatableLLMPricingRules).Execute()

Create or update pricing rules



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
	llmpricingruletypesUpdatableLLMPricingRules := *openapiclient.NewLlmpricingruletypesUpdatableLLMPricingRules([]openapiclient.LlmpricingruletypesUpdatableLLMPricingRule{*openapiclient.NewLlmpricingruletypesUpdatableLLMPricingRule(false, "ModelName_example", []string{"ModelPattern_example"}, *openapiclient.NewLlmpricingruletypesLLMRulePricing(float64(123), float64(123)), "Provider_example", openapiclient.LlmpricingruletypesLLMPricingRuleUnit("per_million_tokens"))}) // LlmpricingruletypesUpdatableLLMPricingRules |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LlmpricingrulesAPI.CreateOrUpdateLLMPricingRules(context.Background()).LlmpricingruletypesUpdatableLLMPricingRules(llmpricingruletypesUpdatableLLMPricingRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmpricingrulesAPI.CreateOrUpdateLLMPricingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateLLMPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **llmpricingruletypesUpdatableLLMPricingRules** | [**LlmpricingruletypesUpdatableLLMPricingRules**](LlmpricingruletypesUpdatableLLMPricingRules.md) |  | 

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


## DeleteLLMPricingRule

> DeleteLLMPricingRule(ctx, id).Execute()

Delete a pricing rule



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
	r, err := apiClient.LlmpricingrulesAPI.DeleteLLMPricingRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmpricingrulesAPI.DeleteLLMPricingRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLLMPricingRuleRequest struct via the builder pattern


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


## GetLLMPricingRule

> GetLLMPricingRule200Response GetLLMPricingRule(ctx, id).Execute()

Get a pricing rule



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
	resp, r, err := apiClient.LlmpricingrulesAPI.GetLLMPricingRule(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmpricingrulesAPI.GetLLMPricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLLMPricingRule`: GetLLMPricingRule200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmpricingrulesAPI.GetLLMPricingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLLMPricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLLMPricingRule200Response**](GetLLMPricingRule200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLLMPricingRules

> ListLLMPricingRules200Response ListLLMPricingRules(ctx).Offset(offset).Limit(limit).Execute()

List pricing rules



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
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmpricingrulesAPI.ListLLMPricingRules(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmpricingrulesAPI.ListLLMPricingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLLMPricingRules`: ListLLMPricingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `LlmpricingrulesAPI.ListLLMPricingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLLMPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ListLLMPricingRules200Response**](ListLLMPricingRules200Response.md)

### Authorization

[api_key](../README.md#api_key), [tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

