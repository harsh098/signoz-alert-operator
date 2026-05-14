# GetMetricsStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**MetricsexplorertypesStatsResponse**](MetricsexplorertypesStatsResponse.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetMetricsStats200Response

`func NewGetMetricsStats200Response(data MetricsexplorertypesStatsResponse, status string, ) *GetMetricsStats200Response`

NewGetMetricsStats200Response instantiates a new GetMetricsStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricsStats200ResponseWithDefaults

`func NewGetMetricsStats200ResponseWithDefaults() *GetMetricsStats200Response`

NewGetMetricsStats200ResponseWithDefaults instantiates a new GetMetricsStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMetricsStats200Response) GetData() MetricsexplorertypesStatsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetricsStats200Response) GetDataOk() (*MetricsexplorertypesStatsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetricsStats200Response) SetData(v MetricsexplorertypesStatsResponse)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetMetricsStats200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMetricsStats200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMetricsStats200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


