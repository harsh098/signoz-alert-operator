# Querybuildertypesv5QueryBuilderTraceOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]Querybuildertypesv5TraceAggregation**](Querybuildertypesv5TraceAggregation.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**Functions** | Pointer to [**[]Querybuildertypesv5Function**](Querybuildertypesv5Function.md) |  | [optional] 
**GroupBy** | Pointer to [**[]Querybuildertypesv5GroupByKey**](Querybuildertypesv5GroupByKey.md) |  | [optional] 
**Having** | Pointer to [**Querybuildertypesv5Having**](Querybuildertypesv5Having.md) |  | [optional] 
**Legend** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to [**[]Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 
**ReturnSpansFrom** | Pointer to **string** |  | [optional] 
**SelectFields** | Pointer to [**[]TelemetrytypesTelemetryFieldKey**](TelemetrytypesTelemetryFieldKey.md) |  | [optional] 
**StepInterval** | Pointer to [**Querybuildertypesv5Step**](Querybuildertypesv5Step.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5QueryBuilderTraceOperator

`func NewQuerybuildertypesv5QueryBuilderTraceOperator() *Querybuildertypesv5QueryBuilderTraceOperator`

NewQuerybuildertypesv5QueryBuilderTraceOperator instantiates a new Querybuildertypesv5QueryBuilderTraceOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5QueryBuilderTraceOperatorWithDefaults

`func NewQuerybuildertypesv5QueryBuilderTraceOperatorWithDefaults() *Querybuildertypesv5QueryBuilderTraceOperator`

NewQuerybuildertypesv5QueryBuilderTraceOperatorWithDefaults instantiates a new Querybuildertypesv5QueryBuilderTraceOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetAggregations() []Querybuildertypesv5TraceAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetAggregationsOk() (*[]Querybuildertypesv5TraceAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetAggregations(v []Querybuildertypesv5TraceAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetCursor

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetDisabled

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExpression

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFilter

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFunctions

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetFunctions() []Querybuildertypesv5Function`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetFunctionsOk() (*[]Querybuildertypesv5Function, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetFunctions(v []Querybuildertypesv5Function)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetGroupBy

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetGroupBy() []Querybuildertypesv5GroupByKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetGroupByOk() (*[]Querybuildertypesv5GroupByKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetGroupBy(v []Querybuildertypesv5GroupByKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetHaving

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetHaving() Querybuildertypesv5Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetHavingOk() (*Querybuildertypesv5Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetHaving(v Querybuildertypesv5Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetLegend

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetLegend() string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetLegendOk() (*string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetLegend(v string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetLimit

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetOrder() []Querybuildertypesv5OrderBy`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetOrderOk() (*[]Querybuildertypesv5OrderBy, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetOrder(v []Querybuildertypesv5OrderBy)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReturnSpansFrom

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetReturnSpansFrom() string`

GetReturnSpansFrom returns the ReturnSpansFrom field if non-nil, zero value otherwise.

### GetReturnSpansFromOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetReturnSpansFromOk() (*string, bool)`

GetReturnSpansFromOk returns a tuple with the ReturnSpansFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSpansFrom

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetReturnSpansFrom(v string)`

SetReturnSpansFrom sets ReturnSpansFrom field to given value.

### HasReturnSpansFrom

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasReturnSpansFrom() bool`

HasReturnSpansFrom returns a boolean if a field has been set.

### GetSelectFields

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetSelectFields() []TelemetrytypesTelemetryFieldKey`

GetSelectFields returns the SelectFields field if non-nil, zero value otherwise.

### GetSelectFieldsOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetSelectFieldsOk() (*[]TelemetrytypesTelemetryFieldKey, bool)`

GetSelectFieldsOk returns a tuple with the SelectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectFields

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetSelectFields(v []TelemetrytypesTelemetryFieldKey)`

SetSelectFields sets SelectFields field to given value.

### HasSelectFields

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasSelectFields() bool`

HasSelectFields returns a boolean if a field has been set.

### GetStepInterval

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetStepInterval() Querybuildertypesv5Step`

GetStepInterval returns the StepInterval field if non-nil, zero value otherwise.

### GetStepIntervalOk

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) GetStepIntervalOk() (*Querybuildertypesv5Step, bool)`

GetStepIntervalOk returns a tuple with the StepInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepInterval

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) SetStepInterval(v Querybuildertypesv5Step)`

SetStepInterval sets StepInterval field to given value.

### HasStepInterval

`func (o *Querybuildertypesv5QueryBuilderTraceOperator) HasStepInterval() bool`

HasStepInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


