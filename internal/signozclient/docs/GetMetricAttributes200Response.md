# GetMetricAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**MetricsexplorertypesMetricAttributesResponse**](MetricsexplorertypesMetricAttributesResponse.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetMetricAttributes200Response

`func NewGetMetricAttributes200Response(data MetricsexplorertypesMetricAttributesResponse, status string, ) *GetMetricAttributes200Response`

NewGetMetricAttributes200Response instantiates a new GetMetricAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricAttributes200ResponseWithDefaults

`func NewGetMetricAttributes200ResponseWithDefaults() *GetMetricAttributes200Response`

NewGetMetricAttributes200ResponseWithDefaults instantiates a new GetMetricAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMetricAttributes200Response) GetData() MetricsexplorertypesMetricAttributesResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetricAttributes200Response) GetDataOk() (*MetricsexplorertypesMetricAttributesResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetricAttributes200Response) SetData(v MetricsexplorertypesMetricAttributesResponse)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetMetricAttributes200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMetricAttributes200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMetricAttributes200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


