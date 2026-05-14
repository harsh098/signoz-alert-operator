# GetHosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ZeustypesGettableHost**](ZeustypesGettableHost.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetHosts200Response

`func NewGetHosts200Response(data ZeustypesGettableHost, status string, ) *GetHosts200Response`

NewGetHosts200Response instantiates a new GetHosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHosts200ResponseWithDefaults

`func NewGetHosts200ResponseWithDefaults() *GetHosts200Response`

NewGetHosts200ResponseWithDefaults instantiates a new GetHosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetHosts200Response) GetData() ZeustypesGettableHost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetHosts200Response) GetDataOk() (*ZeustypesGettableHost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetHosts200Response) SetData(v ZeustypesGettableHost)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetHosts200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHosts200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHosts200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


