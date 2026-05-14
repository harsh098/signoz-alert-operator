# SpantypesPostableSpanMapperGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **map[string]interface{}** |  | 
**Condition** | [**SpantypesSpanMapperGroupCondition**](SpantypesSpanMapperGroupCondition.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewSpantypesPostableSpanMapperGroup

`func NewSpantypesPostableSpanMapperGroup(category map[string]interface{}, condition SpantypesSpanMapperGroupCondition, name string, ) *SpantypesPostableSpanMapperGroup`

NewSpantypesPostableSpanMapperGroup instantiates a new SpantypesPostableSpanMapperGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpantypesPostableSpanMapperGroupWithDefaults

`func NewSpantypesPostableSpanMapperGroupWithDefaults() *SpantypesPostableSpanMapperGroup`

NewSpantypesPostableSpanMapperGroupWithDefaults instantiates a new SpantypesPostableSpanMapperGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SpantypesPostableSpanMapperGroup) GetCategory() map[string]interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SpantypesPostableSpanMapperGroup) GetCategoryOk() (*map[string]interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SpantypesPostableSpanMapperGroup) SetCategory(v map[string]interface{})`

SetCategory sets Category field to given value.


### GetCondition

`func (o *SpantypesPostableSpanMapperGroup) GetCondition() SpantypesSpanMapperGroupCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SpantypesPostableSpanMapperGroup) GetConditionOk() (*SpantypesSpanMapperGroupCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SpantypesPostableSpanMapperGroup) SetCondition(v SpantypesSpanMapperGroupCondition)`

SetCondition sets Condition field to given value.


### GetEnabled

`func (o *SpantypesPostableSpanMapperGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpantypesPostableSpanMapperGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpantypesPostableSpanMapperGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SpantypesPostableSpanMapperGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *SpantypesPostableSpanMapperGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpantypesPostableSpanMapperGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpantypesPostableSpanMapperGroup) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


