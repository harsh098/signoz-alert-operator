# RuletypesBasicRuleThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** |  | [optional] 
**MatchType** | [**RuletypesMatchType**](RuletypesMatchType.md) |  | 
**Name** | **string** |  | 
**Op** | [**RuletypesCompareOperator**](RuletypesCompareOperator.md) |  | 
**RecoveryTarget** | Pointer to **NullableFloat32** |  | [optional] 
**Target** | **NullableFloat32** |  | 
**TargetUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewRuletypesBasicRuleThreshold

`func NewRuletypesBasicRuleThreshold(matchType RuletypesMatchType, name string, op RuletypesCompareOperator, target NullableFloat32, ) *RuletypesBasicRuleThreshold`

NewRuletypesBasicRuleThreshold instantiates a new RuletypesBasicRuleThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesBasicRuleThresholdWithDefaults

`func NewRuletypesBasicRuleThresholdWithDefaults() *RuletypesBasicRuleThreshold`

NewRuletypesBasicRuleThresholdWithDefaults instantiates a new RuletypesBasicRuleThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *RuletypesBasicRuleThreshold) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *RuletypesBasicRuleThreshold) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *RuletypesBasicRuleThreshold) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *RuletypesBasicRuleThreshold) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *RuletypesBasicRuleThreshold) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *RuletypesBasicRuleThreshold) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetMatchType

`func (o *RuletypesBasicRuleThreshold) GetMatchType() RuletypesMatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *RuletypesBasicRuleThreshold) GetMatchTypeOk() (*RuletypesMatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *RuletypesBasicRuleThreshold) SetMatchType(v RuletypesMatchType)`

SetMatchType sets MatchType field to given value.


### GetName

`func (o *RuletypesBasicRuleThreshold) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuletypesBasicRuleThreshold) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuletypesBasicRuleThreshold) SetName(v string)`

SetName sets Name field to given value.


### GetOp

`func (o *RuletypesBasicRuleThreshold) GetOp() RuletypesCompareOperator`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RuletypesBasicRuleThreshold) GetOpOk() (*RuletypesCompareOperator, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RuletypesBasicRuleThreshold) SetOp(v RuletypesCompareOperator)`

SetOp sets Op field to given value.


### GetRecoveryTarget

`func (o *RuletypesBasicRuleThreshold) GetRecoveryTarget() float32`

GetRecoveryTarget returns the RecoveryTarget field if non-nil, zero value otherwise.

### GetRecoveryTargetOk

`func (o *RuletypesBasicRuleThreshold) GetRecoveryTargetOk() (*float32, bool)`

GetRecoveryTargetOk returns a tuple with the RecoveryTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTarget

`func (o *RuletypesBasicRuleThreshold) SetRecoveryTarget(v float32)`

SetRecoveryTarget sets RecoveryTarget field to given value.

### HasRecoveryTarget

`func (o *RuletypesBasicRuleThreshold) HasRecoveryTarget() bool`

HasRecoveryTarget returns a boolean if a field has been set.

### SetRecoveryTargetNil

`func (o *RuletypesBasicRuleThreshold) SetRecoveryTargetNil(b bool)`

 SetRecoveryTargetNil sets the value for RecoveryTarget to be an explicit nil

### UnsetRecoveryTarget
`func (o *RuletypesBasicRuleThreshold) UnsetRecoveryTarget()`

UnsetRecoveryTarget ensures that no value is present for RecoveryTarget, not even an explicit nil
### GetTarget

`func (o *RuletypesBasicRuleThreshold) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RuletypesBasicRuleThreshold) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RuletypesBasicRuleThreshold) SetTarget(v float32)`

SetTarget sets Target field to given value.


### SetTargetNil

`func (o *RuletypesBasicRuleThreshold) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *RuletypesBasicRuleThreshold) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTargetUnit

`func (o *RuletypesBasicRuleThreshold) GetTargetUnit() string`

GetTargetUnit returns the TargetUnit field if non-nil, zero value otherwise.

### GetTargetUnitOk

`func (o *RuletypesBasicRuleThreshold) GetTargetUnitOk() (*string, bool)`

GetTargetUnitOk returns a tuple with the TargetUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUnit

`func (o *RuletypesBasicRuleThreshold) SetTargetUnit(v string)`

SetTargetUnit sets TargetUnit field to given value.

### HasTargetUnit

`func (o *RuletypesBasicRuleThreshold) HasTargetUnit() bool`

HasTargetUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


