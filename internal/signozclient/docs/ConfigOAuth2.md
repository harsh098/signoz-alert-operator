# ConfigOAuth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Claims** | Pointer to **map[string]interface{}** |  | [optional] 
**ClientCertificateKey** | Pointer to **string** |  | [optional] 
**ClientCertificateKeyFile** | Pointer to **string** |  | [optional] 
**ClientCertificateKeyId** | Pointer to **string** |  | [optional] 
**ClientCertificateKeyRef** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ClientSecretFile** | Pointer to **string** |  | [optional] 
**ClientSecretRef** | Pointer to **string** |  | [optional] 
**EndpointParams** | Pointer to **map[string]string** |  | [optional] 
**GrantType** | Pointer to **string** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**NoProxy** | Pointer to **string** |  | [optional] 
**ProxyConnectHeader** | Pointer to **map[string][]string** |  | [optional] 
**ProxyFromEnvironment** | Pointer to **bool** |  | [optional] 
**ProxyUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigOAuth2

`func NewConfigOAuth2() *ConfigOAuth2`

NewConfigOAuth2 instantiates a new ConfigOAuth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigOAuth2WithDefaults

`func NewConfigOAuth2WithDefaults() *ConfigOAuth2`

NewConfigOAuth2WithDefaults instantiates a new ConfigOAuth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *ConfigOAuth2) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ConfigOAuth2) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ConfigOAuth2) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ConfigOAuth2) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClaims

`func (o *ConfigOAuth2) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ConfigOAuth2) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ConfigOAuth2) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ConfigOAuth2) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetClientCertificateKey

`func (o *ConfigOAuth2) GetClientCertificateKey() string`

GetClientCertificateKey returns the ClientCertificateKey field if non-nil, zero value otherwise.

### GetClientCertificateKeyOk

`func (o *ConfigOAuth2) GetClientCertificateKeyOk() (*string, bool)`

GetClientCertificateKeyOk returns a tuple with the ClientCertificateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKey

`func (o *ConfigOAuth2) SetClientCertificateKey(v string)`

SetClientCertificateKey sets ClientCertificateKey field to given value.

### HasClientCertificateKey

`func (o *ConfigOAuth2) HasClientCertificateKey() bool`

HasClientCertificateKey returns a boolean if a field has been set.

### GetClientCertificateKeyFile

`func (o *ConfigOAuth2) GetClientCertificateKeyFile() string`

GetClientCertificateKeyFile returns the ClientCertificateKeyFile field if non-nil, zero value otherwise.

### GetClientCertificateKeyFileOk

`func (o *ConfigOAuth2) GetClientCertificateKeyFileOk() (*string, bool)`

GetClientCertificateKeyFileOk returns a tuple with the ClientCertificateKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyFile

`func (o *ConfigOAuth2) SetClientCertificateKeyFile(v string)`

SetClientCertificateKeyFile sets ClientCertificateKeyFile field to given value.

### HasClientCertificateKeyFile

`func (o *ConfigOAuth2) HasClientCertificateKeyFile() bool`

HasClientCertificateKeyFile returns a boolean if a field has been set.

### GetClientCertificateKeyId

`func (o *ConfigOAuth2) GetClientCertificateKeyId() string`

GetClientCertificateKeyId returns the ClientCertificateKeyId field if non-nil, zero value otherwise.

### GetClientCertificateKeyIdOk

`func (o *ConfigOAuth2) GetClientCertificateKeyIdOk() (*string, bool)`

GetClientCertificateKeyIdOk returns a tuple with the ClientCertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyId

`func (o *ConfigOAuth2) SetClientCertificateKeyId(v string)`

SetClientCertificateKeyId sets ClientCertificateKeyId field to given value.

### HasClientCertificateKeyId

`func (o *ConfigOAuth2) HasClientCertificateKeyId() bool`

HasClientCertificateKeyId returns a boolean if a field has been set.

### GetClientCertificateKeyRef

`func (o *ConfigOAuth2) GetClientCertificateKeyRef() string`

GetClientCertificateKeyRef returns the ClientCertificateKeyRef field if non-nil, zero value otherwise.

### GetClientCertificateKeyRefOk

`func (o *ConfigOAuth2) GetClientCertificateKeyRefOk() (*string, bool)`

GetClientCertificateKeyRefOk returns a tuple with the ClientCertificateKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateKeyRef

`func (o *ConfigOAuth2) SetClientCertificateKeyRef(v string)`

SetClientCertificateKeyRef sets ClientCertificateKeyRef field to given value.

### HasClientCertificateKeyRef

`func (o *ConfigOAuth2) HasClientCertificateKeyRef() bool`

HasClientCertificateKeyRef returns a boolean if a field has been set.

### GetClientId

`func (o *ConfigOAuth2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConfigOAuth2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConfigOAuth2) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConfigOAuth2) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ConfigOAuth2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ConfigOAuth2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ConfigOAuth2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ConfigOAuth2) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretFile

`func (o *ConfigOAuth2) GetClientSecretFile() string`

GetClientSecretFile returns the ClientSecretFile field if non-nil, zero value otherwise.

### GetClientSecretFileOk

`func (o *ConfigOAuth2) GetClientSecretFileOk() (*string, bool)`

GetClientSecretFileOk returns a tuple with the ClientSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretFile

`func (o *ConfigOAuth2) SetClientSecretFile(v string)`

SetClientSecretFile sets ClientSecretFile field to given value.

### HasClientSecretFile

`func (o *ConfigOAuth2) HasClientSecretFile() bool`

HasClientSecretFile returns a boolean if a field has been set.

### GetClientSecretRef

`func (o *ConfigOAuth2) GetClientSecretRef() string`

GetClientSecretRef returns the ClientSecretRef field if non-nil, zero value otherwise.

### GetClientSecretRefOk

`func (o *ConfigOAuth2) GetClientSecretRefOk() (*string, bool)`

GetClientSecretRefOk returns a tuple with the ClientSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRef

`func (o *ConfigOAuth2) SetClientSecretRef(v string)`

SetClientSecretRef sets ClientSecretRef field to given value.

### HasClientSecretRef

`func (o *ConfigOAuth2) HasClientSecretRef() bool`

HasClientSecretRef returns a boolean if a field has been set.

### GetEndpointParams

`func (o *ConfigOAuth2) GetEndpointParams() map[string]string`

GetEndpointParams returns the EndpointParams field if non-nil, zero value otherwise.

### GetEndpointParamsOk

`func (o *ConfigOAuth2) GetEndpointParamsOk() (*map[string]string, bool)`

GetEndpointParamsOk returns a tuple with the EndpointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointParams

`func (o *ConfigOAuth2) SetEndpointParams(v map[string]string)`

SetEndpointParams sets EndpointParams field to given value.

### HasEndpointParams

`func (o *ConfigOAuth2) HasEndpointParams() bool`

HasEndpointParams returns a boolean if a field has been set.

### GetGrantType

`func (o *ConfigOAuth2) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *ConfigOAuth2) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *ConfigOAuth2) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *ConfigOAuth2) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetIss

`func (o *ConfigOAuth2) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *ConfigOAuth2) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *ConfigOAuth2) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *ConfigOAuth2) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetNoProxy

`func (o *ConfigOAuth2) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *ConfigOAuth2) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *ConfigOAuth2) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *ConfigOAuth2) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetProxyConnectHeader

`func (o *ConfigOAuth2) GetProxyConnectHeader() map[string][]string`

GetProxyConnectHeader returns the ProxyConnectHeader field if non-nil, zero value otherwise.

### GetProxyConnectHeaderOk

`func (o *ConfigOAuth2) GetProxyConnectHeaderOk() (*map[string][]string, bool)`

GetProxyConnectHeaderOk returns a tuple with the ProxyConnectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConnectHeader

`func (o *ConfigOAuth2) SetProxyConnectHeader(v map[string][]string)`

SetProxyConnectHeader sets ProxyConnectHeader field to given value.

### HasProxyConnectHeader

`func (o *ConfigOAuth2) HasProxyConnectHeader() bool`

HasProxyConnectHeader returns a boolean if a field has been set.

### GetProxyFromEnvironment

`func (o *ConfigOAuth2) GetProxyFromEnvironment() bool`

GetProxyFromEnvironment returns the ProxyFromEnvironment field if non-nil, zero value otherwise.

### GetProxyFromEnvironmentOk

`func (o *ConfigOAuth2) GetProxyFromEnvironmentOk() (*bool, bool)`

GetProxyFromEnvironmentOk returns a tuple with the ProxyFromEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyFromEnvironment

`func (o *ConfigOAuth2) SetProxyFromEnvironment(v bool)`

SetProxyFromEnvironment sets ProxyFromEnvironment field to given value.

### HasProxyFromEnvironment

`func (o *ConfigOAuth2) HasProxyFromEnvironment() bool`

HasProxyFromEnvironment returns a boolean if a field has been set.

### GetProxyUrl

`func (o *ConfigOAuth2) GetProxyUrl() map[string]interface{}`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *ConfigOAuth2) GetProxyUrlOk() (*map[string]interface{}, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *ConfigOAuth2) SetProxyUrl(v map[string]interface{})`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *ConfigOAuth2) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetScopes

`func (o *ConfigOAuth2) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConfigOAuth2) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConfigOAuth2) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConfigOAuth2) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *ConfigOAuth2) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *ConfigOAuth2) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *ConfigOAuth2) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *ConfigOAuth2) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetTokenUrl

`func (o *ConfigOAuth2) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *ConfigOAuth2) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *ConfigOAuth2) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *ConfigOAuth2) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


