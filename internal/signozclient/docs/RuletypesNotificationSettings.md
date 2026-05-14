# RuletypesNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBy** | Pointer to **[]string** |  | [optional] 
**NewGroupEvalDelay** | Pointer to **string** |  | [optional] 
**Renotify** | Pointer to [**RuletypesRenotify**](RuletypesRenotify.md) |  | [optional] 
**UsePolicy** | Pointer to **bool** |  | [optional] 

## Methods

### NewRuletypesNotificationSettings

`func NewRuletypesNotificationSettings() *RuletypesNotificationSettings`

NewRuletypesNotificationSettings instantiates a new RuletypesNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesNotificationSettingsWithDefaults

`func NewRuletypesNotificationSettingsWithDefaults() *RuletypesNotificationSettings`

NewRuletypesNotificationSettingsWithDefaults instantiates a new RuletypesNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBy

`func (o *RuletypesNotificationSettings) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *RuletypesNotificationSettings) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *RuletypesNotificationSettings) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *RuletypesNotificationSettings) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNewGroupEvalDelay

`func (o *RuletypesNotificationSettings) GetNewGroupEvalDelay() string`

GetNewGroupEvalDelay returns the NewGroupEvalDelay field if non-nil, zero value otherwise.

### GetNewGroupEvalDelayOk

`func (o *RuletypesNotificationSettings) GetNewGroupEvalDelayOk() (*string, bool)`

GetNewGroupEvalDelayOk returns a tuple with the NewGroupEvalDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewGroupEvalDelay

`func (o *RuletypesNotificationSettings) SetNewGroupEvalDelay(v string)`

SetNewGroupEvalDelay sets NewGroupEvalDelay field to given value.

### HasNewGroupEvalDelay

`func (o *RuletypesNotificationSettings) HasNewGroupEvalDelay() bool`

HasNewGroupEvalDelay returns a boolean if a field has been set.

### GetRenotify

`func (o *RuletypesNotificationSettings) GetRenotify() RuletypesRenotify`

GetRenotify returns the Renotify field if non-nil, zero value otherwise.

### GetRenotifyOk

`func (o *RuletypesNotificationSettings) GetRenotifyOk() (*RuletypesRenotify, bool)`

GetRenotifyOk returns a tuple with the Renotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenotify

`func (o *RuletypesNotificationSettings) SetRenotify(v RuletypesRenotify)`

SetRenotify sets Renotify field to given value.

### HasRenotify

`func (o *RuletypesNotificationSettings) HasRenotify() bool`

HasRenotify returns a boolean if a field has been set.

### GetUsePolicy

`func (o *RuletypesNotificationSettings) GetUsePolicy() bool`

GetUsePolicy returns the UsePolicy field if non-nil, zero value otherwise.

### GetUsePolicyOk

`func (o *RuletypesNotificationSettings) GetUsePolicyOk() (*bool, bool)`

GetUsePolicyOk returns a tuple with the UsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePolicy

`func (o *RuletypesNotificationSettings) SetUsePolicy(v bool)`

SetUsePolicy sets UsePolicy field to given value.

### HasUsePolicy

`func (o *RuletypesNotificationSettings) HasUsePolicy() bool`

HasUsePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


