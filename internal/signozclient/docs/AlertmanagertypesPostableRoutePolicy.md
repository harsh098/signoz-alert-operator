# AlertmanagertypesPostableRoutePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Expression** | **string** |  | 
**Kind** | Pointer to [**AlertmanagertypesExpressionKind**](AlertmanagertypesExpressionKind.md) |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAlertmanagertypesPostableRoutePolicy

`func NewAlertmanagertypesPostableRoutePolicy(channels []string, expression string, name string, ) *AlertmanagertypesPostableRoutePolicy`

NewAlertmanagertypesPostableRoutePolicy instantiates a new AlertmanagertypesPostableRoutePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertmanagertypesPostableRoutePolicyWithDefaults

`func NewAlertmanagertypesPostableRoutePolicyWithDefaults() *AlertmanagertypesPostableRoutePolicy`

NewAlertmanagertypesPostableRoutePolicyWithDefaults instantiates a new AlertmanagertypesPostableRoutePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *AlertmanagertypesPostableRoutePolicy) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AlertmanagertypesPostableRoutePolicy) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### SetChannelsNil

`func (o *AlertmanagertypesPostableRoutePolicy) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *AlertmanagertypesPostableRoutePolicy) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetDescription

`func (o *AlertmanagertypesPostableRoutePolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertmanagertypesPostableRoutePolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertmanagertypesPostableRoutePolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *AlertmanagertypesPostableRoutePolicy) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *AlertmanagertypesPostableRoutePolicy) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetKind

`func (o *AlertmanagertypesPostableRoutePolicy) GetKind() AlertmanagertypesExpressionKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetKindOk() (*AlertmanagertypesExpressionKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlertmanagertypesPostableRoutePolicy) SetKind(v AlertmanagertypesExpressionKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AlertmanagertypesPostableRoutePolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *AlertmanagertypesPostableRoutePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertmanagertypesPostableRoutePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *AlertmanagertypesPostableRoutePolicy) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AlertmanagertypesPostableRoutePolicy) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AlertmanagertypesPostableRoutePolicy) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AlertmanagertypesPostableRoutePolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AlertmanagertypesPostableRoutePolicy) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AlertmanagertypesPostableRoutePolicy) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


