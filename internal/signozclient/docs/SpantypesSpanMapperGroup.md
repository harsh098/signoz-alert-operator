# SpantypesSpanMapperGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **map[string]interface{}** |  | 
**Condition** | [**SpantypesSpanMapperGroupCondition**](SpantypesSpanMapperGroupCondition.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrgId** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSpantypesSpanMapperGroup

`func NewSpantypesSpanMapperGroup(category map[string]interface{}, condition SpantypesSpanMapperGroupCondition, enabled bool, id string, name string, orgId string, ) *SpantypesSpanMapperGroup`

NewSpantypesSpanMapperGroup instantiates a new SpantypesSpanMapperGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpantypesSpanMapperGroupWithDefaults

`func NewSpantypesSpanMapperGroupWithDefaults() *SpantypesSpanMapperGroup`

NewSpantypesSpanMapperGroupWithDefaults instantiates a new SpantypesSpanMapperGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SpantypesSpanMapperGroup) GetCategory() map[string]interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SpantypesSpanMapperGroup) GetCategoryOk() (*map[string]interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SpantypesSpanMapperGroup) SetCategory(v map[string]interface{})`

SetCategory sets Category field to given value.


### GetCondition

`func (o *SpantypesSpanMapperGroup) GetCondition() SpantypesSpanMapperGroupCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SpantypesSpanMapperGroup) GetConditionOk() (*SpantypesSpanMapperGroupCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SpantypesSpanMapperGroup) SetCondition(v SpantypesSpanMapperGroupCondition)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *SpantypesSpanMapperGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpantypesSpanMapperGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpantypesSpanMapperGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpantypesSpanMapperGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SpantypesSpanMapperGroup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpantypesSpanMapperGroup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpantypesSpanMapperGroup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpantypesSpanMapperGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *SpantypesSpanMapperGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpantypesSpanMapperGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpantypesSpanMapperGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *SpantypesSpanMapperGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpantypesSpanMapperGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpantypesSpanMapperGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SpantypesSpanMapperGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpantypesSpanMapperGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpantypesSpanMapperGroup) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *SpantypesSpanMapperGroup) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SpantypesSpanMapperGroup) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SpantypesSpanMapperGroup) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetUpdatedAt

`func (o *SpantypesSpanMapperGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SpantypesSpanMapperGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SpantypesSpanMapperGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SpantypesSpanMapperGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SpantypesSpanMapperGroup) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SpantypesSpanMapperGroup) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SpantypesSpanMapperGroup) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SpantypesSpanMapperGroup) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


