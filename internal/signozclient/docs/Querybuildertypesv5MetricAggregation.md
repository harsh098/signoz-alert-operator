# Querybuildertypesv5MetricAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonSpaceAggregationParam** | Pointer to [**MetrictypesComparisonSpaceAggregationParam**](MetrictypesComparisonSpaceAggregationParam.md) |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**ReduceTo** | Pointer to [**Querybuildertypesv5ReduceTo**](Querybuildertypesv5ReduceTo.md) |  | [optional] 
**SpaceAggregation** | Pointer to [**MetrictypesSpaceAggregation**](MetrictypesSpaceAggregation.md) |  | [optional] 
**Temporality** | Pointer to [**MetrictypesTemporality**](MetrictypesTemporality.md) |  | [optional] 
**TimeAggregation** | Pointer to [**MetrictypesTimeAggregation**](MetrictypesTimeAggregation.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5MetricAggregation

`func NewQuerybuildertypesv5MetricAggregation() *Querybuildertypesv5MetricAggregation`

NewQuerybuildertypesv5MetricAggregation instantiates a new Querybuildertypesv5MetricAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5MetricAggregationWithDefaults

`func NewQuerybuildertypesv5MetricAggregationWithDefaults() *Querybuildertypesv5MetricAggregation`

NewQuerybuildertypesv5MetricAggregationWithDefaults instantiates a new Querybuildertypesv5MetricAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonSpaceAggregationParam

`func (o *Querybuildertypesv5MetricAggregation) GetComparisonSpaceAggregationParam() MetrictypesComparisonSpaceAggregationParam`

GetComparisonSpaceAggregationParam returns the ComparisonSpaceAggregationParam field if non-nil, zero value otherwise.

### GetComparisonSpaceAggregationParamOk

`func (o *Querybuildertypesv5MetricAggregation) GetComparisonSpaceAggregationParamOk() (*MetrictypesComparisonSpaceAggregationParam, bool)`

GetComparisonSpaceAggregationParamOk returns a tuple with the ComparisonSpaceAggregationParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonSpaceAggregationParam

`func (o *Querybuildertypesv5MetricAggregation) SetComparisonSpaceAggregationParam(v MetrictypesComparisonSpaceAggregationParam)`

SetComparisonSpaceAggregationParam sets ComparisonSpaceAggregationParam field to given value.

### HasComparisonSpaceAggregationParam

`func (o *Querybuildertypesv5MetricAggregation) HasComparisonSpaceAggregationParam() bool`

HasComparisonSpaceAggregationParam returns a boolean if a field has been set.

### GetMetricName

`func (o *Querybuildertypesv5MetricAggregation) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *Querybuildertypesv5MetricAggregation) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *Querybuildertypesv5MetricAggregation) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *Querybuildertypesv5MetricAggregation) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetReduceTo

`func (o *Querybuildertypesv5MetricAggregation) GetReduceTo() Querybuildertypesv5ReduceTo`

GetReduceTo returns the ReduceTo field if non-nil, zero value otherwise.

### GetReduceToOk

`func (o *Querybuildertypesv5MetricAggregation) GetReduceToOk() (*Querybuildertypesv5ReduceTo, bool)`

GetReduceToOk returns a tuple with the ReduceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceTo

`func (o *Querybuildertypesv5MetricAggregation) SetReduceTo(v Querybuildertypesv5ReduceTo)`

SetReduceTo sets ReduceTo field to given value.

### HasReduceTo

`func (o *Querybuildertypesv5MetricAggregation) HasReduceTo() bool`

HasReduceTo returns a boolean if a field has been set.

### GetSpaceAggregation

`func (o *Querybuildertypesv5MetricAggregation) GetSpaceAggregation() MetrictypesSpaceAggregation`

GetSpaceAggregation returns the SpaceAggregation field if non-nil, zero value otherwise.

### GetSpaceAggregationOk

`func (o *Querybuildertypesv5MetricAggregation) GetSpaceAggregationOk() (*MetrictypesSpaceAggregation, bool)`

GetSpaceAggregationOk returns a tuple with the SpaceAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAggregation

`func (o *Querybuildertypesv5MetricAggregation) SetSpaceAggregation(v MetrictypesSpaceAggregation)`

SetSpaceAggregation sets SpaceAggregation field to given value.

### HasSpaceAggregation

`func (o *Querybuildertypesv5MetricAggregation) HasSpaceAggregation() bool`

HasSpaceAggregation returns a boolean if a field has been set.

### GetTemporality

`func (o *Querybuildertypesv5MetricAggregation) GetTemporality() MetrictypesTemporality`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *Querybuildertypesv5MetricAggregation) GetTemporalityOk() (*MetrictypesTemporality, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *Querybuildertypesv5MetricAggregation) SetTemporality(v MetrictypesTemporality)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *Querybuildertypesv5MetricAggregation) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetTimeAggregation

`func (o *Querybuildertypesv5MetricAggregation) GetTimeAggregation() MetrictypesTimeAggregation`

GetTimeAggregation returns the TimeAggregation field if non-nil, zero value otherwise.

### GetTimeAggregationOk

`func (o *Querybuildertypesv5MetricAggregation) GetTimeAggregationOk() (*MetrictypesTimeAggregation, bool)`

GetTimeAggregationOk returns a tuple with the TimeAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregation

`func (o *Querybuildertypesv5MetricAggregation) SetTimeAggregation(v MetrictypesTimeAggregation)`

SetTimeAggregation sets TimeAggregation field to given value.

### HasTimeAggregation

`func (o *Querybuildertypesv5MetricAggregation) HasTimeAggregation() bool`

HasTimeAggregation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


