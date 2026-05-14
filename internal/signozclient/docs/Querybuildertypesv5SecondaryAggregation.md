# Querybuildertypesv5SecondaryAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**GroupBy** | Pointer to [**[]Querybuildertypesv5GroupByKey**](Querybuildertypesv5GroupByKey.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**LimitBy** | Pointer to [**Querybuildertypesv5LimitBy**](Querybuildertypesv5LimitBy.md) |  | [optional] 
**Order** | Pointer to [**[]Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 
**StepInterval** | Pointer to [**Querybuildertypesv5Step**](Querybuildertypesv5Step.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5SecondaryAggregation

`func NewQuerybuildertypesv5SecondaryAggregation() *Querybuildertypesv5SecondaryAggregation`

NewQuerybuildertypesv5SecondaryAggregation instantiates a new Querybuildertypesv5SecondaryAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5SecondaryAggregationWithDefaults

`func NewQuerybuildertypesv5SecondaryAggregationWithDefaults() *Querybuildertypesv5SecondaryAggregation`

NewQuerybuildertypesv5SecondaryAggregationWithDefaults instantiates a new Querybuildertypesv5SecondaryAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Querybuildertypesv5SecondaryAggregation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Querybuildertypesv5SecondaryAggregation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Querybuildertypesv5SecondaryAggregation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetExpression

`func (o *Querybuildertypesv5SecondaryAggregation) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Querybuildertypesv5SecondaryAggregation) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Querybuildertypesv5SecondaryAggregation) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetGroupBy

`func (o *Querybuildertypesv5SecondaryAggregation) GetGroupBy() []Querybuildertypesv5GroupByKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetGroupByOk() (*[]Querybuildertypesv5GroupByKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Querybuildertypesv5SecondaryAggregation) SetGroupBy(v []Querybuildertypesv5GroupByKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Querybuildertypesv5SecondaryAggregation) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *Querybuildertypesv5SecondaryAggregation) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Querybuildertypesv5SecondaryAggregation) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Querybuildertypesv5SecondaryAggregation) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitBy

`func (o *Querybuildertypesv5SecondaryAggregation) GetLimitBy() Querybuildertypesv5LimitBy`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetLimitByOk() (*Querybuildertypesv5LimitBy, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *Querybuildertypesv5SecondaryAggregation) SetLimitBy(v Querybuildertypesv5LimitBy)`

SetLimitBy sets LimitBy field to given value.

### HasLimitBy

`func (o *Querybuildertypesv5SecondaryAggregation) HasLimitBy() bool`

HasLimitBy returns a boolean if a field has been set.

### GetOrder

`func (o *Querybuildertypesv5SecondaryAggregation) GetOrder() []Querybuildertypesv5OrderBy`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetOrderOk() (*[]Querybuildertypesv5OrderBy, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Querybuildertypesv5SecondaryAggregation) SetOrder(v []Querybuildertypesv5OrderBy)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Querybuildertypesv5SecondaryAggregation) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetStepInterval

`func (o *Querybuildertypesv5SecondaryAggregation) GetStepInterval() Querybuildertypesv5Step`

GetStepInterval returns the StepInterval field if non-nil, zero value otherwise.

### GetStepIntervalOk

`func (o *Querybuildertypesv5SecondaryAggregation) GetStepIntervalOk() (*Querybuildertypesv5Step, bool)`

GetStepIntervalOk returns a tuple with the StepInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepInterval

`func (o *Querybuildertypesv5SecondaryAggregation) SetStepInterval(v Querybuildertypesv5Step)`

SetStepInterval sets StepInterval field to given value.

### HasStepInterval

`func (o *Querybuildertypesv5SecondaryAggregation) HasStepInterval() bool`

HasStepInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


