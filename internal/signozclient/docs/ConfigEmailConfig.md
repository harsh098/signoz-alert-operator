# ConfigEmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthIdentity** | Pointer to **string** |  | [optional] 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthPasswordFile** | Pointer to **string** |  | [optional] 
**AuthSecret** | Pointer to **string** |  | [optional] 
**AuthSecretFile** | Pointer to **string** |  | [optional] 
**AuthUsername** | Pointer to **string** |  | [optional] 
**ForceImplicitTls** | Pointer to **NullableBool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Hello** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **string** |  | [optional] 
**RequireTls** | Pointer to **NullableBool** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Smarthost** | Pointer to **map[string]interface{}** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Threading** | Pointer to [**ConfigThreadingConfig**](ConfigThreadingConfig.md) |  | [optional] 
**TlsConfig** | Pointer to [**ConfigTLSConfig**](ConfigTLSConfig.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigEmailConfig

`func NewConfigEmailConfig() *ConfigEmailConfig`

NewConfigEmailConfig instantiates a new ConfigEmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigEmailConfigWithDefaults

`func NewConfigEmailConfigWithDefaults() *ConfigEmailConfig`

NewConfigEmailConfigWithDefaults instantiates a new ConfigEmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthIdentity

`func (o *ConfigEmailConfig) GetAuthIdentity() string`

GetAuthIdentity returns the AuthIdentity field if non-nil, zero value otherwise.

### GetAuthIdentityOk

`func (o *ConfigEmailConfig) GetAuthIdentityOk() (*string, bool)`

GetAuthIdentityOk returns a tuple with the AuthIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdentity

`func (o *ConfigEmailConfig) SetAuthIdentity(v string)`

SetAuthIdentity sets AuthIdentity field to given value.

### HasAuthIdentity

`func (o *ConfigEmailConfig) HasAuthIdentity() bool`

HasAuthIdentity returns a boolean if a field has been set.

### GetAuthPassword

`func (o *ConfigEmailConfig) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *ConfigEmailConfig) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *ConfigEmailConfig) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *ConfigEmailConfig) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthPasswordFile

`func (o *ConfigEmailConfig) GetAuthPasswordFile() string`

GetAuthPasswordFile returns the AuthPasswordFile field if non-nil, zero value otherwise.

### GetAuthPasswordFileOk

`func (o *ConfigEmailConfig) GetAuthPasswordFileOk() (*string, bool)`

GetAuthPasswordFileOk returns a tuple with the AuthPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPasswordFile

`func (o *ConfigEmailConfig) SetAuthPasswordFile(v string)`

SetAuthPasswordFile sets AuthPasswordFile field to given value.

### HasAuthPasswordFile

`func (o *ConfigEmailConfig) HasAuthPasswordFile() bool`

HasAuthPasswordFile returns a boolean if a field has been set.

### GetAuthSecret

`func (o *ConfigEmailConfig) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *ConfigEmailConfig) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *ConfigEmailConfig) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *ConfigEmailConfig) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### GetAuthSecretFile

`func (o *ConfigEmailConfig) GetAuthSecretFile() string`

GetAuthSecretFile returns the AuthSecretFile field if non-nil, zero value otherwise.

### GetAuthSecretFileOk

`func (o *ConfigEmailConfig) GetAuthSecretFileOk() (*string, bool)`

GetAuthSecretFileOk returns a tuple with the AuthSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecretFile

`func (o *ConfigEmailConfig) SetAuthSecretFile(v string)`

SetAuthSecretFile sets AuthSecretFile field to given value.

### HasAuthSecretFile

`func (o *ConfigEmailConfig) HasAuthSecretFile() bool`

HasAuthSecretFile returns a boolean if a field has been set.

### GetAuthUsername

`func (o *ConfigEmailConfig) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *ConfigEmailConfig) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *ConfigEmailConfig) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *ConfigEmailConfig) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetForceImplicitTls

`func (o *ConfigEmailConfig) GetForceImplicitTls() bool`

GetForceImplicitTls returns the ForceImplicitTls field if non-nil, zero value otherwise.

### GetForceImplicitTlsOk

`func (o *ConfigEmailConfig) GetForceImplicitTlsOk() (*bool, bool)`

GetForceImplicitTlsOk returns a tuple with the ForceImplicitTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceImplicitTls

`func (o *ConfigEmailConfig) SetForceImplicitTls(v bool)`

SetForceImplicitTls sets ForceImplicitTls field to given value.

### HasForceImplicitTls

`func (o *ConfigEmailConfig) HasForceImplicitTls() bool`

HasForceImplicitTls returns a boolean if a field has been set.

### SetForceImplicitTlsNil

`func (o *ConfigEmailConfig) SetForceImplicitTlsNil(b bool)`

 SetForceImplicitTlsNil sets the value for ForceImplicitTls to be an explicit nil

### UnsetForceImplicitTls
`func (o *ConfigEmailConfig) UnsetForceImplicitTls()`

UnsetForceImplicitTls ensures that no value is present for ForceImplicitTls, not even an explicit nil
### GetFrom

`func (o *ConfigEmailConfig) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ConfigEmailConfig) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ConfigEmailConfig) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ConfigEmailConfig) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *ConfigEmailConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConfigEmailConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConfigEmailConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConfigEmailConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHello

`func (o *ConfigEmailConfig) GetHello() string`

GetHello returns the Hello field if non-nil, zero value otherwise.

### GetHelloOk

`func (o *ConfigEmailConfig) GetHelloOk() (*string, bool)`

GetHelloOk returns a tuple with the Hello field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHello

`func (o *ConfigEmailConfig) SetHello(v string)`

SetHello sets Hello field to given value.

### HasHello

`func (o *ConfigEmailConfig) HasHello() bool`

HasHello returns a boolean if a field has been set.

### GetHtml

`func (o *ConfigEmailConfig) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *ConfigEmailConfig) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *ConfigEmailConfig) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *ConfigEmailConfig) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetRequireTls

`func (o *ConfigEmailConfig) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *ConfigEmailConfig) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *ConfigEmailConfig) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.

### HasRequireTls

`func (o *ConfigEmailConfig) HasRequireTls() bool`

HasRequireTls returns a boolean if a field has been set.

### SetRequireTlsNil

`func (o *ConfigEmailConfig) SetRequireTlsNil(b bool)`

 SetRequireTlsNil sets the value for RequireTls to be an explicit nil

### UnsetRequireTls
`func (o *ConfigEmailConfig) UnsetRequireTls()`

UnsetRequireTls ensures that no value is present for RequireTls, not even an explicit nil
### GetSendResolved

`func (o *ConfigEmailConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigEmailConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigEmailConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigEmailConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSmarthost

`func (o *ConfigEmailConfig) GetSmarthost() map[string]interface{}`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *ConfigEmailConfig) GetSmarthostOk() (*map[string]interface{}, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *ConfigEmailConfig) SetSmarthost(v map[string]interface{})`

SetSmarthost sets Smarthost field to given value.

### HasSmarthost

`func (o *ConfigEmailConfig) HasSmarthost() bool`

HasSmarthost returns a boolean if a field has been set.

### GetText

`func (o *ConfigEmailConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigEmailConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigEmailConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigEmailConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThreading

`func (o *ConfigEmailConfig) GetThreading() ConfigThreadingConfig`

GetThreading returns the Threading field if non-nil, zero value otherwise.

### GetThreadingOk

`func (o *ConfigEmailConfig) GetThreadingOk() (*ConfigThreadingConfig, bool)`

GetThreadingOk returns a tuple with the Threading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreading

`func (o *ConfigEmailConfig) SetThreading(v ConfigThreadingConfig)`

SetThreading sets Threading field to given value.

### HasThreading

`func (o *ConfigEmailConfig) HasThreading() bool`

HasThreading returns a boolean if a field has been set.

### GetTlsConfig

`func (o *ConfigEmailConfig) GetTlsConfig() ConfigTLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *ConfigEmailConfig) GetTlsConfigOk() (*ConfigTLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *ConfigEmailConfig) SetTlsConfig(v ConfigTLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *ConfigEmailConfig) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetTo

`func (o *ConfigEmailConfig) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ConfigEmailConfig) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ConfigEmailConfig) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ConfigEmailConfig) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


