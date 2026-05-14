# MetricsexplorertypesInspectMetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** |  | 
**Filter** | Pointer to [**Querybuildertypesv5Filter**](Querybuildertypesv5Filter.md) |  | [optional] 
**MetricName** | **string** |  | 
**Start** | **int64** |  | 

## Methods

### NewMetricsexplorertypesInspectMetricsRequest

`func NewMetricsexplorertypesInspectMetricsRequest(end int64, metricName string, start int64, ) *MetricsexplorertypesInspectMetricsRequest`

NewMetricsexplorertypesInspectMetricsRequest instantiates a new MetricsexplorertypesInspectMetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesInspectMetricsRequestWithDefaults

`func NewMetricsexplorertypesInspectMetricsRequestWithDefaults() *MetricsexplorertypesInspectMetricsRequest`

NewMetricsexplorertypesInspectMetricsRequestWithDefaults instantiates a new MetricsexplorertypesInspectMetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *MetricsexplorertypesInspectMetricsRequest) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricsexplorertypesInspectMetricsRequest) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricsexplorertypesInspectMetricsRequest) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *MetricsexplorertypesInspectMetricsRequest) GetFilter() Querybuildertypesv5Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricsexplorertypesInspectMetricsRequest) GetFilterOk() (*Querybuildertypesv5Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricsexplorertypesInspectMetricsRequest) SetFilter(v Querybuildertypesv5Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetricsexplorertypesInspectMetricsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricsexplorertypesInspectMetricsRequest) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricsexplorertypesInspectMetricsRequest) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricsexplorertypesInspectMetricsRequest) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetStart

`func (o *MetricsexplorertypesInspectMetricsRequest) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricsexplorertypesInspectMetricsRequest) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricsexplorertypesInspectMetricsRequest) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


