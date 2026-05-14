# ConfigMSTeamsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**WebhookUrlFile** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigMSTeamsConfig

`func NewConfigMSTeamsConfig() *ConfigMSTeamsConfig`

NewConfigMSTeamsConfig instantiates a new ConfigMSTeamsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMSTeamsConfigWithDefaults

`func NewConfigMSTeamsConfigWithDefaults() *ConfigMSTeamsConfig`

NewConfigMSTeamsConfigWithDefaults instantiates a new ConfigMSTeamsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpConfig

`func (o *ConfigMSTeamsConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigMSTeamsConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigMSTeamsConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigMSTeamsConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigMSTeamsConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigMSTeamsConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigMSTeamsConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigMSTeamsConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSummary

`func (o *ConfigMSTeamsConfig) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConfigMSTeamsConfig) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConfigMSTeamsConfig) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConfigMSTeamsConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetText

`func (o *ConfigMSTeamsConfig) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConfigMSTeamsConfig) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConfigMSTeamsConfig) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConfigMSTeamsConfig) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *ConfigMSTeamsConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfigMSTeamsConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfigMSTeamsConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfigMSTeamsConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *ConfigMSTeamsConfig) GetWebhookUrl() map[string]interface{}`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ConfigMSTeamsConfig) GetWebhookUrlOk() (*map[string]interface{}, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ConfigMSTeamsConfig) SetWebhookUrl(v map[string]interface{})`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ConfigMSTeamsConfig) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlFile

`func (o *ConfigMSTeamsConfig) GetWebhookUrlFile() string`

GetWebhookUrlFile returns the WebhookUrlFile field if non-nil, zero value otherwise.

### GetWebhookUrlFileOk

`func (o *ConfigMSTeamsConfig) GetWebhookUrlFileOk() (*string, bool)`

GetWebhookUrlFileOk returns a tuple with the WebhookUrlFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlFile

`func (o *ConfigMSTeamsConfig) SetWebhookUrlFile(v string)`

SetWebhookUrlFile sets WebhookUrlFile field to given value.

### HasWebhookUrlFile

`func (o *ConfigMSTeamsConfig) HasWebhookUrlFile() bool`

HasWebhookUrlFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


