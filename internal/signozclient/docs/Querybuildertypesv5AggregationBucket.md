# Querybuildertypesv5AggregationBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**AnomalyScores** | Pointer to [**[]Querybuildertypesv5TimeSeries**](Querybuildertypesv5TimeSeries.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**LowerBoundSeries** | Pointer to [**[]Querybuildertypesv5TimeSeries**](Querybuildertypesv5TimeSeries.md) |  | [optional] 
**Meta** | Pointer to [**Querybuildertypesv5AggregationBucketMeta**](Querybuildertypesv5AggregationBucketMeta.md) |  | [optional] 
**PredictedSeries** | Pointer to [**[]Querybuildertypesv5TimeSeries**](Querybuildertypesv5TimeSeries.md) |  | [optional] 
**Series** | Pointer to [**[]Querybuildertypesv5TimeSeries**](Querybuildertypesv5TimeSeries.md) |  | [optional] 
**UpperBoundSeries** | Pointer to [**[]Querybuildertypesv5TimeSeries**](Querybuildertypesv5TimeSeries.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5AggregationBucket

`func NewQuerybuildertypesv5AggregationBucket() *Querybuildertypesv5AggregationBucket`

NewQuerybuildertypesv5AggregationBucket instantiates a new Querybuildertypesv5AggregationBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5AggregationBucketWithDefaults

`func NewQuerybuildertypesv5AggregationBucketWithDefaults() *Querybuildertypesv5AggregationBucket`

NewQuerybuildertypesv5AggregationBucketWithDefaults instantiates a new Querybuildertypesv5AggregationBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Querybuildertypesv5AggregationBucket) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Querybuildertypesv5AggregationBucket) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Querybuildertypesv5AggregationBucket) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Querybuildertypesv5AggregationBucket) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAnomalyScores

`func (o *Querybuildertypesv5AggregationBucket) GetAnomalyScores() []Querybuildertypesv5TimeSeries`

GetAnomalyScores returns the AnomalyScores field if non-nil, zero value otherwise.

### GetAnomalyScoresOk

`func (o *Querybuildertypesv5AggregationBucket) GetAnomalyScoresOk() (*[]Querybuildertypesv5TimeSeries, bool)`

GetAnomalyScoresOk returns a tuple with the AnomalyScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyScores

`func (o *Querybuildertypesv5AggregationBucket) SetAnomalyScores(v []Querybuildertypesv5TimeSeries)`

SetAnomalyScores sets AnomalyScores field to given value.

### HasAnomalyScores

`func (o *Querybuildertypesv5AggregationBucket) HasAnomalyScores() bool`

HasAnomalyScores returns a boolean if a field has been set.

### GetIndex

`func (o *Querybuildertypesv5AggregationBucket) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Querybuildertypesv5AggregationBucket) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Querybuildertypesv5AggregationBucket) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Querybuildertypesv5AggregationBucket) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLowerBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) GetLowerBoundSeries() []Querybuildertypesv5TimeSeries`

GetLowerBoundSeries returns the LowerBoundSeries field if non-nil, zero value otherwise.

### GetLowerBoundSeriesOk

`func (o *Querybuildertypesv5AggregationBucket) GetLowerBoundSeriesOk() (*[]Querybuildertypesv5TimeSeries, bool)`

GetLowerBoundSeriesOk returns a tuple with the LowerBoundSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) SetLowerBoundSeries(v []Querybuildertypesv5TimeSeries)`

SetLowerBoundSeries sets LowerBoundSeries field to given value.

### HasLowerBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) HasLowerBoundSeries() bool`

HasLowerBoundSeries returns a boolean if a field has been set.

### GetMeta

`func (o *Querybuildertypesv5AggregationBucket) GetMeta() Querybuildertypesv5AggregationBucketMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Querybuildertypesv5AggregationBucket) GetMetaOk() (*Querybuildertypesv5AggregationBucketMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Querybuildertypesv5AggregationBucket) SetMeta(v Querybuildertypesv5AggregationBucketMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Querybuildertypesv5AggregationBucket) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPredictedSeries

`func (o *Querybuildertypesv5AggregationBucket) GetPredictedSeries() []Querybuildertypesv5TimeSeries`

GetPredictedSeries returns the PredictedSeries field if non-nil, zero value otherwise.

### GetPredictedSeriesOk

`func (o *Querybuildertypesv5AggregationBucket) GetPredictedSeriesOk() (*[]Querybuildertypesv5TimeSeries, bool)`

GetPredictedSeriesOk returns a tuple with the PredictedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedSeries

`func (o *Querybuildertypesv5AggregationBucket) SetPredictedSeries(v []Querybuildertypesv5TimeSeries)`

SetPredictedSeries sets PredictedSeries field to given value.

### HasPredictedSeries

`func (o *Querybuildertypesv5AggregationBucket) HasPredictedSeries() bool`

HasPredictedSeries returns a boolean if a field has been set.

### GetSeries

`func (o *Querybuildertypesv5AggregationBucket) GetSeries() []Querybuildertypesv5TimeSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Querybuildertypesv5AggregationBucket) GetSeriesOk() (*[]Querybuildertypesv5TimeSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Querybuildertypesv5AggregationBucket) SetSeries(v []Querybuildertypesv5TimeSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Querybuildertypesv5AggregationBucket) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *Querybuildertypesv5AggregationBucket) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *Querybuildertypesv5AggregationBucket) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetUpperBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) GetUpperBoundSeries() []Querybuildertypesv5TimeSeries`

GetUpperBoundSeries returns the UpperBoundSeries field if non-nil, zero value otherwise.

### GetUpperBoundSeriesOk

`func (o *Querybuildertypesv5AggregationBucket) GetUpperBoundSeriesOk() (*[]Querybuildertypesv5TimeSeries, bool)`

GetUpperBoundSeriesOk returns a tuple with the UpperBoundSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) SetUpperBoundSeries(v []Querybuildertypesv5TimeSeries)`

SetUpperBoundSeries sets UpperBoundSeries field to given value.

### HasUpperBoundSeries

`func (o *Querybuildertypesv5AggregationBucket) HasUpperBoundSeries() bool`

HasUpperBoundSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


