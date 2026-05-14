# LlmpricingruletypesLLMPricingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**IsOverride** | **bool** |  | 
**ModelName** | **string** |  | 
**ModelPattern** | **[]string** |  | 
**OrgId** | **string** |  | 
**Pricing** | [**LlmpricingruletypesLLMRulePricing**](LlmpricingruletypesLLMRulePricing.md) |  | 
**Provider** | **string** |  | 
**SourceId** | Pointer to **string** |  | [optional] 
**SyncedAt** | Pointer to **NullableTime** |  | [optional] 
**Unit** | [**LlmpricingruletypesLLMPricingRuleUnit**](LlmpricingruletypesLLMPricingRuleUnit.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewLlmpricingruletypesLLMPricingRule

`func NewLlmpricingruletypesLLMPricingRule(enabled bool, id string, isOverride bool, modelName string, modelPattern []string, orgId string, pricing LlmpricingruletypesLLMRulePricing, provider string, unit LlmpricingruletypesLLMPricingRuleUnit, ) *LlmpricingruletypesLLMPricingRule`

NewLlmpricingruletypesLLMPricingRule instantiates a new LlmpricingruletypesLLMPricingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmpricingruletypesLLMPricingRuleWithDefaults

`func NewLlmpricingruletypesLLMPricingRuleWithDefaults() *LlmpricingruletypesLLMPricingRule`

NewLlmpricingruletypesLLMPricingRuleWithDefaults instantiates a new LlmpricingruletypesLLMPricingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LlmpricingruletypesLLMPricingRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmpricingruletypesLLMPricingRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmpricingruletypesLLMPricingRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LlmpricingruletypesLLMPricingRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *LlmpricingruletypesLLMPricingRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LlmpricingruletypesLLMPricingRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LlmpricingruletypesLLMPricingRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *LlmpricingruletypesLLMPricingRule) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEnabled

`func (o *LlmpricingruletypesLLMPricingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LlmpricingruletypesLLMPricingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LlmpricingruletypesLLMPricingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *LlmpricingruletypesLLMPricingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmpricingruletypesLLMPricingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmpricingruletypesLLMPricingRule) SetId(v string)`

SetId sets Id field to given value.


### GetIsOverride

`func (o *LlmpricingruletypesLLMPricingRule) GetIsOverride() bool`

GetIsOverride returns the IsOverride field if non-nil, zero value otherwise.

### GetIsOverrideOk

`func (o *LlmpricingruletypesLLMPricingRule) GetIsOverrideOk() (*bool, bool)`

GetIsOverrideOk returns a tuple with the IsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverride

`func (o *LlmpricingruletypesLLMPricingRule) SetIsOverride(v bool)`

SetIsOverride sets IsOverride field to given value.


### GetModelName

`func (o *LlmpricingruletypesLLMPricingRule) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *LlmpricingruletypesLLMPricingRule) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *LlmpricingruletypesLLMPricingRule) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetModelPattern

`func (o *LlmpricingruletypesLLMPricingRule) GetModelPattern() []string`

GetModelPattern returns the ModelPattern field if non-nil, zero value otherwise.

### GetModelPatternOk

`func (o *LlmpricingruletypesLLMPricingRule) GetModelPatternOk() (*[]string, bool)`

GetModelPatternOk returns a tuple with the ModelPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPattern

`func (o *LlmpricingruletypesLLMPricingRule) SetModelPattern(v []string)`

SetModelPattern sets ModelPattern field to given value.


### SetModelPatternNil

`func (o *LlmpricingruletypesLLMPricingRule) SetModelPatternNil(b bool)`

 SetModelPatternNil sets the value for ModelPattern to be an explicit nil

### UnsetModelPattern
`func (o *LlmpricingruletypesLLMPricingRule) UnsetModelPattern()`

UnsetModelPattern ensures that no value is present for ModelPattern, not even an explicit nil
### GetOrgId

`func (o *LlmpricingruletypesLLMPricingRule) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *LlmpricingruletypesLLMPricingRule) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *LlmpricingruletypesLLMPricingRule) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPricing

`func (o *LlmpricingruletypesLLMPricingRule) GetPricing() LlmpricingruletypesLLMRulePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *LlmpricingruletypesLLMPricingRule) GetPricingOk() (*LlmpricingruletypesLLMRulePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *LlmpricingruletypesLLMPricingRule) SetPricing(v LlmpricingruletypesLLMRulePricing)`

SetPricing sets Pricing field to given value.


### GetProvider

`func (o *LlmpricingruletypesLLMPricingRule) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LlmpricingruletypesLLMPricingRule) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LlmpricingruletypesLLMPricingRule) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSourceId

`func (o *LlmpricingruletypesLLMPricingRule) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *LlmpricingruletypesLLMPricingRule) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *LlmpricingruletypesLLMPricingRule) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *LlmpricingruletypesLLMPricingRule) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSyncedAt

`func (o *LlmpricingruletypesLLMPricingRule) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *LlmpricingruletypesLLMPricingRule) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *LlmpricingruletypesLLMPricingRule) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *LlmpricingruletypesLLMPricingRule) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *LlmpricingruletypesLLMPricingRule) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *LlmpricingruletypesLLMPricingRule) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetUnit

`func (o *LlmpricingruletypesLLMPricingRule) GetUnit() LlmpricingruletypesLLMPricingRuleUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LlmpricingruletypesLLMPricingRule) GetUnitOk() (*LlmpricingruletypesLLMPricingRuleUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LlmpricingruletypesLLMPricingRule) SetUnit(v LlmpricingruletypesLLMPricingRuleUnit)`

SetUnit sets Unit field to given value.


### GetUpdatedAt

`func (o *LlmpricingruletypesLLMPricingRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LlmpricingruletypesLLMPricingRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LlmpricingruletypesLLMPricingRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LlmpricingruletypesLLMPricingRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *LlmpricingruletypesLLMPricingRule) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *LlmpricingruletypesLLMPricingRule) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *LlmpricingruletypesLLMPricingRule) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *LlmpricingruletypesLLMPricingRule) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


