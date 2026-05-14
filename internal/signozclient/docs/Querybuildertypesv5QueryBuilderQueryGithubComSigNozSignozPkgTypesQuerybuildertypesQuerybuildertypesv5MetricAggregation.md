# Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]Querybuildertypesv5MetricAggregation**](Querybuildertypesv5MetricAggregation.md) |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**Functions** | Pointer to [**[]Querybuildertypesv5Function**](Querybuildertypesv5Function.md) |  | [optional] 
**GroupBy** | Pointer to [**[]Querybuildertypesv5GroupByKey**](Querybuildertypesv5GroupByKey.md) |  | [optional] 
**Having** | Pointer to [**Querybuildertypesv5Having**](Querybuildertypesv5Having.md) |  | [optional] 
**Legend** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**LimitBy** | Pointer to [**Querybuildertypesv5LimitBy**](Querybuildertypesv5LimitBy.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to [**[]Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 
**SecondaryAggregations** | Pointer to [**[]Querybuildertypesv5SecondaryAggregation**](Querybuildertypesv5SecondaryAggregation.md) |  | [optional] 
**SelectFields** | Pointer to [**[]TelemetrytypesTelemetryFieldKey**](TelemetrytypesTelemetryFieldKey.md) |  | [optional] 
**Signal** | Pointer to [**TelemetrytypesSignal**](TelemetrytypesSignal.md) |  | [optional] 
**Source** | Pointer to [**TelemetrytypesSource**](TelemetrytypesSource.md) |  | [optional] 
**StepInterval** | Pointer to [**Querybuildertypesv5Step**](Querybuildertypesv5Step.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation

`func NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation() *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation`

NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation instantiates a new Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregationWithDefaults

`func NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregationWithDefaults() *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation`

NewQuerybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregationWithDefaults instantiates a new Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetAggregations() []Querybuildertypesv5MetricAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetAggregationsOk() (*[]Querybuildertypesv5MetricAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetAggregations(v []Querybuildertypesv5MetricAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetCursor

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetDisabled

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFilter

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFunctions

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetFunctions() []Querybuildertypesv5Function`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetFunctionsOk() (*[]Querybuildertypesv5Function, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetFunctions(v []Querybuildertypesv5Function)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetGroupBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetGroupBy() []Querybuildertypesv5GroupByKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetGroupByOk() (*[]Querybuildertypesv5GroupByKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetGroupBy(v []Querybuildertypesv5GroupByKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetHaving

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetHaving() Querybuildertypesv5Having`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetHavingOk() (*Querybuildertypesv5Having, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetHaving(v Querybuildertypesv5Having)`

SetHaving sets Having field to given value.

### HasHaving

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### GetLegend

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLegend() string`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLegendOk() (*string, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegend

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetLegend(v string)`

SetLegend sets Legend field to given value.

### HasLegend

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### GetLimit

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLimitBy() Querybuildertypesv5LimitBy`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetLimitByOk() (*Querybuildertypesv5LimitBy, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetLimitBy(v Querybuildertypesv5LimitBy)`

SetLimitBy sets LimitBy field to given value.

### HasLimitBy

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasLimitBy() bool`

HasLimitBy returns a boolean if a field has been set.

### GetName

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetOrder() []Querybuildertypesv5OrderBy`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetOrderOk() (*[]Querybuildertypesv5OrderBy, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetOrder(v []Querybuildertypesv5OrderBy)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSecondaryAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSecondaryAggregations() []Querybuildertypesv5SecondaryAggregation`

GetSecondaryAggregations returns the SecondaryAggregations field if non-nil, zero value otherwise.

### GetSecondaryAggregationsOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSecondaryAggregationsOk() (*[]Querybuildertypesv5SecondaryAggregation, bool)`

GetSecondaryAggregationsOk returns a tuple with the SecondaryAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetSecondaryAggregations(v []Querybuildertypesv5SecondaryAggregation)`

SetSecondaryAggregations sets SecondaryAggregations field to given value.

### HasSecondaryAggregations

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasSecondaryAggregations() bool`

HasSecondaryAggregations returns a boolean if a field has been set.

### GetSelectFields

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSelectFields() []TelemetrytypesTelemetryFieldKey`

GetSelectFields returns the SelectFields field if non-nil, zero value otherwise.

### GetSelectFieldsOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSelectFieldsOk() (*[]TelemetrytypesTelemetryFieldKey, bool)`

GetSelectFieldsOk returns a tuple with the SelectFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectFields

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetSelectFields(v []TelemetrytypesTelemetryFieldKey)`

SetSelectFields sets SelectFields field to given value.

### HasSelectFields

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasSelectFields() bool`

HasSelectFields returns a boolean if a field has been set.

### GetSignal

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSignal() TelemetrytypesSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSignalOk() (*TelemetrytypesSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetSignal(v TelemetrytypesSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSource

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSource() TelemetrytypesSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetSourceOk() (*TelemetrytypesSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetSource(v TelemetrytypesSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStepInterval

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetStepInterval() Querybuildertypesv5Step`

GetStepInterval returns the StepInterval field if non-nil, zero value otherwise.

### GetStepIntervalOk

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) GetStepIntervalOk() (*Querybuildertypesv5Step, bool)`

GetStepIntervalOk returns a tuple with the StepInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepInterval

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) SetStepInterval(v Querybuildertypesv5Step)`

SetStepInterval sets StepInterval field to given value.

### HasStepInterval

`func (o *Querybuildertypesv5QueryBuilderQueryGithubComSigNozSignozPkgTypesQuerybuildertypesQuerybuildertypesv5MetricAggregation) HasStepInterval() bool`

HasStepInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


