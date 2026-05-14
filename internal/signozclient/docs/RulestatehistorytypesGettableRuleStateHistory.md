# RulestatehistorytypesGettableRuleStateHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **int32** |  | 
**Labels** | [**[]Querybuildertypesv5Label**](Querybuildertypesv5Label.md) |  | 
**OverallState** | [**RuletypesAlertState**](RuletypesAlertState.md) |  | 
**OverallStateChanged** | **bool** |  | 
**RuleId** | **string** |  | 
**RuleName** | **string** |  | 
**State** | [**RuletypesAlertState**](RuletypesAlertState.md) |  | 
**StateChanged** | **bool** |  | 
**UnixMilli** | **int64** |  | 
**Value** | **float64** |  | 

## Methods

### NewRulestatehistorytypesGettableRuleStateHistory

`func NewRulestatehistorytypesGettableRuleStateHistory(fingerprint int32, labels []Querybuildertypesv5Label, overallState RuletypesAlertState, overallStateChanged bool, ruleId string, ruleName string, state RuletypesAlertState, stateChanged bool, unixMilli int64, value float64, ) *RulestatehistorytypesGettableRuleStateHistory`

NewRulestatehistorytypesGettableRuleStateHistory instantiates a new RulestatehistorytypesGettableRuleStateHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulestatehistorytypesGettableRuleStateHistoryWithDefaults

`func NewRulestatehistorytypesGettableRuleStateHistoryWithDefaults() *RulestatehistorytypesGettableRuleStateHistory`

NewRulestatehistorytypesGettableRuleStateHistoryWithDefaults instantiates a new RulestatehistorytypesGettableRuleStateHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetFingerprint() int32`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetFingerprintOk() (*int32, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetFingerprint(v int32)`

SetFingerprint sets Fingerprint field to given value.


### GetLabels

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetLabels() []Querybuildertypesv5Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetLabelsOk() (*[]Querybuildertypesv5Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetLabels(v []Querybuildertypesv5Label)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *RulestatehistorytypesGettableRuleStateHistory) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOverallState

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetOverallState() RuletypesAlertState`

GetOverallState returns the OverallState field if non-nil, zero value otherwise.

### GetOverallStateOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetOverallStateOk() (*RuletypesAlertState, bool)`

GetOverallStateOk returns a tuple with the OverallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallState

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetOverallState(v RuletypesAlertState)`

SetOverallState sets OverallState field to given value.


### GetOverallStateChanged

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetOverallStateChanged() bool`

GetOverallStateChanged returns the OverallStateChanged field if non-nil, zero value otherwise.

### GetOverallStateChangedOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetOverallStateChangedOk() (*bool, bool)`

GetOverallStateChangedOk returns a tuple with the OverallStateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStateChanged

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetOverallStateChanged(v bool)`

SetOverallStateChanged sets OverallStateChanged field to given value.


### GetRuleId

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetRuleName

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.


### GetState

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetState() RuletypesAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetStateOk() (*RuletypesAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetState(v RuletypesAlertState)`

SetState sets State field to given value.


### GetStateChanged

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetStateChanged() bool`

GetStateChanged returns the StateChanged field if non-nil, zero value otherwise.

### GetStateChangedOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetStateChangedOk() (*bool, bool)`

GetStateChangedOk returns a tuple with the StateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanged

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetStateChanged(v bool)`

SetStateChanged sets StateChanged field to given value.


### GetUnixMilli

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetUnixMilli() int64`

GetUnixMilli returns the UnixMilli field if non-nil, zero value otherwise.

### GetUnixMilliOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetUnixMilliOk() (*int64, bool)`

GetUnixMilliOk returns a tuple with the UnixMilli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixMilli

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetUnixMilli(v int64)`

SetUnixMilli sets UnixMilli field to given value.


### GetValue

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RulestatehistorytypesGettableRuleStateHistory) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RulestatehistorytypesGettableRuleStateHistory) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


