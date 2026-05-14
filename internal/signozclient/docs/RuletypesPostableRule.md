# RuletypesPostableRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** |  | 
**AlertType** | [**RuletypesAlertType**](RuletypesAlertType.md) |  | 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Condition** | [**RuletypesRuleCondition**](RuletypesRuleCondition.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**EvalWindow** | Pointer to **string** |  | [optional] 
**Evaluation** | Pointer to [**RuletypesEvaluationEnvelope**](RuletypesEvaluationEnvelope.md) |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**NotificationSettings** | Pointer to [**RuletypesNotificationSettings**](RuletypesNotificationSettings.md) |  | [optional] 
**PreferredChannels** | Pointer to **[]string** |  | [optional] 
**RuleType** | [**RuletypesRuleType**](RuletypesRuleType.md) |  | 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewRuletypesPostableRule

`func NewRuletypesPostableRule(alert string, alertType RuletypesAlertType, condition RuletypesRuleCondition, ruleType RuletypesRuleType, ) *RuletypesPostableRule`

NewRuletypesPostableRule instantiates a new RuletypesPostableRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesPostableRuleWithDefaults

`func NewRuletypesPostableRuleWithDefaults() *RuletypesPostableRule`

NewRuletypesPostableRuleWithDefaults instantiates a new RuletypesPostableRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *RuletypesPostableRule) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *RuletypesPostableRule) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *RuletypesPostableRule) SetAlert(v string)`

SetAlert sets Alert field to given value.


### GetAlertType

`func (o *RuletypesPostableRule) GetAlertType() RuletypesAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *RuletypesPostableRule) GetAlertTypeOk() (*RuletypesAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *RuletypesPostableRule) SetAlertType(v RuletypesAlertType)`

SetAlertType sets AlertType field to given value.


### GetAnnotations

`func (o *RuletypesPostableRule) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RuletypesPostableRule) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RuletypesPostableRule) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RuletypesPostableRule) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCondition

`func (o *RuletypesPostableRule) GetCondition() RuletypesRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RuletypesPostableRule) GetConditionOk() (*RuletypesRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RuletypesPostableRule) SetCondition(v RuletypesRuleCondition)`

SetCondition sets Condition field to given value.


### GetDescription

`func (o *RuletypesPostableRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuletypesPostableRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuletypesPostableRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuletypesPostableRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *RuletypesPostableRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuletypesPostableRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuletypesPostableRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuletypesPostableRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEvalWindow

`func (o *RuletypesPostableRule) GetEvalWindow() string`

GetEvalWindow returns the EvalWindow field if non-nil, zero value otherwise.

### GetEvalWindowOk

`func (o *RuletypesPostableRule) GetEvalWindowOk() (*string, bool)`

GetEvalWindowOk returns a tuple with the EvalWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalWindow

`func (o *RuletypesPostableRule) SetEvalWindow(v string)`

SetEvalWindow sets EvalWindow field to given value.

### HasEvalWindow

`func (o *RuletypesPostableRule) HasEvalWindow() bool`

HasEvalWindow returns a boolean if a field has been set.

### GetEvaluation

`func (o *RuletypesPostableRule) GetEvaluation() RuletypesEvaluationEnvelope`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *RuletypesPostableRule) GetEvaluationOk() (*RuletypesEvaluationEnvelope, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *RuletypesPostableRule) SetEvaluation(v RuletypesEvaluationEnvelope)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *RuletypesPostableRule) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetFrequency

`func (o *RuletypesPostableRule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RuletypesPostableRule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RuletypesPostableRule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RuletypesPostableRule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetLabels

`func (o *RuletypesPostableRule) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RuletypesPostableRule) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RuletypesPostableRule) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RuletypesPostableRule) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNotificationSettings

`func (o *RuletypesPostableRule) GetNotificationSettings() RuletypesNotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *RuletypesPostableRule) GetNotificationSettingsOk() (*RuletypesNotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *RuletypesPostableRule) SetNotificationSettings(v RuletypesNotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *RuletypesPostableRule) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetPreferredChannels

`func (o *RuletypesPostableRule) GetPreferredChannels() []string`

GetPreferredChannels returns the PreferredChannels field if non-nil, zero value otherwise.

### GetPreferredChannelsOk

`func (o *RuletypesPostableRule) GetPreferredChannelsOk() (*[]string, bool)`

GetPreferredChannelsOk returns a tuple with the PreferredChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredChannels

`func (o *RuletypesPostableRule) SetPreferredChannels(v []string)`

SetPreferredChannels sets PreferredChannels field to given value.

### HasPreferredChannels

`func (o *RuletypesPostableRule) HasPreferredChannels() bool`

HasPreferredChannels returns a boolean if a field has been set.

### GetRuleType

`func (o *RuletypesPostableRule) GetRuleType() RuletypesRuleType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *RuletypesPostableRule) GetRuleTypeOk() (*RuletypesRuleType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *RuletypesPostableRule) SetRuleType(v RuletypesRuleType)`

SetRuleType sets RuleType field to given value.


### GetSchemaVersion

`func (o *RuletypesPostableRule) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RuletypesPostableRule) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RuletypesPostableRule) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RuletypesPostableRule) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *RuletypesPostableRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RuletypesPostableRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RuletypesPostableRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RuletypesPostableRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *RuletypesPostableRule) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuletypesPostableRule) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuletypesPostableRule) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RuletypesPostableRule) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


