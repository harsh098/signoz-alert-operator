# RuletypesRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsentFor** | Pointer to **int32** |  | [optional] 
**AlertOnAbsent** | Pointer to **bool** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] 
**CompositeQuery** | [**RuletypesAlertCompositeQuery**](RuletypesAlertCompositeQuery.md) |  | 
**MatchType** | Pointer to [**RuletypesMatchType**](RuletypesMatchType.md) |  | [optional] 
**Op** | Pointer to [**RuletypesCompareOperator**](RuletypesCompareOperator.md) |  | [optional] 
**RequireMinPoints** | Pointer to **bool** |  | [optional] 
**RequiredNumPoints** | Pointer to **int32** |  | [optional] 
**Seasonality** | Pointer to [**RuletypesSeasonality**](RuletypesSeasonality.md) |  | [optional] 
**SelectedQueryName** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **NullableFloat32** |  | [optional] 
**TargetUnit** | Pointer to **string** |  | [optional] 
**Thresholds** | Pointer to [**RuletypesThresholdBasic**](RuletypesThresholdBasic.md) |  | [optional] 

## Methods

### NewRuletypesRuleCondition

`func NewRuletypesRuleCondition(compositeQuery RuletypesAlertCompositeQuery, ) *RuletypesRuleCondition`

NewRuletypesRuleCondition instantiates a new RuletypesRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesRuleConditionWithDefaults

`func NewRuletypesRuleConditionWithDefaults() *RuletypesRuleCondition`

NewRuletypesRuleConditionWithDefaults instantiates a new RuletypesRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsentFor

`func (o *RuletypesRuleCondition) GetAbsentFor() int32`

GetAbsentFor returns the AbsentFor field if non-nil, zero value otherwise.

### GetAbsentForOk

`func (o *RuletypesRuleCondition) GetAbsentForOk() (*int32, bool)`

GetAbsentForOk returns a tuple with the AbsentFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsentFor

`func (o *RuletypesRuleCondition) SetAbsentFor(v int32)`

SetAbsentFor sets AbsentFor field to given value.

### HasAbsentFor

`func (o *RuletypesRuleCondition) HasAbsentFor() bool`

HasAbsentFor returns a boolean if a field has been set.

### GetAlertOnAbsent

`func (o *RuletypesRuleCondition) GetAlertOnAbsent() bool`

GetAlertOnAbsent returns the AlertOnAbsent field if non-nil, zero value otherwise.

### GetAlertOnAbsentOk

`func (o *RuletypesRuleCondition) GetAlertOnAbsentOk() (*bool, bool)`

GetAlertOnAbsentOk returns a tuple with the AlertOnAbsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnAbsent

`func (o *RuletypesRuleCondition) SetAlertOnAbsent(v bool)`

SetAlertOnAbsent sets AlertOnAbsent field to given value.

### HasAlertOnAbsent

`func (o *RuletypesRuleCondition) HasAlertOnAbsent() bool`

HasAlertOnAbsent returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RuletypesRuleCondition) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RuletypesRuleCondition) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RuletypesRuleCondition) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RuletypesRuleCondition) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCompositeQuery

`func (o *RuletypesRuleCondition) GetCompositeQuery() RuletypesAlertCompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *RuletypesRuleCondition) GetCompositeQueryOk() (*RuletypesAlertCompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *RuletypesRuleCondition) SetCompositeQuery(v RuletypesAlertCompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.


### GetMatchType

`func (o *RuletypesRuleCondition) GetMatchType() RuletypesMatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *RuletypesRuleCondition) GetMatchTypeOk() (*RuletypesMatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *RuletypesRuleCondition) SetMatchType(v RuletypesMatchType)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *RuletypesRuleCondition) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetOp

`func (o *RuletypesRuleCondition) GetOp() RuletypesCompareOperator`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *RuletypesRuleCondition) GetOpOk() (*RuletypesCompareOperator, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *RuletypesRuleCondition) SetOp(v RuletypesCompareOperator)`

SetOp sets Op field to given value.

### HasOp

`func (o *RuletypesRuleCondition) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRequireMinPoints

`func (o *RuletypesRuleCondition) GetRequireMinPoints() bool`

GetRequireMinPoints returns the RequireMinPoints field if non-nil, zero value otherwise.

### GetRequireMinPointsOk

`func (o *RuletypesRuleCondition) GetRequireMinPointsOk() (*bool, bool)`

GetRequireMinPointsOk returns a tuple with the RequireMinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMinPoints

`func (o *RuletypesRuleCondition) SetRequireMinPoints(v bool)`

SetRequireMinPoints sets RequireMinPoints field to given value.

### HasRequireMinPoints

`func (o *RuletypesRuleCondition) HasRequireMinPoints() bool`

HasRequireMinPoints returns a boolean if a field has been set.

### GetRequiredNumPoints

`func (o *RuletypesRuleCondition) GetRequiredNumPoints() int32`

GetRequiredNumPoints returns the RequiredNumPoints field if non-nil, zero value otherwise.

### GetRequiredNumPointsOk

`func (o *RuletypesRuleCondition) GetRequiredNumPointsOk() (*int32, bool)`

GetRequiredNumPointsOk returns a tuple with the RequiredNumPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredNumPoints

`func (o *RuletypesRuleCondition) SetRequiredNumPoints(v int32)`

SetRequiredNumPoints sets RequiredNumPoints field to given value.

### HasRequiredNumPoints

`func (o *RuletypesRuleCondition) HasRequiredNumPoints() bool`

HasRequiredNumPoints returns a boolean if a field has been set.

### GetSeasonality

`func (o *RuletypesRuleCondition) GetSeasonality() RuletypesSeasonality`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *RuletypesRuleCondition) GetSeasonalityOk() (*RuletypesSeasonality, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *RuletypesRuleCondition) SetSeasonality(v RuletypesSeasonality)`

SetSeasonality sets Seasonality field to given value.

### HasSeasonality

`func (o *RuletypesRuleCondition) HasSeasonality() bool`

HasSeasonality returns a boolean if a field has been set.

### GetSelectedQueryName

`func (o *RuletypesRuleCondition) GetSelectedQueryName() string`

GetSelectedQueryName returns the SelectedQueryName field if non-nil, zero value otherwise.

### GetSelectedQueryNameOk

`func (o *RuletypesRuleCondition) GetSelectedQueryNameOk() (*string, bool)`

GetSelectedQueryNameOk returns a tuple with the SelectedQueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedQueryName

`func (o *RuletypesRuleCondition) SetSelectedQueryName(v string)`

SetSelectedQueryName sets SelectedQueryName field to given value.

### HasSelectedQueryName

`func (o *RuletypesRuleCondition) HasSelectedQueryName() bool`

HasSelectedQueryName returns a boolean if a field has been set.

### GetTarget

`func (o *RuletypesRuleCondition) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RuletypesRuleCondition) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RuletypesRuleCondition) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RuletypesRuleCondition) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *RuletypesRuleCondition) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *RuletypesRuleCondition) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTargetUnit

`func (o *RuletypesRuleCondition) GetTargetUnit() string`

GetTargetUnit returns the TargetUnit field if non-nil, zero value otherwise.

### GetTargetUnitOk

`func (o *RuletypesRuleCondition) GetTargetUnitOk() (*string, bool)`

GetTargetUnitOk returns a tuple with the TargetUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUnit

`func (o *RuletypesRuleCondition) SetTargetUnit(v string)`

SetTargetUnit sets TargetUnit field to given value.

### HasTargetUnit

`func (o *RuletypesRuleCondition) HasTargetUnit() bool`

HasTargetUnit returns a boolean if a field has been set.

### GetThresholds

`func (o *RuletypesRuleCondition) GetThresholds() RuletypesThresholdBasic`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *RuletypesRuleCondition) GetThresholdsOk() (*RuletypesThresholdBasic, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *RuletypesRuleCondition) SetThresholds(v RuletypesThresholdBasic)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *RuletypesRuleCondition) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


