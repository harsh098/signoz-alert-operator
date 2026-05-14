# GetMetricMetadata200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**MetricsexplorertypesMetricMetadata**](MetricsexplorertypesMetricMetadata.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetMetricMetadata200Response

`func NewGetMetricMetadata200Response(data MetricsexplorertypesMetricMetadata, status string, ) *GetMetricMetadata200Response`

NewGetMetricMetadata200Response instantiates a new GetMetricMetadata200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricMetadata200ResponseWithDefaults

`func NewGetMetricMetadata200ResponseWithDefaults() *GetMetricMetadata200Response`

NewGetMetricMetadata200ResponseWithDefaults instantiates a new GetMetricMetadata200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMetricMetadata200Response) GetData() MetricsexplorertypesMetricMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetricMetadata200Response) GetDataOk() (*MetricsexplorertypesMetricMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetricMetadata200Response) SetData(v MetricsexplorertypesMetricMetadata)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetMetricMetadata200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMetricMetadata200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMetricMetadata200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


