# ListChannels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlertmanagertypesChannel**](AlertmanagertypesChannel.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListChannels200Response

`func NewListChannels200Response(data []AlertmanagertypesChannel, status string, ) *ListChannels200Response`

NewListChannels200Response instantiates a new ListChannels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChannels200ResponseWithDefaults

`func NewListChannels200ResponseWithDefaults() *ListChannels200Response`

NewListChannels200ResponseWithDefaults instantiates a new ListChannels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListChannels200Response) GetData() []AlertmanagertypesChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListChannels200Response) GetDataOk() (*[]AlertmanagertypesChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListChannels200Response) SetData(v []AlertmanagertypesChannel)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListChannels200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListChannels200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListChannels200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


