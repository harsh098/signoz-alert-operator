# RuletypesRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** |  | 
**AlertType** | [**RuletypesAlertType**](RuletypesAlertType.md) |  | 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Condition** | [**RuletypesRuleCondition**](RuletypesRuleCondition.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**EvalWindow** | Pointer to **string** |  | [optional] 
**Evaluation** | Pointer to [**RuletypesEvaluationEnvelope**](RuletypesEvaluationEnvelope.md) |  | [optional] 
**Frequency** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**NotificationSettings** | Pointer to [**RuletypesNotificationSettings**](RuletypesNotificationSettings.md) |  | [optional] 
**PreferredChannels** | Pointer to **[]string** |  | [optional] 
**RuleType** | [**RuletypesRuleType**](RuletypesRuleType.md) |  | 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**State** | [**RuletypesAlertState**](RuletypesAlertState.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewRuletypesRule

`func NewRuletypesRule(alert string, alertType RuletypesAlertType, condition RuletypesRuleCondition, id string, ruleType RuletypesRuleType, state RuletypesAlertState, ) *RuletypesRule`

NewRuletypesRule instantiates a new RuletypesRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesRuleWithDefaults

`func NewRuletypesRuleWithDefaults() *RuletypesRule`

NewRuletypesRuleWithDefaults instantiates a new RuletypesRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *RuletypesRule) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *RuletypesRule) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *RuletypesRule) SetAlert(v string)`

SetAlert sets Alert field to given value.


### GetAlertType

`func (o *RuletypesRule) GetAlertType() RuletypesAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *RuletypesRule) GetAlertTypeOk() (*RuletypesAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *RuletypesRule) SetAlertType(v RuletypesAlertType)`

SetAlertType sets AlertType field to given value.


### GetAnnotations

`func (o *RuletypesRule) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RuletypesRule) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RuletypesRule) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RuletypesRule) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCondition

`func (o *RuletypesRule) GetCondition() RuletypesRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RuletypesRule) GetConditionOk() (*RuletypesRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RuletypesRule) SetCondition(v RuletypesRuleCondition)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *RuletypesRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuletypesRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuletypesRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RuletypesRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RuletypesRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RuletypesRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RuletypesRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RuletypesRule) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *RuletypesRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuletypesRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuletypesRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuletypesRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *RuletypesRule) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuletypesRule) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuletypesRule) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuletypesRule) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEvalWindow

`func (o *RuletypesRule) GetEvalWindow() string`

GetEvalWindow returns the EvalWindow field if non-nil, zero value otherwise.

### GetEvalWindowOk

`func (o *RuletypesRule) GetEvalWindowOk() (*string, bool)`

GetEvalWindowOk returns a tuple with the EvalWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalWindow

`func (o *RuletypesRule) SetEvalWindow(v string)`

SetEvalWindow sets EvalWindow field to given value.

### HasEvalWindow

`func (o *RuletypesRule) HasEvalWindow() bool`

HasEvalWindow returns a boolean if a field has been set.

### GetEvaluation

`func (o *RuletypesRule) GetEvaluation() RuletypesEvaluationEnvelope`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *RuletypesRule) GetEvaluationOk() (*RuletypesEvaluationEnvelope, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *RuletypesRule) SetEvaluation(v RuletypesEvaluationEnvelope)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *RuletypesRule) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetFrequency

`func (o *RuletypesRule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RuletypesRule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RuletypesRule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RuletypesRule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetId

`func (o *RuletypesRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuletypesRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuletypesRule) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *RuletypesRule) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RuletypesRule) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RuletypesRule) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RuletypesRule) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNotificationSettings

`func (o *RuletypesRule) GetNotificationSettings() RuletypesNotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *RuletypesRule) GetNotificationSettingsOk() (*RuletypesNotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *RuletypesRule) SetNotificationSettings(v RuletypesNotificationSettings)`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *RuletypesRule) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetPreferredChannels

`func (o *RuletypesRule) GetPreferredChannels() []string`

GetPreferredChannels returns the PreferredChannels field if non-nil, zero value otherwise.

### GetPreferredChannelsOk

`func (o *RuletypesRule) GetPreferredChannelsOk() (*[]string, bool)`

GetPreferredChannelsOk returns a tuple with the PreferredChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredChannels

`func (o *RuletypesRule) SetPreferredChannels(v []string)`

SetPreferredChannels sets PreferredChannels field to given value.

### HasPreferredChannels

`func (o *RuletypesRule) HasPreferredChannels() bool`

HasPreferredChannels returns a boolean if a field has been set.

### GetRuleType

`func (o *RuletypesRule) GetRuleType() RuletypesRuleType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *RuletypesRule) GetRuleTypeOk() (*RuletypesRuleType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *RuletypesRule) SetRuleType(v RuletypesRuleType)`

SetRuleType sets RuleType field to given value.


### GetSchemaVersion

`func (o *RuletypesRule) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RuletypesRule) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RuletypesRule) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RuletypesRule) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *RuletypesRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RuletypesRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RuletypesRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RuletypesRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *RuletypesRule) GetState() RuletypesAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RuletypesRule) GetStateOk() (*RuletypesAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RuletypesRule) SetState(v RuletypesAlertState)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *RuletypesRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuletypesRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuletypesRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RuletypesRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RuletypesRule) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RuletypesRule) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RuletypesRule) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RuletypesRule) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *RuletypesRule) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RuletypesRule) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RuletypesRule) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RuletypesRule) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


