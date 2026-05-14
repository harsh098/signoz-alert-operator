# InframonitoringtypesPostablePods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** |  | 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]Querybuildertypesv5GroupByKey**](Querybuildertypesv5GroupByKey.md) |  | [optional] 
**Limit** | **int32** |  | 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**Querybuildertypesv5OrderBy**](Querybuildertypesv5OrderBy.md) |  | [optional] 
**Start** | **int64** |  | 

## Methods

### NewInframonitoringtypesPostablePods

`func NewInframonitoringtypesPostablePods(end int64, limit int32, start int64, ) *InframonitoringtypesPostablePods`

NewInframonitoringtypesPostablePods instantiates a new InframonitoringtypesPostablePods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInframonitoringtypesPostablePodsWithDefaults

`func NewInframonitoringtypesPostablePodsWithDefaults() *InframonitoringtypesPostablePods`

NewInframonitoringtypesPostablePodsWithDefaults instantiates a new InframonitoringtypesPostablePods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *InframonitoringtypesPostablePods) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InframonitoringtypesPostablePods) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InframonitoringtypesPostablePods) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *InframonitoringtypesPostablePods) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InframonitoringtypesPostablePods) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InframonitoringtypesPostablePods) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InframonitoringtypesPostablePods) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *InframonitoringtypesPostablePods) GetGroupBy() []Querybuildertypesv5GroupByKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *InframonitoringtypesPostablePods) GetGroupByOk() (*[]Querybuildertypesv5GroupByKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *InframonitoringtypesPostablePods) SetGroupBy(v []Querybuildertypesv5GroupByKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *InframonitoringtypesPostablePods) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### SetGroupByNil

`func (o *InframonitoringtypesPostablePods) SetGroupByNil(b bool)`

 SetGroupByNil sets the value for GroupBy to be an explicit nil

### UnsetGroupBy
`func (o *InframonitoringtypesPostablePods) UnsetGroupBy()`

UnsetGroupBy ensures that no value is present for GroupBy, not even an explicit nil
### GetLimit

`func (o *InframonitoringtypesPostablePods) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InframonitoringtypesPostablePods) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InframonitoringtypesPostablePods) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *InframonitoringtypesPostablePods) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InframonitoringtypesPostablePods) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InframonitoringtypesPostablePods) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *InframonitoringtypesPostablePods) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *InframonitoringtypesPostablePods) GetOrderBy() Querybuildertypesv5OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *InframonitoringtypesPostablePods) GetOrderByOk() (*Querybuildertypesv5OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *InframonitoringtypesPostablePods) SetOrderBy(v Querybuildertypesv5OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *InframonitoringtypesPostablePods) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *InframonitoringtypesPostablePods) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InframonitoringtypesPostablePods) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InframonitoringtypesPostablePods) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


