# ConfigHTTPClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**ConfigAuthorization**](ConfigAuthorization.md) |  | [optional] 
**BasicAuth** | Pointer to [**ConfigBasicAuth**](ConfigBasicAuth.md) |  | [optional] 
**BearerToken** | Pointer to **string** |  | [optional] 
**BearerTokenFile** | Pointer to **string** |  | [optional] 
**EnableHttp2** | Pointer to **bool** |  | [optional] 
**FollowRedirects** | Pointer to **bool** |  | [optional] 
**HttpHeaders** | Pointer to **map[string]interface{}** |  | [optional] 
**NoProxy** | Pointer to **string** |  | [optional] 
**Oauth2** | Pointer to [**ConfigOAuth2**](ConfigOAuth2.md) |  | [optional] 
**ProxyConnectHeader** | Pointer to **map[string][]string** |  | [optional] 
**ProxyFromEnvironment** | Pointer to **bool** |  | [optional] 
**ProxyUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**TlsConfig** | Pointer to [**ConfigTLSConfig**](ConfigTLSConfig.md) |  | [optional] 

## Methods

### NewConfigHTTPClientConfig

`func NewConfigHTTPClientConfig() *ConfigHTTPClientConfig`

NewConfigHTTPClientConfig instantiates a new ConfigHTTPClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigHTTPClientConfigWithDefaults

`func NewConfigHTTPClientConfigWithDefaults() *ConfigHTTPClientConfig`

NewConfigHTTPClientConfigWithDefaults instantiates a new ConfigHTTPClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *ConfigHTTPClientConfig) GetAuthorization() ConfigAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ConfigHTTPClientConfig) GetAuthorizationOk() (*ConfigAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ConfigHTTPClientConfig) SetAuthorization(v ConfigAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ConfigHTTPClientConfig) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetBasicAuth

`func (o *ConfigHTTPClientConfig) GetBasicAuth() ConfigBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *ConfigHTTPClientConfig) GetBasicAuthOk() (*ConfigBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *ConfigHTTPClientConfig) SetBasicAuth(v ConfigBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *ConfigHTTPClientConfig) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBearerToken

`func (o *ConfigHTTPClientConfig) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *ConfigHTTPClientConfig) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *ConfigHTTPClientConfig) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *ConfigHTTPClientConfig) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetBearerTokenFile

`func (o *ConfigHTTPClientConfig) GetBearerTokenFile() string`

GetBearerTokenFile returns the BearerTokenFile field if non-nil, zero value otherwise.

### GetBearerTokenFileOk

`func (o *ConfigHTTPClientConfig) GetBearerTokenFileOk() (*string, bool)`

GetBearerTokenFileOk returns a tuple with the BearerTokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerTokenFile

`func (o *ConfigHTTPClientConfig) SetBearerTokenFile(v string)`

SetBearerTokenFile sets BearerTokenFile field to given value.

### HasBearerTokenFile

`func (o *ConfigHTTPClientConfig) HasBearerTokenFile() bool`

HasBearerTokenFile returns a boolean if a field has been set.

### GetEnableHttp2

`func (o *ConfigHTTPClientConfig) GetEnableHttp2() bool`

GetEnableHttp2 returns the EnableHttp2 field if non-nil, zero value otherwise.

### GetEnableHttp2Ok

`func (o *ConfigHTTPClientConfig) GetEnableHttp2Ok() (*bool, bool)`

GetEnableHttp2Ok returns a tuple with the EnableHttp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttp2

`func (o *ConfigHTTPClientConfig) SetEnableHttp2(v bool)`

SetEnableHttp2 sets EnableHttp2 field to given value.

### HasEnableHttp2

`func (o *ConfigHTTPClientConfig) HasEnableHttp2() bool`

HasEnableHttp2 returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *ConfigHTTPClientConfig) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *ConfigHTTPClientConfig) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *ConfigHTTPClientConfig) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *ConfigHTTPClientConfig) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *ConfigHTTPClientConfig) GetHttpHeaders() map[string]interface{}`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *ConfigHTTPClientConfig) GetHttpHeadersOk() (*map[string]interface{}, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *ConfigHTTPClientConfig) SetHttpHeaders(v map[string]interface{})`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *ConfigHTTPClientConfig) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetNoProxy

`func (o *ConfigHTTPClientConfig) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *ConfigHTTPClientConfig) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *ConfigHTTPClientConfig) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *ConfigHTTPClientConfig) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetOauth2

`func (o *ConfigHTTPClientConfig) GetOauth2() ConfigOAuth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *ConfigHTTPClientConfig) GetOauth2Ok() (*ConfigOAuth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *ConfigHTTPClientConfig) SetOauth2(v ConfigOAuth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *ConfigHTTPClientConfig) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetProxyConnectHeader

`func (o *ConfigHTTPClientConfig) GetProxyConnectHeader() map[string][]string`

GetProxyConnectHeader returns the ProxyConnectHeader field if non-nil, zero value otherwise.

### GetProxyConnectHeaderOk

`func (o *ConfigHTTPClientConfig) GetProxyConnectHeaderOk() (*map[string][]string, bool)`

GetProxyConnectHeaderOk returns a tuple with the ProxyConnectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConnectHeader

`func (o *ConfigHTTPClientConfig) SetProxyConnectHeader(v map[string][]string)`

SetProxyConnectHeader sets ProxyConnectHeader field to given value.

### HasProxyConnectHeader

`func (o *ConfigHTTPClientConfig) HasProxyConnectHeader() bool`

HasProxyConnectHeader returns a boolean if a field has been set.

### GetProxyFromEnvironment

`func (o *ConfigHTTPClientConfig) GetProxyFromEnvironment() bool`

GetProxyFromEnvironment returns the ProxyFromEnvironment field if non-nil, zero value otherwise.

### GetProxyFromEnvironmentOk

`func (o *ConfigHTTPClientConfig) GetProxyFromEnvironmentOk() (*bool, bool)`

GetProxyFromEnvironmentOk returns a tuple with the ProxyFromEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyFromEnvironment

`func (o *ConfigHTTPClientConfig) SetProxyFromEnvironment(v bool)`

SetProxyFromEnvironment sets ProxyFromEnvironment field to given value.

### HasProxyFromEnvironment

`func (o *ConfigHTTPClientConfig) HasProxyFromEnvironment() bool`

HasProxyFromEnvironment returns a boolean if a field has been set.

### GetProxyUrl

`func (o *ConfigHTTPClientConfig) GetProxyUrl() map[string]interface{}`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *ConfigHTTPClientConfig) GetProxyUrlOk() (*map[string]interface{}, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *ConfigHTTPClientConfig) SetProxyUrl(v map[string]interface{})`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *ConfigHTTPClientConfig) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetTlsConfig

`func (o *ConfigHTTPClientConfig) GetTlsConfig() ConfigTLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *ConfigHTTPClientConfig) GetTlsConfigOk() (*ConfigTLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *ConfigHTTPClientConfig) SetTlsConfig(v ConfigTLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *ConfigHTTPClientConfig) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


