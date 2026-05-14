# ConfigOpsGenieConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiKeyFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Responders** | Pointer to [**[]ConfigOpsGenieConfigResponder**](ConfigOpsGenieConfigResponder.md) |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdateAlerts** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfigOpsGenieConfig

`func NewConfigOpsGenieConfig() *ConfigOpsGenieConfig`

NewConfigOpsGenieConfig instantiates a new ConfigOpsGenieConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigOpsGenieConfigWithDefaults

`func NewConfigOpsGenieConfigWithDefaults() *ConfigOpsGenieConfig`

NewConfigOpsGenieConfigWithDefaults instantiates a new ConfigOpsGenieConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ConfigOpsGenieConfig) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConfigOpsGenieConfig) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConfigOpsGenieConfig) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ConfigOpsGenieConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiKey

`func (o *ConfigOpsGenieConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ConfigOpsGenieConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ConfigOpsGenieConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ConfigOpsGenieConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeyFile

`func (o *ConfigOpsGenieConfig) GetApiKeyFile() string`

GetApiKeyFile returns the ApiKeyFile field if non-nil, zero value otherwise.

### GetApiKeyFileOk

`func (o *ConfigOpsGenieConfig) GetApiKeyFileOk() (*string, bool)`

GetApiKeyFileOk returns a tuple with the ApiKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyFile

`func (o *ConfigOpsGenieConfig) SetApiKeyFile(v string)`

SetApiKeyFile sets ApiKeyFile field to given value.

### HasApiKeyFile

`func (o *ConfigOpsGenieConfig) HasApiKeyFile() bool`

HasApiKeyFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigOpsGenieConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigOpsGenieConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigOpsGenieConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigOpsGenieConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigOpsGenieConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigOpsGenieConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigOpsGenieConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigOpsGenieConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *ConfigOpsGenieConfig) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConfigOpsGenieConfig) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConfigOpsGenieConfig) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConfigOpsGenieConfig) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEntity

`func (o *ConfigOpsGenieConfig) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ConfigOpsGenieConfig) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ConfigOpsGenieConfig) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ConfigOpsGenieConfig) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigOpsGenieConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigOpsGenieConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigOpsGenieConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigOpsGenieConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *ConfigOpsGenieConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConfigOpsGenieConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConfigOpsGenieConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConfigOpsGenieConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNote

`func (o *ConfigOpsGenieConfig) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ConfigOpsGenieConfig) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ConfigOpsGenieConfig) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ConfigOpsGenieConfig) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPriority

`func (o *ConfigOpsGenieConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConfigOpsGenieConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConfigOpsGenieConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConfigOpsGenieConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResponders

`func (o *ConfigOpsGenieConfig) GetResponders() []ConfigOpsGenieConfigResponder`

GetResponders returns the Responders field if non-nil, zero value otherwise.

### GetRespondersOk

`func (o *ConfigOpsGenieConfig) GetRespondersOk() (*[]ConfigOpsGenieConfigResponder, bool)`

GetRespondersOk returns a tuple with the Responders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponders

`func (o *ConfigOpsGenieConfig) SetResponders(v []ConfigOpsGenieConfigResponder)`

SetResponders sets Responders field to given value.

### HasResponders

`func (o *ConfigOpsGenieConfig) HasResponders() bool`

HasResponders returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigOpsGenieConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigOpsGenieConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigOpsGenieConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigOpsGenieConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSource

`func (o *ConfigOpsGenieConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConfigOpsGenieConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConfigOpsGenieConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConfigOpsGenieConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *ConfigOpsGenieConfig) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigOpsGenieConfig) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigOpsGenieConfig) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigOpsGenieConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdateAlerts

`func (o *ConfigOpsGenieConfig) GetUpdateAlerts() bool`

GetUpdateAlerts returns the UpdateAlerts field if non-nil, zero value otherwise.

### GetUpdateAlertsOk

`func (o *ConfigOpsGenieConfig) GetUpdateAlertsOk() (*bool, bool)`

GetUpdateAlertsOk returns a tuple with the UpdateAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAlerts

`func (o *ConfigOpsGenieConfig) SetUpdateAlerts(v bool)`

SetUpdateAlerts sets UpdateAlerts field to given value.

### HasUpdateAlerts

`func (o *ConfigOpsGenieConfig) HasUpdateAlerts() bool`

HasUpdateAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


