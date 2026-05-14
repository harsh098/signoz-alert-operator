# Querybuildertypesv5QueryBuilderFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**Functions** | Pointer to [**[]Querybuildertypesv5Function**](Querybuildertypesv5Function.md) |  | [optional] 
**Having** | Pointer to [**Querybuildertypesv5Having**](Querybuildertypesv5Having.md) |  | [optional] 
**Legend** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**[]Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5QueryBuilderFormula

`func NewQuerybuildertypesv5QueryBuilderFormula() *Querybuildertypesv5QueryBuilderFormula`

NewQuerybuildertypesv5QueryBuilderFormula instantiates a new Querybuildertypesv5QueryBuilderFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5QueryBuilderFormulaWithDefaults

`func NewQuerybuildertypesv5QueryBuilderFormulaWithDefaults() *Querybuildertypesv5QueryBuilderFormula`

NewQuerybuildertypesv5QueryBuilderFormulaWithDefaults instantiates a new Querybuildertypesv5QueryBuilderFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *Querybuildertypesv5QueryBuilderFormula) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Querybuildertypesv5QueryBuilderFormula) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Querybuildertypesv5QueryBuilderFormula) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExpression

`func (o *Querybuildertypesv5QueryBuilderFormula) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Querybuildertypesv5QueryBuilderFormula) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Querybuildertypesv5QueryBuilderFormula) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFunctions

`func (o *Querybuildertypesv5QueryBuilderFormula) GetFunctions() []Querybuildertypesv5Function`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetFunctionsOk() (*[]Querybuildertypesv5Function, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Querybuildertypesv5QueryBuilderFormula) SetFunctions(v []Querybuildertypesv5Function)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *Querybuildertypesv5QueryBuilderFormula) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetHaving

`func (o *Querybuildertypesv5QueryBuilderFormula) GetHaving() Querybuildertypesv5Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetHavingOk() (*Querybuildertypesv5Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *Querybuildertypesv5QueryBuilderFormula) SetHaving(v Querybuildertypesv5Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *Querybuildertypesv5QueryBuilderFormula) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetLegend

`func (o *Querybuildertypesv5QueryBuilderFormula) GetLegend() string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetLegendOk() (*string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *Querybuildertypesv5QueryBuilderFormula) SetLegend(v string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *Querybuildertypesv5QueryBuilderFormula) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetLimit

`func (o *Querybuildertypesv5QueryBuilderFormula) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Querybuildertypesv5QueryBuilderFormula) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Querybuildertypesv5QueryBuilderFormula) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *Querybuildertypesv5QueryBuilderFormula) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Querybuildertypesv5QueryBuilderFormula) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Querybuildertypesv5QueryBuilderFormula) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *Querybuildertypesv5QueryBuilderFormula) GetOrder() []Querybuildertypesv5OrderBy`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Querybuildertypesv5QueryBuilderFormula) GetOrderOk() (*[]Querybuildertypesv5OrderBy, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Querybuildertypesv5QueryBuilderFormula) SetOrder(v []Querybuildertypesv5OrderBy)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Querybuildertypesv5QueryBuilderFormula) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


