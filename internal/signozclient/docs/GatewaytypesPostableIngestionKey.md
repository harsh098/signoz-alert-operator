# GatewaytypesPostableIngestionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGatewaytypesPostableIngestionKey

`func NewGatewaytypesPostableIngestionKey(name string, ) *GatewaytypesPostableIngestionKey`

NewGatewaytypesPostableIngestionKey instantiates a new GatewaytypesPostableIngestionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaytypesPostableIngestionKeyWithDefaults

`func NewGatewaytypesPostableIngestionKeyWithDefaults() *GatewaytypesPostableIngestionKey`

NewGatewaytypesPostableIngestionKeyWithDefaults instantiates a new GatewaytypesPostableIngestionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *GatewaytypesPostableIngestionKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GatewaytypesPostableIngestionKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GatewaytypesPostableIngestionKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GatewaytypesPostableIngestionKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *GatewaytypesPostableIngestionKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewaytypesPostableIngestionKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewaytypesPostableIngestionKey) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *GatewaytypesPostableIngestionKey) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewaytypesPostableIngestionKey) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewaytypesPostableIngestionKey) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewaytypesPostableIngestionKey) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GatewaytypesPostableIngestionKey) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GatewaytypesPostableIngestionKey) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


