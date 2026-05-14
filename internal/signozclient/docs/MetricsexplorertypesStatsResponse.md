# MetricsexplorertypesStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]MetricsexplorertypesStat**](MetricsexplorertypesStat.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewMetricsexplorertypesStatsResponse

`func NewMetricsexplorertypesStatsResponse(metrics []MetricsexplorertypesStat, total int32, ) *MetricsexplorertypesStatsResponse`

NewMetricsexplorertypesStatsResponse instantiates a new MetricsexplorertypesStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesStatsResponseWithDefaults

`func NewMetricsexplorertypesStatsResponseWithDefaults() *MetricsexplorertypesStatsResponse`

NewMetricsexplorertypesStatsResponseWithDefaults instantiates a new MetricsexplorertypesStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *MetricsexplorertypesStatsResponse) GetMetrics() []MetricsexplorertypesStat`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsexplorertypesStatsResponse) GetMetricsOk() (*[]MetricsexplorertypesStat, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricsexplorertypesStatsResponse) SetMetrics(v []MetricsexplorertypesStat)`

SetMetrics sets Metrics field to given value.


### SetMetricsNil

`func (o *MetricsexplorertypesStatsResponse) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *MetricsexplorertypesStatsResponse) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetTotal

`func (o *MetricsexplorertypesStatsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetricsexplorertypesStatsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetricsexplorertypesStatsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


