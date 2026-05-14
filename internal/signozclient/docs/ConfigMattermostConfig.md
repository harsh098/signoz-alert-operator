# ConfigMattermostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]ConfigMattermostAttachment**](ConfigMattermostAttachment.md) |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**IconEmoji** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to [**ConfigMattermostPriority**](ConfigMattermostPriority.md) |  | [optional] 
**Props** | Pointer to [**ConfigMattermostProps**](ConfigMattermostProps.md) |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigMattermostConfig

`func NewConfigMattermostConfig() *ConfigMattermostConfig`

NewConfigMattermostConfig instantiates a new ConfigMattermostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMattermostConfigWithDefaults

`func NewConfigMattermostConfigWithDefaults() *ConfigMattermostConfig`

NewConfigMattermostConfigWithDefaults instantiates a new ConfigMattermostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *ConfigMattermostConfig) GetAttachments() []ConfigMattermostAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ConfigMattermostConfig) GetAttachmentsOk() (*[]ConfigMattermostAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ConfigMattermostConfig) SetAttachments(v []ConfigMattermostAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ConfigMattermostConfig) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetChannel

`func (o *ConfigMattermostConfig) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ConfigMattermostConfig) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ConfigMattermostConfig) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ConfigMattermostConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigMattermostConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigMattermostConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigMattermostConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigMattermostConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIconEmoji

`func (o *ConfigMattermostConfig) GetIconEmoji() string`

GetIconEmoji returns the IconEmoji field if non-nil, zero value otherwise.

### GetIconEmojiOk

`func (o *ConfigMattermostConfig) GetIconEmojiOk() (*string, bool)`

GetIconEmojiOk returns a tuple with the IconEmoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconEmoji

`func (o *ConfigMattermostConfig) SetIconEmoji(v string)`

SetIconEmoji sets IconEmoji field to given value.

### HasIconEmoji

`func (o *ConfigMattermostConfig) HasIconEmoji() bool`

HasIconEmoji returns a boolean if a field has been set.

### GetIconUrl

`func (o *ConfigMattermostConfig) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConfigMattermostConfig) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConfigMattermostConfig) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConfigMattermostConfig) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetPriority

`func (o *ConfigMattermostConfig) GetPriority() ConfigMattermostPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConfigMattermostConfig) GetPriorityOk() (*ConfigMattermostPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConfigMattermostConfig) SetPriority(v ConfigMattermostPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConfigMattermostConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProps

`func (o *ConfigMattermostConfig) GetProps() ConfigMattermostProps`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *ConfigMattermostConfig) GetPropsOk() (*ConfigMattermostProps, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *ConfigMattermostConfig) SetProps(v ConfigMattermostProps)`

SetProps sets Props field to given value.

### HasProps

`func (o *ConfigMattermostConfig) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigMattermostConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigMattermostConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigMattermostConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigMattermostConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetText

`func (o *ConfigMattermostConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigMattermostConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigMattermostConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigMattermostConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *ConfigMattermostConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigMattermostConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigMattermostConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigMattermostConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *ConfigMattermostConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConfigMattermostConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConfigMattermostConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConfigMattermostConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ConfigMattermostConfig) GetWebhookUrl() map[string]interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ConfigMattermostConfig) GetWebhookUrlOk() (*map[string]interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ConfigMattermostConfig) SetWebhookUrl(v map[string]interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ConfigMattermostConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlFile

`func (o *ConfigMattermostConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *ConfigMattermostConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *ConfigMattermostConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *ConfigMattermostConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


