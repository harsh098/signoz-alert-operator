# GatewaytypesLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**GatewaytypesLimitConfig**](GatewaytypesLimitConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**KeyId** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to [**GatewaytypesLimitMetric**](GatewaytypesLimitMetric.md) |  | [optional] 
**Signal** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGatewaytypesLimit

`func NewGatewaytypesLimit() *GatewaytypesLimit`

NewGatewaytypesLimit instantiates a new GatewaytypesLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaytypesLimitWithDefaults

`func NewGatewaytypesLimitWithDefaults() *GatewaytypesLimit`

NewGatewaytypesLimitWithDefaults instantiates a new GatewaytypesLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *GatewaytypesLimit) GetConfig() GatewaytypesLimitConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GatewaytypesLimit) GetConfigOk() (*GatewaytypesLimitConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GatewaytypesLimit) SetConfig(v GatewaytypesLimitConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GatewaytypesLimit) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GatewaytypesLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GatewaytypesLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GatewaytypesLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GatewaytypesLimit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *GatewaytypesLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewaytypesLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewaytypesLimit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewaytypesLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyId

`func (o *GatewaytypesLimit) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GatewaytypesLimit) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GatewaytypesLimit) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *GatewaytypesLimit) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetMetric

`func (o *GatewaytypesLimit) GetMetric() GatewaytypesLimitMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GatewaytypesLimit) GetMetricOk() (*GatewaytypesLimitMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GatewaytypesLimit) SetMetric(v GatewaytypesLimitMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GatewaytypesLimit) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetSignal

`func (o *GatewaytypesLimit) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *GatewaytypesLimit) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *GatewaytypesLimit) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *GatewaytypesLimit) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTags

`func (o *GatewaytypesLimit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewaytypesLimit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewaytypesLimit) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewaytypesLimit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GatewaytypesLimit) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GatewaytypesLimit) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUpdatedAt

`func (o *GatewaytypesLimit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GatewaytypesLimit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GatewaytypesLimit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GatewaytypesLimit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


