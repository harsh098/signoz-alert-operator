# TracedetailtypesGettableWaterfallTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]TracedetailtypesSpanAggregationResult**](TracedetailtypesSpanAggregationResult.md) |  | [optional] 
**EndTimestampMillis** | Pointer to **int32** |  | [optional] 
**HasMissingSpans** | Pointer to **bool** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**RootServiceEntryPoint** | Pointer to **string** |  | [optional] 
**RootServiceName** | Pointer to **string** |  | [optional] 
**ServiceNameToTotalDurationMap** | Pointer to **map[string]int32** |  | [optional] 
**Spans** | Pointer to [**[]TracedetailtypesWaterfallSpan**](TracedetailtypesWaterfallSpan.md) |  | [optional] 
**StartTimestampMillis** | Pointer to **int32** |  | [optional] 
**TotalErrorSpansCount** | Pointer to **int32** |  | [optional] 
**TotalSpansCount** | Pointer to **int32** |  | [optional] 
**UncollapsedSpans** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTracedetailtypesGettableWaterfallTrace

`func NewTracedetailtypesGettableWaterfallTrace() *TracedetailtypesGettableWaterfallTrace`

NewTracedetailtypesGettableWaterfallTrace instantiates a new TracedetailtypesGettableWaterfallTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracedetailtypesGettableWaterfallTraceWithDefaults

`func NewTracedetailtypesGettableWaterfallTraceWithDefaults() *TracedetailtypesGettableWaterfallTrace`

NewTracedetailtypesGettableWaterfallTraceWithDefaults instantiates a new TracedetailtypesGettableWaterfallTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *TracedetailtypesGettableWaterfallTrace) GetAggregations() []TracedetailtypesSpanAggregationResult`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetAggregationsOk() (*[]TracedetailtypesSpanAggregationResult, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *TracedetailtypesGettableWaterfallTrace) SetAggregations(v []TracedetailtypesSpanAggregationResult)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *TracedetailtypesGettableWaterfallTrace) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### SetAggregationsNil

`func (o *TracedetailtypesGettableWaterfallTrace) SetAggregationsNil(b bool)`

 SetAggregationsNil sets the value for Aggregations to be an explicit nil

### UnsetAggregations
`func (o *TracedetailtypesGettableWaterfallTrace) UnsetAggregations()`

UnsetAggregations ensures that no value is present for Aggregations, not even an explicit nil
### GetEndTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) GetEndTimestampMillis() int32`

GetEndTimestampMillis returns the EndTimestampMillis field if non-nil, zero value otherwise.

### GetEndTimestampMillisOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetEndTimestampMillisOk() (*int32, bool)`

GetEndTimestampMillisOk returns a tuple with the EndTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) SetEndTimestampMillis(v int32)`

SetEndTimestampMillis sets EndTimestampMillis field to given value.

### HasEndTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) HasEndTimestampMillis() bool`

HasEndTimestampMillis returns a boolean if a field has been set.

### GetHasMissingSpans

`func (o *TracedetailtypesGettableWaterfallTrace) GetHasMissingSpans() bool`

GetHasMissingSpans returns the HasMissingSpans field if non-nil, zero value otherwise.

### GetHasMissingSpansOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetHasMissingSpansOk() (*bool, bool)`

GetHasMissingSpansOk returns a tuple with the HasMissingSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMissingSpans

`func (o *TracedetailtypesGettableWaterfallTrace) SetHasMissingSpans(v bool)`

SetHasMissingSpans sets HasMissingSpans field to given value.

### HasHasMissingSpans

`func (o *TracedetailtypesGettableWaterfallTrace) HasHasMissingSpans() bool`

HasHasMissingSpans returns a boolean if a field has been set.

### GetHasMore

`func (o *TracedetailtypesGettableWaterfallTrace) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TracedetailtypesGettableWaterfallTrace) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *TracedetailtypesGettableWaterfallTrace) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetRootServiceEntryPoint

`func (o *TracedetailtypesGettableWaterfallTrace) GetRootServiceEntryPoint() string`

GetRootServiceEntryPoint returns the RootServiceEntryPoint field if non-nil, zero value otherwise.

### GetRootServiceEntryPointOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetRootServiceEntryPointOk() (*string, bool)`

GetRootServiceEntryPointOk returns a tuple with the RootServiceEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootServiceEntryPoint

`func (o *TracedetailtypesGettableWaterfallTrace) SetRootServiceEntryPoint(v string)`

SetRootServiceEntryPoint sets RootServiceEntryPoint field to given value.

### HasRootServiceEntryPoint

`func (o *TracedetailtypesGettableWaterfallTrace) HasRootServiceEntryPoint() bool`

HasRootServiceEntryPoint returns a boolean if a field has been set.

### GetRootServiceName

`func (o *TracedetailtypesGettableWaterfallTrace) GetRootServiceName() string`

GetRootServiceName returns the RootServiceName field if non-nil, zero value otherwise.

### GetRootServiceNameOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetRootServiceNameOk() (*string, bool)`

GetRootServiceNameOk returns a tuple with the RootServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootServiceName

`func (o *TracedetailtypesGettableWaterfallTrace) SetRootServiceName(v string)`

SetRootServiceName sets RootServiceName field to given value.

### HasRootServiceName

`func (o *TracedetailtypesGettableWaterfallTrace) HasRootServiceName() bool`

HasRootServiceName returns a boolean if a field has been set.

### GetServiceNameToTotalDurationMap

`func (o *TracedetailtypesGettableWaterfallTrace) GetServiceNameToTotalDurationMap() map[string]int32`

GetServiceNameToTotalDurationMap returns the ServiceNameToTotalDurationMap field if non-nil, zero value otherwise.

### GetServiceNameToTotalDurationMapOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetServiceNameToTotalDurationMapOk() (*map[string]int32, bool)`

GetServiceNameToTotalDurationMapOk returns a tuple with the ServiceNameToTotalDurationMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNameToTotalDurationMap

`func (o *TracedetailtypesGettableWaterfallTrace) SetServiceNameToTotalDurationMap(v map[string]int32)`

SetServiceNameToTotalDurationMap sets ServiceNameToTotalDurationMap field to given value.

### HasServiceNameToTotalDurationMap

`func (o *TracedetailtypesGettableWaterfallTrace) HasServiceNameToTotalDurationMap() bool`

HasServiceNameToTotalDurationMap returns a boolean if a field has been set.

### SetServiceNameToTotalDurationMapNil

`func (o *TracedetailtypesGettableWaterfallTrace) SetServiceNameToTotalDurationMapNil(b bool)`

 SetServiceNameToTotalDurationMapNil sets the value for ServiceNameToTotalDurationMap to be an explicit nil

### UnsetServiceNameToTotalDurationMap
`func (o *TracedetailtypesGettableWaterfallTrace) UnsetServiceNameToTotalDurationMap()`

UnsetServiceNameToTotalDurationMap ensures that no value is present for ServiceNameToTotalDurationMap, not even an explicit nil
### GetSpans

`func (o *TracedetailtypesGettableWaterfallTrace) GetSpans() []TracedetailtypesWaterfallSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetSpansOk() (*[]TracedetailtypesWaterfallSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *TracedetailtypesGettableWaterfallTrace) SetSpans(v []TracedetailtypesWaterfallSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *TracedetailtypesGettableWaterfallTrace) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### SetSpansNil

`func (o *TracedetailtypesGettableWaterfallTrace) SetSpansNil(b bool)`

 SetSpansNil sets the value for Spans to be an explicit nil

### UnsetSpans
`func (o *TracedetailtypesGettableWaterfallTrace) UnsetSpans()`

UnsetSpans ensures that no value is present for Spans, not even an explicit nil
### GetStartTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) GetStartTimestampMillis() int32`

GetStartTimestampMillis returns the StartTimestampMillis field if non-nil, zero value otherwise.

### GetStartTimestampMillisOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetStartTimestampMillisOk() (*int32, bool)`

GetStartTimestampMillisOk returns a tuple with the StartTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) SetStartTimestampMillis(v int32)`

SetStartTimestampMillis sets StartTimestampMillis field to given value.

### HasStartTimestampMillis

`func (o *TracedetailtypesGettableWaterfallTrace) HasStartTimestampMillis() bool`

HasStartTimestampMillis returns a boolean if a field has been set.

### GetTotalErrorSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) GetTotalErrorSpansCount() int32`

GetTotalErrorSpansCount returns the TotalErrorSpansCount field if non-nil, zero value otherwise.

### GetTotalErrorSpansCountOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetTotalErrorSpansCountOk() (*int32, bool)`

GetTotalErrorSpansCountOk returns a tuple with the TotalErrorSpansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) SetTotalErrorSpansCount(v int32)`

SetTotalErrorSpansCount sets TotalErrorSpansCount field to given value.

### HasTotalErrorSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) HasTotalErrorSpansCount() bool`

HasTotalErrorSpansCount returns a boolean if a field has been set.

### GetTotalSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) GetTotalSpansCount() int32`

GetTotalSpansCount returns the TotalSpansCount field if non-nil, zero value otherwise.

### GetTotalSpansCountOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetTotalSpansCountOk() (*int32, bool)`

GetTotalSpansCountOk returns a tuple with the TotalSpansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) SetTotalSpansCount(v int32)`

SetTotalSpansCount sets TotalSpansCount field to given value.

### HasTotalSpansCount

`func (o *TracedetailtypesGettableWaterfallTrace) HasTotalSpansCount() bool`

HasTotalSpansCount returns a boolean if a field has been set.

### GetUncollapsedSpans

`func (o *TracedetailtypesGettableWaterfallTrace) GetUncollapsedSpans() []string`

GetUncollapsedSpans returns the UncollapsedSpans field if non-nil, zero value otherwise.

### GetUncollapsedSpansOk

`func (o *TracedetailtypesGettableWaterfallTrace) GetUncollapsedSpansOk() (*[]string, bool)`

GetUncollapsedSpansOk returns a tuple with the UncollapsedSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncollapsedSpans

`func (o *TracedetailtypesGettableWaterfallTrace) SetUncollapsedSpans(v []string)`

SetUncollapsedSpans sets UncollapsedSpans field to given value.

### HasUncollapsedSpans

`func (o *TracedetailtypesGettableWaterfallTrace) HasUncollapsedSpans() bool`

HasUncollapsedSpans returns a boolean if a field has been set.

### SetUncollapsedSpansNil

`func (o *TracedetailtypesGettableWaterfallTrace) SetUncollapsedSpansNil(b bool)`

 SetUncollapsedSpansNil sets the value for UncollapsedSpans to be an explicit nil

### UnsetUncollapsedSpans
`func (o *TracedetailtypesGettableWaterfallTrace) UnsetUncollapsedSpans()`

UnsetUncollapsedSpans ensures that no value is present for UncollapsedSpans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


