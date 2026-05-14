# ListUsersDeprecated200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TypesDeprecatedUser**](TypesDeprecatedUser.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListUsersDeprecated200Response

`func NewListUsersDeprecated200Response(data []TypesDeprecatedUser, status string, ) *ListUsersDeprecated200Response`

NewListUsersDeprecated200Response instantiates a new ListUsersDeprecated200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersDeprecated200ResponseWithDefaults

`func NewListUsersDeprecated200ResponseWithDefaults() *ListUsersDeprecated200Response`

NewListUsersDeprecated200ResponseWithDefaults instantiates a new ListUsersDeprecated200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListUsersDeprecated200Response) GetData() []TypesDeprecatedUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUsersDeprecated200Response) GetDataOk() (*[]TypesDeprecatedUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUsersDeprecated200Response) SetData(v []TypesDeprecatedUser)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListUsersDeprecated200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListUsersDeprecated200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListUsersDeprecated200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


