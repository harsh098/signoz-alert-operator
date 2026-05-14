# \SessionsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSessionByEmailPassword**](SessionsAPI.md#CreateSessionByEmailPassword) | **Post** /api/v2/sessions/email_password | Create session by email and password
[**CreateSessionByGoogleCallback**](SessionsAPI.md#CreateSessionByGoogleCallback) | **Get** /api/v1/complete/google | Create session by google callback
[**CreateSessionByOIDCCallback**](SessionsAPI.md#CreateSessionByOIDCCallback) | **Get** /api/v1/complete/oidc | Create session by oidc callback
[**CreateSessionBySAMLCallback**](SessionsAPI.md#CreateSessionBySAMLCallback) | **Post** /api/v1/complete/saml | Create session by saml callback
[**DeleteSession**](SessionsAPI.md#DeleteSession) | **Delete** /api/v2/sessions | Delete session
[**GetSessionContext**](SessionsAPI.md#GetSessionContext) | **Get** /api/v2/sessions/context | Get session context
[**RotateSession**](SessionsAPI.md#RotateSession) | **Post** /api/v2/sessions/rotate | Rotate session



## CreateSessionByEmailPassword

> CreateSessionByGoogleCallback303Response CreateSessionByEmailPassword(ctx).AuthtypesPostableEmailPasswordSession(authtypesPostableEmailPasswordSession).Execute()

Create session by email and password



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
	authtypesPostableEmailPasswordSession := *openapiclient.NewAuthtypesPostableEmailPasswordSession() // AuthtypesPostableEmailPasswordSession |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.CreateSessionByEmailPassword(context.Background()).AuthtypesPostableEmailPasswordSession(authtypesPostableEmailPasswordSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateSessionByEmailPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSessionByEmailPassword`: CreateSessionByGoogleCallback303Response
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.CreateSessionByEmailPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionByEmailPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authtypesPostableEmailPasswordSession** | [**AuthtypesPostableEmailPasswordSession**](AuthtypesPostableEmailPasswordSession.md) |  | 

### Return type

[**CreateSessionByGoogleCallback303Response**](CreateSessionByGoogleCallback303Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionByGoogleCallback

> CreateSessionByGoogleCallback(ctx).Execute()

Create session by google callback



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
	r, err := apiClient.SessionsAPI.CreateSessionByGoogleCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateSessionByGoogleCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionByGoogleCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionByOIDCCallback

> CreateSessionByOIDCCallback(ctx).Execute()

Create session by oidc callback



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
	r, err := apiClient.SessionsAPI.CreateSessionByOIDCCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateSessionByOIDCCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionByOIDCCallbackRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionBySAMLCallback

> CreateSessionBySAMLCallback(ctx).RelayState(relayState).SAMLResponse(sAMLResponse).RelayState2(relayState2).SAMLResponse2(sAMLResponse2).Execute()

Create session by saml callback



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
	relayState := "relayState_example" // string |  (optional)
	sAMLResponse := "sAMLResponse_example" // string |  (optional)
	relayState2 := "relayState_example" // string |  (optional)
	sAMLResponse2 := "sAMLResponse_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionsAPI.CreateSessionBySAMLCallback(context.Background()).RelayState(relayState).SAMLResponse(sAMLResponse).RelayState2(relayState2).SAMLResponse2(sAMLResponse2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.CreateSessionBySAMLCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionBySAMLCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relayState** | **string** |  | 
 **sAMLResponse** | **string** |  | 
 **relayState2** | **string** |  | 
 **sAMLResponse2** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> DeleteSession(ctx).Execute()

Delete session



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
	r, err := apiClient.SessionsAPI.DeleteSession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.DeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[tokenizer](../README.md#tokenizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionContext

> GetSessionContext200Response GetSessionContext(ctx).Execute()

Get session context



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
	resp, r, err := apiClient.SessionsAPI.GetSessionContext(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.GetSessionContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionContext`: GetSessionContext200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.GetSessionContext`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionContextRequest struct via the builder pattern


### Return type

[**GetSessionContext200Response**](GetSessionContext200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSession

> CreateSessionByGoogleCallback303Response RotateSession(ctx).AuthtypesPostableRotateToken(authtypesPostableRotateToken).Execute()

Rotate session



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
	authtypesPostableRotateToken := *openapiclient.NewAuthtypesPostableRotateToken() // AuthtypesPostableRotateToken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionsAPI.RotateSession(context.Background()).AuthtypesPostableRotateToken(authtypesPostableRotateToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.RotateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateSession`: CreateSessionByGoogleCallback303Response
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.RotateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authtypesPostableRotateToken** | [**AuthtypesPostableRotateToken**](AuthtypesPostableRotateToken.md) |  | 

### Return type

[**CreateSessionByGoogleCallback303Response**](CreateSessionByGoogleCallback303Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

