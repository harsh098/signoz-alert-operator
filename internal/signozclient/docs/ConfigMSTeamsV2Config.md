# ConfigMSTeamsV2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigMSTeamsV2Config

`func NewConfigMSTeamsV2Config() *ConfigMSTeamsV2Config`

NewConfigMSTeamsV2Config instantiates a new ConfigMSTeamsV2Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMSTeamsV2ConfigWithDefaults

`func NewConfigMSTeamsV2ConfigWithDefaults() *ConfigMSTeamsV2Config`

NewConfigMSTeamsV2ConfigWithDefaults instantiates a new ConfigMSTeamsV2Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpConfig

`func (o *ConfigMSTeamsV2Config) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigMSTeamsV2Config) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigMSTeamsV2Config) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigMSTeamsV2Config) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigMSTeamsV2Config) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigMSTeamsV2Config) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigMSTeamsV2Config) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigMSTeamsV2Config) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetText

`func (o *ConfigMSTeamsV2Config) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigMSTeamsV2Config) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigMSTeamsV2Config) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigMSTeamsV2Config) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigMSTeamsV2Config) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigMSTeamsV2Config) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigMSTeamsV2Config) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigMSTeamsV2Config) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ConfigMSTeamsV2Config) GetWebhookUrl() map[string]interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ConfigMSTeamsV2Config) GetWebhookUrlOk() (*map[string]interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ConfigMSTeamsV2Config) SetWebhookUrl(v map[string]interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ConfigMSTeamsV2Config) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlFile

`func (o *ConfigMSTeamsV2Config) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *ConfigMSTeamsV2Config) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *ConfigMSTeamsV2Config) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *ConfigMSTeamsV2Config) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


