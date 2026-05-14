# Querybuildertypesv5TimeSeriesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**[]Querybuildertypesv5AggregationBucket**](Querybuildertypesv5AggregationBucket.md) |  | [optional] 
**QueryName** | Pointer to **string** |  | [optional] 

## Methods

### NewQuerybuildertypesv5TimeSeriesData

`func NewQuerybuildertypesv5TimeSeriesData() *Querybuildertypesv5TimeSeriesData`

NewQuerybuildertypesv5TimeSeriesData instantiates a new Querybuildertypesv5TimeSeriesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5TimeSeriesDataWithDefaults

`func NewQuerybuildertypesv5TimeSeriesDataWithDefaults() *Querybuildertypesv5TimeSeriesData`

NewQuerybuildertypesv5TimeSeriesDataWithDefaults instantiates a new Querybuildertypesv5TimeSeriesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *Querybuildertypesv5TimeSeriesData) GetAggregations() []Querybuildertypesv5AggregationBucket`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *Querybuildertypesv5TimeSeriesData) GetAggregationsOk() (*[]Querybuildertypesv5AggregationBucket, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *Querybuildertypesv5TimeSeriesData) SetAggregations(v []Querybuildertypesv5AggregationBucket)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *Querybuildertypesv5TimeSeriesData) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### SetAggregationsNil

`func (o *Querybuildertypesv5TimeSeriesData) SetAggregationsNil(b bool)`

 SetAggregationsNil sets the value for Aggregations to be an explicit nil

### UnsetAggregations
`func (o *Querybuildertypesv5TimeSeriesData) UnsetAggregations()`

UnsetAggregations ensures that no value is present for Aggregations, not even an explicit nil
### GetQueryName

`func (o *Querybuildertypesv5TimeSeriesData) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *Querybuildertypesv5TimeSeriesData) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *Querybuildertypesv5TimeSeriesData) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *Querybuildertypesv5TimeSeriesData) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


