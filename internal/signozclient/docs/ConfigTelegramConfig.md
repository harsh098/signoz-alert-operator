# ConfigTelegramConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Chat** | Pointer to **int64** |  | [optional] 
**ChatFile** | Pointer to **string** |  | [optional] 
**DisableNotifications** | Pointer to **bool** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageThreadId** | Pointer to **int32** |  | [optional] 
**ParseMode** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigTelegramConfig

`func NewConfigTelegramConfig() *ConfigTelegramConfig`

NewConfigTelegramConfig instantiates a new ConfigTelegramConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTelegramConfigWithDefaults

`func NewConfigTelegramConfigWithDefaults() *ConfigTelegramConfig`

NewConfigTelegramConfigWithDefaults instantiates a new ConfigTelegramConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiUrl

`func (o *ConfigTelegramConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigTelegramConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigTelegramConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigTelegramConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetChat

`func (o *ConfigTelegramConfig) GetChat() int64`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ConfigTelegramConfig) GetChatOk() (*int64, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ConfigTelegramConfig) SetChat(v int64)`

SetChat sets Chat field to given value.

### HasChat

`func (o *ConfigTelegramConfig) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetChatFile

`func (o *ConfigTelegramConfig) GetChatFile() string`

GetChatFile returns the ChatFile field if non-nil, zero value otherwise.

### GetChatFileOk

`func (o *ConfigTelegramConfig) GetChatFileOk() (*string, bool)`

GetChatFileOk returns a tuple with the ChatFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatFile

`func (o *ConfigTelegramConfig) SetChatFile(v string)`

SetChatFile sets ChatFile field to given value.

### HasChatFile

`func (o *ConfigTelegramConfig) HasChatFile() bool`

HasChatFile returns a boolean if a field has been set.

### GetDisableNotifications

`func (o *ConfigTelegramConfig) GetDisableNotifications() bool`

GetDisableNotifications returns the DisableNotifications field if non-nil, zero value otherwise.

### GetDisableNotificationsOk

`func (o *ConfigTelegramConfig) GetDisableNotificationsOk() (*bool, bool)`

GetDisableNotificationsOk returns a tuple with the DisableNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotifications

`func (o *ConfigTelegramConfig) SetDisableNotifications(v bool)`

SetDisableNotifications sets DisableNotifications field to given value.

### HasDisableNotifications

`func (o *ConfigTelegramConfig) HasDisableNotifications() bool`

HasDisableNotifications returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigTelegramConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigTelegramConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigTelegramConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigTelegramConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigTelegramConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigTelegramConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigTelegramConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigTelegramConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageThreadId

`func (o *ConfigTelegramConfig) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *ConfigTelegramConfig) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *ConfigTelegramConfig) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *ConfigTelegramConfig) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetParseMode

`func (o *ConfigTelegramConfig) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *ConfigTelegramConfig) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *ConfigTelegramConfig) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *ConfigTelegramConfig) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigTelegramConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigTelegramConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigTelegramConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigTelegramConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetToken

`func (o *ConfigTelegramConfig) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfigTelegramConfig) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfigTelegramConfig) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConfigTelegramConfig) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenFile

`func (o *ConfigTelegramConfig) GetTokenFile() string`

GetTokenFile returns the TokenFile field if non-nil, zero value otherwise.

### GetTokenFileOk

`func (o *ConfigTelegramConfig) GetTokenFileOk() (*string, bool)`

GetTokenFileOk returns a tuple with the TokenFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFile

`func (o *ConfigTelegramConfig) SetTokenFile(v string)`

SetTokenFile sets TokenFile field to given value.

### HasTokenFile

`func (o *ConfigTelegramConfig) HasTokenFile() bool`

HasTokenFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


