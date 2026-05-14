# MetricsexplorertypesStatsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** |  | 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**Limit** | **int32** |  | 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 
**Start** | **int64** |  | 

## Methods

### NewMetricsexplorertypesStatsRequest

`func NewMetricsexplorertypesStatsRequest(end int64, limit int32, start int64, ) *MetricsexplorertypesStatsRequest`

NewMetricsexplorertypesStatsRequest instantiates a new MetricsexplorertypesStatsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesStatsRequestWithDefaults

`func NewMetricsexplorertypesStatsRequestWithDefaults() *MetricsexplorertypesStatsRequest`

NewMetricsexplorertypesStatsRequestWithDefaults instantiates a new MetricsexplorertypesStatsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MetricsexplorertypesStatsRequest) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsexplorertypesStatsRequest) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricsexplorertypesStatsRequest) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *MetricsexplorertypesStatsRequest) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricsexplorertypesStatsRequest) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricsexplorertypesStatsRequest) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetricsexplorertypesStatsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *MetricsexplorertypesStatsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsexplorertypesStatsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsexplorertypesStatsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *MetricsexplorertypesStatsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MetricsexplorertypesStatsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MetricsexplorertypesStatsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MetricsexplorertypesStatsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *MetricsexplorertypesStatsRequest) GetOrderBy() Querybuildertypesv5OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *MetricsexplorertypesStatsRequest) GetOrderByOk() (*Querybuildertypesv5OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *MetricsexplorertypesStatsRequest) SetOrderBy(v Querybuildertypesv5OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *MetricsexplorertypesStatsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *MetricsexplorertypesStatsRequest) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsexplorertypesStatsRequest) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricsexplorertypesStatsRequest) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


