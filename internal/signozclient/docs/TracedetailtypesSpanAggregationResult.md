# TracedetailtypesSpanAggregationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to [**TracedetailtypesSpanAggregationType**](TracedetailtypesSpanAggregationType.md) |  | [optional] 
**Field** | Pointer to [**TelemetrytypesTelemetryFieldKey**](TelemetrytypesTelemetryFieldKey.md) |  | [optional] 
**Value** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewTracedetailtypesSpanAggregationResult

`func NewTracedetailtypesSpanAggregationResult() *TracedetailtypesSpanAggregationResult`

NewTracedetailtypesSpanAggregationResult instantiates a new TracedetailtypesSpanAggregationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracedetailtypesSpanAggregationResultWithDefaults

`func NewTracedetailtypesSpanAggregationResultWithDefaults() *TracedetailtypesSpanAggregationResult`

NewTracedetailtypesSpanAggregationResultWithDefaults instantiates a new TracedetailtypesSpanAggregationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *TracedetailtypesSpanAggregationResult) GetAggregation() TracedetailtypesSpanAggregationType`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TracedetailtypesSpanAggregationResult) GetAggregationOk() (*TracedetailtypesSpanAggregationType, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TracedetailtypesSpanAggregationResult) SetAggregation(v TracedetailtypesSpanAggregationType)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *TracedetailtypesSpanAggregationResult) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetField

`func (o *TracedetailtypesSpanAggregationResult) GetField() TelemetrytypesTelemetryFieldKey`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TracedetailtypesSpanAggregationResult) GetFieldOk() (*TelemetrytypesTelemetryFieldKey, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TracedetailtypesSpanAggregationResult) SetField(v TelemetrytypesTelemetryFieldKey)`

SetField sets Field field to given value.

### HasField

`func (o *TracedetailtypesSpanAggregationResult) HasField() bool`

HasField returns a boolean if a field has been set.

### GetValue

`func (o *TracedetailtypesSpanAggregationResult) GetValue() map[string]int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TracedetailtypesSpanAggregationResult) GetValueOk() (*map[string]int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TracedetailtypesSpanAggregationResult) SetValue(v map[string]int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *TracedetailtypesSpanAggregationResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TracedetailtypesSpanAggregationResult) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TracedetailtypesSpanAggregationResult) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


