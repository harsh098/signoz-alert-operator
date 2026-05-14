# MetricsexplorertypesTreemapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** |  | 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**Limit** | **int32** |  | 
**Mode** | [**MetricsexplorertypesTreemapMode**](MetricsexplorertypesTreemapMode.md) |  | 
**Start** | **int64** |  | 

## Methods

### NewMetricsexplorertypesTreemapRequest

`func NewMetricsexplorertypesTreemapRequest(end int64, limit int32, mode MetricsexplorertypesTreemapMode, start int64, ) *MetricsexplorertypesTreemapRequest`

NewMetricsexplorertypesTreemapRequest instantiates a new MetricsexplorertypesTreemapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesTreemapRequestWithDefaults

`func NewMetricsexplorertypesTreemapRequestWithDefaults() *MetricsexplorertypesTreemapRequest`

NewMetricsexplorertypesTreemapRequestWithDefaults instantiates a new MetricsexplorertypesTreemapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MetricsexplorertypesTreemapRequest) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsexplorertypesTreemapRequest) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricsexplorertypesTreemapRequest) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *MetricsexplorertypesTreemapRequest) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricsexplorertypesTreemapRequest) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricsexplorertypesTreemapRequest) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetricsexplorertypesTreemapRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *MetricsexplorertypesTreemapRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsexplorertypesTreemapRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsexplorertypesTreemapRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetMode

`func (o *MetricsexplorertypesTreemapRequest) GetMode() MetricsexplorertypesTreemapMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MetricsexplorertypesTreemapRequest) GetModeOk() (*MetricsexplorertypesTreemapMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MetricsexplorertypesTreemapRequest) SetMode(v MetricsexplorertypesTreemapMode)`

SetMode sets Mode field to given value.


### GetStart

`func (o *MetricsexplorertypesTreemapRequest) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsexplorertypesTreemapRequest) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricsexplorertypesTreemapRequest) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


