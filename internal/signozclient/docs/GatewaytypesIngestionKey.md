# GatewaytypesIngestionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]GatewaytypesLimit**](GatewaytypesLimit.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewaytypesIngestionKey

`func NewGatewaytypesIngestionKey() *GatewaytypesIngestionKey`

NewGatewaytypesIngestionKey instantiates a new GatewaytypesIngestionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaytypesIngestionKeyWithDefaults

`func NewGatewaytypesIngestionKeyWithDefaults() *GatewaytypesIngestionKey`

NewGatewaytypesIngestionKeyWithDefaults instantiates a new GatewaytypesIngestionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GatewaytypesIngestionKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GatewaytypesIngestionKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GatewaytypesIngestionKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GatewaytypesIngestionKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GatewaytypesIngestionKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GatewaytypesIngestionKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GatewaytypesIngestionKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GatewaytypesIngestionKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *GatewaytypesIngestionKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewaytypesIngestionKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewaytypesIngestionKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewaytypesIngestionKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLimits

`func (o *GatewaytypesIngestionKey) GetLimits() []GatewaytypesLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *GatewaytypesIngestionKey) GetLimitsOk() (*[]GatewaytypesLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *GatewaytypesIngestionKey) SetLimits(v []GatewaytypesLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *GatewaytypesIngestionKey) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *GatewaytypesIngestionKey) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *GatewaytypesIngestionKey) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetName

`func (o *GatewaytypesIngestionKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewaytypesIngestionKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewaytypesIngestionKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewaytypesIngestionKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *GatewaytypesIngestionKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewaytypesIngestionKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewaytypesIngestionKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewaytypesIngestionKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GatewaytypesIngestionKey) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GatewaytypesIngestionKey) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUpdatedAt

`func (o *GatewaytypesIngestionKey) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GatewaytypesIngestionKey) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GatewaytypesIngestionKey) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GatewaytypesIngestionKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *GatewaytypesIngestionKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GatewaytypesIngestionKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GatewaytypesIngestionKey) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GatewaytypesIngestionKey) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *GatewaytypesIngestionKey) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GatewaytypesIngestionKey) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GatewaytypesIngestionKey) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *GatewaytypesIngestionKey) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


