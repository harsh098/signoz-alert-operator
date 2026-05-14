# GetRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AuthtypesRole**](AuthtypesRole.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetRole200Response

`func NewGetRole200Response(data AuthtypesRole, status string, ) *GetRole200Response`

NewGetRole200Response instantiates a new GetRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRole200ResponseWithDefaults

`func NewGetRole200ResponseWithDefaults() *GetRole200Response`

NewGetRole200ResponseWithDefaults instantiates a new GetRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRole200Response) GetData() AuthtypesRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRole200Response) GetDataOk() (*AuthtypesRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRole200Response) SetData(v AuthtypesRole)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetRole200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRole200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRole200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


