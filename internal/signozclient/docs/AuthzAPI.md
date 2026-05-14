# \AuthzAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthzCheck**](AuthzAPI.md#AuthzCheck) | **Post** /api/v1/authz/check | Check permissions



## AuthzCheck

> AuthzCheck200Response AuthzCheck(ctx).AuthtypesTransaction(authtypesTransaction).Execute()

Check permissions



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
	authtypesTransaction := []openapiclient.AuthtypesTransaction{*openapiclient.NewAuthtypesTransaction(*openapiclient.NewCoretypesObject(*openapiclient.NewCoretypesResourceRef("Kind_example", openapiclient.CoretypesType("user")), "Selector_example"), openapiclient.AuthtypesRelation("create"))} // []AuthtypesTransaction |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthzAPI.AuthzCheck(context.Background()).AuthtypesTransaction(authtypesTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthzAPI.AuthzCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzCheck`: AuthzCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthzAPI.AuthzCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthzCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authtypesTransaction** | [**[]AuthtypesTransaction**](AuthtypesTransaction.md) |  | 

### Return type

[**AuthzCheck200Response**](AuthzCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

