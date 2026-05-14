# AuthzCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AuthtypesGettableTransaction**](AuthtypesGettableTransaction.md) |  | 
**Status** | **string** |  | 

## Methods

### NewAuthzCheck200Response

`func NewAuthzCheck200Response(data []AuthtypesGettableTransaction, status string, ) *AuthzCheck200Response`

NewAuthzCheck200Response instantiates a new AuthzCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzCheck200ResponseWithDefaults

`func NewAuthzCheck200ResponseWithDefaults() *AuthzCheck200Response`

NewAuthzCheck200ResponseWithDefaults instantiates a new AuthzCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuthzCheck200Response) GetData() []AuthtypesGettableTransaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthzCheck200Response) GetDataOk() (*[]AuthtypesGettableTransaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthzCheck200Response) SetData(v []AuthtypesGettableTransaction)`

SetData sets Data field to given value.


### GetStatus

`func (o *AuthzCheck200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthzCheck200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthzCheck200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


