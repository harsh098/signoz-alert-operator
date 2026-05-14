# SpantypesSpanMapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SpantypesSpanMapperConfig**](SpantypesSpanMapperConfig.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**FieldContext** | [**SpantypesFieldContext**](SpantypesFieldContext.md) |  | 
**GroupId** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewSpantypesSpanMapper

`func NewSpantypesSpanMapper(config SpantypesSpanMapperConfig, enabled bool, fieldContext SpantypesFieldContext, groupId string, id string, name string, ) *SpantypesSpanMapper`

NewSpantypesSpanMapper instantiates a new SpantypesSpanMapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpantypesSpanMapperWithDefaults

`func NewSpantypesSpanMapperWithDefaults() *SpantypesSpanMapper`

NewSpantypesSpanMapperWithDefaults instantiates a new SpantypesSpanMapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SpantypesSpanMapper) GetConfig() SpantypesSpanMapperConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SpantypesSpanMapper) GetConfigOk() (*SpantypesSpanMapperConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SpantypesSpanMapper) SetConfig(v SpantypesSpanMapperConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *SpantypesSpanMapper) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpantypesSpanMapper) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpantypesSpanMapper) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpantypesSpanMapper) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SpantypesSpanMapper) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpantypesSpanMapper) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpantypesSpanMapper) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpantypesSpanMapper) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *SpantypesSpanMapper) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpantypesSpanMapper) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpantypesSpanMapper) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFieldContext

`func (o *SpantypesSpanMapper) GetFieldContext() SpantypesFieldContext`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *SpantypesSpanMapper) GetFieldContextOk() (*SpantypesFieldContext, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *SpantypesSpanMapper) SetFieldContext(v SpantypesFieldContext)`

SetFieldContext sets FieldContext field to given value.


### GetGroupId

`func (o *SpantypesSpanMapper) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SpantypesSpanMapper) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SpantypesSpanMapper) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *SpantypesSpanMapper) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpantypesSpanMapper) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpantypesSpanMapper) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SpantypesSpanMapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpantypesSpanMapper) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpantypesSpanMapper) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *SpantypesSpanMapper) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SpantypesSpanMapper) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SpantypesSpanMapper) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SpantypesSpanMapper) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SpantypesSpanMapper) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SpantypesSpanMapper) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SpantypesSpanMapper) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SpantypesSpanMapper) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


