# LlmpricingruletypesUpdatableLLMPricingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsOverride** | Pointer to **NullableBool** |  | [optional] 
**ModelName** | **string** |  | 
**ModelPattern** | **[]string** |  | 
**Pricing** | [**LlmpricingruletypesLLMRulePricing**](LlmpricingruletypesLLMRulePricing.md) |  | 
**Provider** | **string** |  | 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**Unit** | [**LlmpricingruletypesLLMPricingRuleUnit**](LlmpricingruletypesLLMPricingRuleUnit.md) |  | 

## Methods

### NewLlmpricingruletypesUpdatableLLMPricingRule

`func NewLlmpricingruletypesUpdatableLLMPricingRule(enabled bool, modelName string, modelPattern []string, pricing LlmpricingruletypesLLMRulePricing, provider string, unit LlmpricingruletypesLLMPricingRuleUnit, ) *LlmpricingruletypesUpdatableLLMPricingRule`

NewLlmpricingruletypesUpdatableLLMPricingRule instantiates a new LlmpricingruletypesUpdatableLLMPricingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmpricingruletypesUpdatableLLMPricingRuleWithDefaults

`func NewLlmpricingruletypesUpdatableLLMPricingRuleWithDefaults() *LlmpricingruletypesUpdatableLLMPricingRule`

NewLlmpricingruletypesUpdatableLLMPricingRuleWithDefaults instantiates a new LlmpricingruletypesUpdatableLLMPricingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LlmpricingruletypesUpdatableLLMPricingRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsOverride

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetIsOverride() bool`

GetIsOverride returns the IsOverride field if non-nil, zero value otherwise.

### GetIsOverrideOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetIsOverrideOk() (*bool, bool)`

GetIsOverrideOk returns a tuple with the IsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverride

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetIsOverride(v bool)`

SetIsOverride sets IsOverride field to given value.

### HasIsOverride

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) HasIsOverride() bool`

HasIsOverride returns a boolean if a field has been set.

### SetIsOverrideNil

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetIsOverrideNil(b bool)`

 SetIsOverrideNil sets the value for IsOverride to be an explicit nil

### UnsetIsOverride
`func (o *LlmpricingruletypesUpdatableLLMPricingRule) UnsetIsOverride()`

UnsetIsOverride ensures that no value is present for IsOverride, not even an explicit nil
### GetModelName

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelPattern

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetModelPattern() []string`

GetModelPattern returns the ModelPattern field if non-nil, zero value otherwise.

### GetModelPatternOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetModelPatternOk() (*[]string, bool)`

GetModelPatternOk returns a tuple with the ModelPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPattern

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetModelPattern(v []string)`

SetModelPattern sets ModelPattern field to given value.


### SetModelPatternNil

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetModelPatternNil(b bool)`

 SetModelPatternNil sets the value for ModelPattern to be an explicit nil

### UnsetModelPattern
`func (o *LlmpricingruletypesUpdatableLLMPricingRule) UnsetModelPattern()`

UnsetModelPattern ensures that no value is present for ModelPattern, not even an explicit nil
### GetPricing

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetPricing() LlmpricingruletypesLLMRulePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetPricingOk() (*LlmpricingruletypesLLMRulePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetPricing(v LlmpricingruletypesLLMRulePricing)`

SetPricing sets Pricing field to given value.


### GetProvider

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSourceId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *LlmpricingruletypesUpdatableLLMPricingRule) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetUnit

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetUnit() LlmpricingruletypesLLMPricingRuleUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) GetUnitOk() (*LlmpricingruletypesLLMPricingRuleUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LlmpricingruletypesUpdatableLLMPricingRule) SetUnit(v LlmpricingruletypesLLMPricingRuleUnit)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


