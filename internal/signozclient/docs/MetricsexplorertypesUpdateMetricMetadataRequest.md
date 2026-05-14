# MetricsexplorertypesUpdateMetricMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsMonotonic** | **bool** |  | 
**MetricName** | **string** |  | 
**Temporality** | [**MetrictypesTemporality**](MetrictypesTemporality.md) |  | 
**Type** | [**MetrictypesType**](MetrictypesType.md) |  | 
**Unit** | **string** |  | 

## Methods

### NewMetricsexplorertypesUpdateMetricMetadataRequest

`func NewMetricsexplorertypesUpdateMetricMetadataRequest(description string, isMonotonic bool, metricName string, temporality MetrictypesTemporality, type_ MetrictypesType, unit string, ) *MetricsexplorertypesUpdateMetricMetadataRequest`

NewMetricsexplorertypesUpdateMetricMetadataRequest instantiates a new MetricsexplorertypesUpdateMetricMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesUpdateMetricMetadataRequestWithDefaults

`func NewMetricsexplorertypesUpdateMetricMetadataRequestWithDefaults() *MetricsexplorertypesUpdateMetricMetadataRequest`

NewMetricsexplorertypesUpdateMetricMetadataRequestWithDefaults instantiates a new MetricsexplorertypesUpdateMetricMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsMonotonic

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetIsMonotonic() bool`

GetIsMonotonic returns the IsMonotonic field if non-nil, zero value otherwise.

### GetIsMonotonicOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetIsMonotonicOk() (*bool, bool)`

GetIsMonotonicOk returns a tuple with the IsMonotonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonotonic

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetIsMonotonic(v bool)`

SetIsMonotonic sets IsMonotonic field to given value.


### GetMetricName

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetTemporality

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetTemporality() MetrictypesTemporality`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetTemporalityOk() (*MetrictypesTemporality, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetTemporality(v MetrictypesTemporality)`

SetTemporality sets Temporality field to given value.


### GetType

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetType() MetrictypesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetTypeOk() (*MetrictypesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetType(v MetrictypesType)`

SetType sets Type field to given value.


### GetUnit

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsexplorertypesUpdateMetricMetadataRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


