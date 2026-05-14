# ConfigDiscordConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigDiscordConfig

`func NewConfigDiscordConfig() *ConfigDiscordConfig`

NewConfigDiscordConfig instantiates a new ConfigDiscordConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDiscordConfigWithDefaults

`func NewConfigDiscordConfigWithDefaults() *ConfigDiscordConfig`

NewConfigDiscordConfigWithDefaults instantiates a new ConfigDiscordConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *ConfigDiscordConfig) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ConfigDiscordConfig) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ConfigDiscordConfig) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ConfigDiscordConfig) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetContent

`func (o *ConfigDiscordConfig) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConfigDiscordConfig) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConfigDiscordConfig) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConfigDiscordConfig) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigDiscordConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigDiscordConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigDiscordConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigDiscordConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigDiscordConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigDiscordConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigDiscordConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigDiscordConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigDiscordConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigDiscordConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigDiscordConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigDiscordConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigDiscordConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigDiscordConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigDiscordConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigDiscordConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsername

`func (o *ConfigDiscordConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConfigDiscordConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConfigDiscordConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConfigDiscordConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ConfigDiscordConfig) GetWebhookUrl() map[string]interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ConfigDiscordConfig) GetWebhookUrlOk() (*map[string]interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ConfigDiscordConfig) SetWebhookUrl(v map[string]interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ConfigDiscordConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlFile

`func (o *ConfigDiscordConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *ConfigDiscordConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *ConfigDiscordConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *ConfigDiscordConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


