# GetUsersByRoleID200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TypesUser**](TypesUser.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetUsersByRoleID200Response

`func NewGetUsersByRoleID200Response(data []TypesUser, status string, ) *GetUsersByRoleID200Response`

NewGetUsersByRoleID200Response instantiates a new GetUsersByRoleID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersByRoleID200ResponseWithDefaults

`func NewGetUsersByRoleID200ResponseWithDefaults() *GetUsersByRoleID200Response`

NewGetUsersByRoleID200ResponseWithDefaults instantiates a new GetUsersByRoleID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUsersByRoleID200Response) GetData() []TypesUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUsersByRoleID200Response) GetDataOk() (*[]TypesUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUsersByRoleID200Response) SetData(v []TypesUser)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetUsersByRoleID200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUsersByRoleID200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUsersByRoleID200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


