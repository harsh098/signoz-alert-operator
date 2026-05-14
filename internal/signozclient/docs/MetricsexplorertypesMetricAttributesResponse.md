# MetricsexplorertypesMetricAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]MetricsexplorertypesMetricAttribute**](MetricsexplorertypesMetricAttribute.md) |  | 
**TotalKeys** | **int64** |  | 

## Methods

### NewMetricsexplorertypesMetricAttributesResponse

`func NewMetricsexplorertypesMetricAttributesResponse(attributes []MetricsexplorertypesMetricAttribute, totalKeys int64, ) *MetricsexplorertypesMetricAttributesResponse`

NewMetricsexplorertypesMetricAttributesResponse instantiates a new MetricsexplorertypesMetricAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsexplorertypesMetricAttributesResponseWithDefaults

`func NewMetricsexplorertypesMetricAttributesResponseWithDefaults() *MetricsexplorertypesMetricAttributesResponse`

NewMetricsexplorertypesMetricAttributesResponseWithDefaults instantiates a new MetricsexplorertypesMetricAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MetricsexplorertypesMetricAttributesResponse) GetAttributes() []MetricsexplorertypesMetricAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricsexplorertypesMetricAttributesResponse) GetAttributesOk() (*[]MetricsexplorertypesMetricAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricsexplorertypesMetricAttributesResponse) SetAttributes(v []MetricsexplorertypesMetricAttribute)`

SetAttributes sets Attributes field to given value.


### SetAttributesNil

`func (o *MetricsexplorertypesMetricAttributesResponse) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *MetricsexplorertypesMetricAttributesResponse) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetTotalKeys

`func (o *MetricsexplorertypesMetricAttributesResponse) GetTotalKeys() int64`

GetTotalKeys returns the TotalKeys field if non-nil, zero value otherwise.

### GetTotalKeysOk

`func (o *MetricsexplorertypesMetricAttributesResponse) GetTotalKeysOk() (*int64, bool)`

GetTotalKeysOk returns a tuple with the TotalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKeys

`func (o *MetricsexplorertypesMetricAttributesResponse) SetTotalKeys(v int64)`

SetTotalKeys sets TotalKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


