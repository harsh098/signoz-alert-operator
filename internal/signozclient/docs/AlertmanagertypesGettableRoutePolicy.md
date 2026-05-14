# AlertmanagertypesGettableRoutePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expression** | **string** |  | 
**Id** | **string** |  | 
**Kind** | Pointer to [**AlertmanagertypesExpressionKind**](AlertmanagertypesExpressionKind.md) |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertmanagertypesGettableRoutePolicy

`func NewAlertmanagertypesGettableRoutePolicy(channels []string, createdAt time.Time, expression string, id string, name string, updatedAt time.Time, ) *AlertmanagertypesGettableRoutePolicy`

NewAlertmanagertypesGettableRoutePolicy instantiates a new AlertmanagertypesGettableRoutePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertmanagertypesGettableRoutePolicyWithDefaults

`func NewAlertmanagertypesGettableRoutePolicyWithDefaults() *AlertmanagertypesGettableRoutePolicy`

NewAlertmanagertypesGettableRoutePolicyWithDefaults instantiates a new AlertmanagertypesGettableRoutePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *AlertmanagertypesGettableRoutePolicy) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AlertmanagertypesGettableRoutePolicy) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### SetChannelsNil

`func (o *AlertmanagertypesGettableRoutePolicy) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *AlertmanagertypesGettableRoutePolicy) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetCreatedAt

`func (o *AlertmanagertypesGettableRoutePolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertmanagertypesGettableRoutePolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AlertmanagertypesGettableRoutePolicy) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AlertmanagertypesGettableRoutePolicy) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDescription

`func (o *AlertmanagertypesGettableRoutePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertmanagertypesGettableRoutePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertmanagertypesGettableRoutePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *AlertmanagertypesGettableRoutePolicy) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *AlertmanagertypesGettableRoutePolicy) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetId

`func (o *AlertmanagertypesGettableRoutePolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertmanagertypesGettableRoutePolicy) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *AlertmanagertypesGettableRoutePolicy) GetKind() AlertmanagertypesExpressionKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetKindOk() (*AlertmanagertypesExpressionKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlertmanagertypesGettableRoutePolicy) SetKind(v AlertmanagertypesExpressionKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AlertmanagertypesGettableRoutePolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *AlertmanagertypesGettableRoutePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertmanagertypesGettableRoutePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *AlertmanagertypesGettableRoutePolicy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AlertmanagertypesGettableRoutePolicy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AlertmanagertypesGettableRoutePolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AlertmanagertypesGettableRoutePolicy) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AlertmanagertypesGettableRoutePolicy) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUpdatedAt

`func (o *AlertmanagertypesGettableRoutePolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AlertmanagertypesGettableRoutePolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AlertmanagertypesGettableRoutePolicy) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AlertmanagertypesGettableRoutePolicy) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *AlertmanagertypesGettableRoutePolicy) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *AlertmanagertypesGettableRoutePolicy) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


