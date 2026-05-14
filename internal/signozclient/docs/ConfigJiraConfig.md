# ConfigJiraConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiType** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to [**ConfigJiraFieldConfig**](ConfigJiraFieldConfig.md) |  | [optional] 
**HttpConfig** | Pointer to [**ConfigHTTPClientConfig**](ConfigHTTPClientConfig.md) |  | [optional] 
**IssueType** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ReopenDuration** | Pointer to **int64** |  | [optional] 
**ReopenTransition** | Pointer to **string** |  | [optional] 
**ResolveTransition** | Pointer to **string** |  | [optional] 
**SendResolved** | Pointer to **bool** |  | [optional] 
**Summary** | Pointer to [**ConfigJiraFieldConfig**](ConfigJiraFieldConfig.md) |  | [optional] 
**WontFixResolution** | Pointer to **string** |  | [optional] 

## Methods

### NewConfigJiraConfig

`func NewConfigJiraConfig() *ConfigJiraConfig`

NewConfigJiraConfig instantiates a new ConfigJiraConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigJiraConfigWithDefaults

`func NewConfigJiraConfigWithDefaults() *ConfigJiraConfig`

NewConfigJiraConfigWithDefaults instantiates a new ConfigJiraConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiType

`func (o *ConfigJiraConfig) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *ConfigJiraConfig) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *ConfigJiraConfig) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *ConfigJiraConfig) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetApiUrl

`func (o *ConfigJiraConfig) GetApiUrl() map[string]interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ConfigJiraConfig) GetApiUrlOk() (*map[string]interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ConfigJiraConfig) SetApiUrl(v map[string]interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *ConfigJiraConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### GetCustomFields

`func (o *ConfigJiraConfig) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ConfigJiraConfig) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ConfigJiraConfig) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ConfigJiraConfig) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigJiraConfig) GetDescription() ConfigJiraFieldConfig`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigJiraConfig) GetDescriptionOk() (*ConfigJiraFieldConfig, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigJiraConfig) SetDescription(v ConfigJiraFieldConfig)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigJiraConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHttpConfig

`func (o *ConfigJiraConfig) GetHttpConfig() ConfigHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *ConfigJiraConfig) GetHttpConfigOk() (*ConfigHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *ConfigJiraConfig) SetHttpConfig(v ConfigHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *ConfigJiraConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetIssueType

`func (o *ConfigJiraConfig) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *ConfigJiraConfig) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *ConfigJiraConfig) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *ConfigJiraConfig) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLabels

`func (o *ConfigJiraConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConfigJiraConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConfigJiraConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ConfigJiraConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPriority

`func (o *ConfigJiraConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConfigJiraConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConfigJiraConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConfigJiraConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProject

`func (o *ConfigJiraConfig) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ConfigJiraConfig) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ConfigJiraConfig) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ConfigJiraConfig) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReopenDuration

`func (o *ConfigJiraConfig) GetReopenDuration() int64`

GetReopenDuration returns the ReopenDuration field if non-nil, zero value otherwise.

### GetReopenDurationOk

`func (o *ConfigJiraConfig) GetReopenDurationOk() (*int64, bool)`

GetReopenDurationOk returns a tuple with the ReopenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenDuration

`func (o *ConfigJiraConfig) SetReopenDuration(v int64)`

SetReopenDuration sets ReopenDuration field to given value.

### HasReopenDuration

`func (o *ConfigJiraConfig) HasReopenDuration() bool`

HasReopenDuration returns a boolean if a field has been set.

### GetReopenTransition

`func (o *ConfigJiraConfig) GetReopenTransition() string`

GetReopenTransition returns the ReopenTransition field if non-nil, zero value otherwise.

### GetReopenTransitionOk

`func (o *ConfigJiraConfig) GetReopenTransitionOk() (*string, bool)`

GetReopenTransitionOk returns a tuple with the ReopenTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReopenTransition

`func (o *ConfigJiraConfig) SetReopenTransition(v string)`

SetReopenTransition sets ReopenTransition field to given value.

### HasReopenTransition

`func (o *ConfigJiraConfig) HasReopenTransition() bool`

HasReopenTransition returns a boolean if a field has been set.

### GetResolveTransition

`func (o *ConfigJiraConfig) GetResolveTransition() string`

GetResolveTransition returns the ResolveTransition field if non-nil, zero value otherwise.

### GetResolveTransitionOk

`func (o *ConfigJiraConfig) GetResolveTransitionOk() (*string, bool)`

GetResolveTransitionOk returns a tuple with the ResolveTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTransition

`func (o *ConfigJiraConfig) SetResolveTransition(v string)`

SetResolveTransition sets ResolveTransition field to given value.

### HasResolveTransition

`func (o *ConfigJiraConfig) HasResolveTransition() bool`

HasResolveTransition returns a boolean if a field has been set.

### GetSendResolved

`func (o *ConfigJiraConfig) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *ConfigJiraConfig) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *ConfigJiraConfig) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *ConfigJiraConfig) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetSummary

`func (o *ConfigJiraConfig) GetSummary() ConfigJiraFieldConfig`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConfigJiraConfig) GetSummaryOk() (*ConfigJiraFieldConfig, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConfigJiraConfig) SetSummary(v ConfigJiraFieldConfig)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConfigJiraConfig) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWontFixResolution

`func (o *ConfigJiraConfig) GetWontFixResolution() string`

GetWontFixResolution returns the WontFixResolution field if non-nil, zero value otherwise.

### GetWontFixResolutionOk

`func (o *ConfigJiraConfig) GetWontFixResolutionOk() (*string, bool)`

GetWontFixResolutionOk returns a tuple with the WontFixResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWontFixResolution

`func (o *ConfigJiraConfig) SetWontFixResolution(v string)`

SetWontFixResolution sets WontFixResolution field to given value.

### HasWontFixResolution

`func (o *ConfigJiraConfig) HasWontFixResolution() bool`

HasWontFixResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


