# GetObjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CoretypesObjectGroup**](CoretypesObjectGroup.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetObjects200Response

`func NewGetObjects200Response(data []CoretypesObjectGroup, status string, ) *GetObjects200Response`

NewGetObjects200Response instantiates a new GetObjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjects200ResponseWithDefaults

`func NewGetObjects200ResponseWithDefaults() *GetObjects200Response`

NewGetObjects200ResponseWithDefaults instantiates a new GetObjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetObjects200Response) GetData() []CoretypesObjectGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetObjects200Response) GetDataOk() (*[]CoretypesObjectGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetObjects200Response) SetData(v []CoretypesObjectGroup)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetObjects200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetObjects200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetObjects200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


