# GetSessionContext200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AuthtypesSessionContext**](AuthtypesSessionContext.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetSessionContext200Response

`func NewGetSessionContext200Response(data AuthtypesSessionContext, status string, ) *GetSessionContext200Response`

NewGetSessionContext200Response instantiates a new GetSessionContext200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionContext200ResponseWithDefaults

`func NewGetSessionContext200ResponseWithDefaults() *GetSessionContext200Response`

NewGetSessionContext200ResponseWithDefaults instantiates a new GetSessionContext200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSessionContext200Response) GetData() AuthtypesSessionContext`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSessionContext200Response) GetDataOk() (*AuthtypesSessionContext, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSessionContext200Response) SetData(v AuthtypesSessionContext)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetSessionContext200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSessionContext200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSessionContext200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


