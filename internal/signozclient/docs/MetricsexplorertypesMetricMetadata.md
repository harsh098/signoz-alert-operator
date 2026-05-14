# MetricsexplorertypesMetricMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsMonotonic** | **bool** |  | 
**Temporality** | [**MetrictypesTemporality**](MetrictypesTemporality.md) |  | 
**Type** | [**MetrictypesType**](MetrictypesType.md) |  | 
**Unit** | **string** |  | 

## Methods

### NewMetricsexplorertypesMetricMetadata

`func NewMetricsexplorertypesMetricMetadata(description string, isMonotonic bool, temporality MetrictypesTemporality, type_ MetrictypesType, unit string, ) *MetricsexplorertypesMetricMetadata`

NewMetricsexplorertypesMetricMetadata instantiates a new MetricsexplorertypesMetricMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesMetricMetadataWithDefaults

`func NewMetricsexplorertypesMetricMetadataWithDefaults() *MetricsexplorertypesMetricMetadata`

NewMetricsexplorertypesMetricMetadataWithDefaults instantiates a new MetricsexplorertypesMetricMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricsexplorertypesMetricMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsexplorertypesMetricMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsexplorertypesMetricMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsMonotonic

`func (o *MetricsexplorertypesMetricMetadata) GetIsMonotonic() bool`

GetIsMonotonic returns the IsMonotonic field if non-nil, zero value otherwise.

### GetIsMonotonicOk

`func (o *MetricsexplorertypesMetricMetadata) GetIsMonotonicOk() (*bool, bool)`

GetIsMonotonicOk returns a tuple with the IsMonotonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonotonic

`func (o *MetricsexplorertypesMetricMetadata) SetIsMonotonic(v bool)`

SetIsMonotonic sets IsMonotonic field to given value.


### GetTemporality

`func (o *MetricsexplorertypesMetricMetadata) GetTemporality() MetrictypesTemporality`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *MetricsexplorertypesMetricMetadata) GetTemporalityOk() (*MetrictypesTemporality, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *MetricsexplorertypesMetricMetadata) SetTemporality(v MetrictypesTemporality)`

SetTemporality sets Temporality field to given value.


### GetType

`func (o *MetricsexplorertypesMetricMetadata) GetType() MetrictypesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricsexplorertypesMetricMetadata) GetTypeOk() (*MetrictypesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricsexplorertypesMetricMetadata) SetType(v MetrictypesType)`

SetType sets Type field to given value.


### GetUnit

`func (o *MetricsexplorertypesMetricMetadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricsexplorertypesMetricMetadata) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricsexplorertypesMetricMetadata) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


