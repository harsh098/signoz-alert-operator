# ConfigPushoverConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **bool** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Monospace** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Retry** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Sound** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlTitle** | Pointer to **string** |  | [optional] 
**UserKey** | Pointer to **string** |  | [optional] 
**UserKeyFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigPushoverConfig

`func NewConfigPushoverConfig() *ConfigPushoverConfig`

NewConfigPushoverConfig instantiates a new ConfigPushoverConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPushoverConfigWithDefaults

`func NewConfigPushoverConfigWithDefaults() *ConfigPushoverConfig`

NewConfigPushoverConfigWithDefaults instantiates a new ConfigPushoverConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ConfigPushoverConfig) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ConfigPushoverConfig) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ConfigPushoverConfig) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ConfigPushoverConfig) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetExpire

`func (o *ConfigPushoverConfig) GetExpire() string`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ConfigPushoverConfig) GetExpireOk() (*string, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ConfigPushoverConfig) SetExpire(v string)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ConfigPushoverConfig) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetHtml

`func (o *ConfigPushoverConfig) GetHtml() bool`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *ConfigPushoverConfig) GetHtmlOk() (*bool, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *ConfigPushoverConfig) SetHtml(v bool)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *ConfigPushoverConfig) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigPushoverConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigPushoverConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigPushoverConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigPushoverConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigPushoverConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigPushoverConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigPushoverConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigPushoverConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonospace

`func (o *ConfigPushoverConfig) GetMonospace() bool`

GetMonospace returns the Monospace field if non-nil, zero value otherwise.

### GetMonospaceOk

`func (o *ConfigPushoverConfig) GetMonospaceOk() (*bool, bool)`

GetMonospaceOk returns a tuple with the Monospace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonospace

`func (o *ConfigPushoverConfig) SetMonospace(v bool)`

SetMonospace sets Monospace field to given value.

### HasMonospace

`func (o *ConfigPushoverConfig) HasMonospace() bool`

HasMonospace returns a boolean if a field has been set.

### GetPriority

`func (o *ConfigPushoverConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConfigPushoverConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConfigPushoverConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConfigPushoverConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetry

`func (o *ConfigPushoverConfig) GetRetry() string`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *ConfigPushoverConfig) GetRetryOk() (*string, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *ConfigPushoverConfig) SetRetry(v string)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *ConfigPushoverConfig) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigPushoverConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigPushoverConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigPushoverConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigPushoverConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSound

`func (o *ConfigPushoverConfig) GetSound() string`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *ConfigPushoverConfig) GetSoundOk() (*string, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *ConfigPushoverConfig) SetSound(v string)`

SetSound sets Sound field to given value.

### HasSound

`func (o *ConfigPushoverConfig) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigPushoverConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigPushoverConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigPushoverConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigPushoverConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToken

`func (o *ConfigPushoverConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfigPushoverConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfigPushoverConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConfigPushoverConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenFile

`func (o *ConfigPushoverConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *ConfigPushoverConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *ConfigPushoverConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *ConfigPushoverConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.

### GetTtl

`func (o *ConfigPushoverConfig) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ConfigPushoverConfig) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ConfigPushoverConfig) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ConfigPushoverConfig) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigPushoverConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigPushoverConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigPushoverConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConfigPushoverConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlTitle

`func (o *ConfigPushoverConfig) GetUrlTitle() string`

GetUrlTitle returns the UrlTitle field if non-nil, zero value otherwise.

### GetUrlTitleOk

`func (o *ConfigPushoverConfig) GetUrlTitleOk() (*string, bool)`

GetUrlTitleOk returns a tuple with the UrlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTitle

`func (o *ConfigPushoverConfig) SetUrlTitle(v string)`

SetUrlTitle sets UrlTitle field to given value.

### HasUrlTitle

`func (o *ConfigPushoverConfig) HasUrlTitle() bool`

HasUrlTitle returns a boolean if a field has been set.

### GetUserKey

`func (o *ConfigPushoverConfig) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *ConfigPushoverConfig) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *ConfigPushoverConfig) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.

### HasUserKey

`func (o *ConfigPushoverConfig) HasUserKey() bool`

HasUserKey returns a boolean if a field has been set.

### GetUserKeyFile

`func (o *ConfigPushoverConfig) GetUserKeyFile() string`

GetUserKeyFile returns the UserKeyFile field if non-nil, zero value otherwise.

### GetUserKeyFileOk

`func (o *ConfigPushoverConfig) GetUserKeyFileOk() (*string, bool)`

GetUserKeyFileOk returns a tuple with the UserKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyFile

`func (o *ConfigPushoverConfig) SetUserKeyFile(v string)`

SetUserKeyFile sets UserKeyFile field to given value.

### HasUserKeyFile

`func (o *ConfigPushoverConfig) HasUserKeyFile() bool`

HasUserKeyFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


