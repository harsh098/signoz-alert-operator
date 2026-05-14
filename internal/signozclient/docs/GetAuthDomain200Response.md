# GetAuthDomain200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AuthtypesGettableAuthDomain**](AuthtypesGettableAuthDomain.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetAuthDomain200Response

`func NewGetAuthDomain200Response(data AuthtypesGettableAuthDomain, status string, ) *GetAuthDomain200Response`

NewGetAuthDomain200Response instantiates a new GetAuthDomain200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthDomain200ResponseWithDefaults

`func NewGetAuthDomain200ResponseWithDefaults() *GetAuthDomain200Response`

NewGetAuthDomain200ResponseWithDefaults instantiates a new GetAuthDomain200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAuthDomain200Response) GetData() AuthtypesGettableAuthDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAuthDomain200Response) GetDataOk() (*AuthtypesGettableAuthDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAuthDomain200Response) SetData(v AuthtypesGettableAuthDomain)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetAuthDomain200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAuthDomain200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAuthDomain200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


