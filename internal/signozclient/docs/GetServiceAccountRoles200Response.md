# GetServiceAccountRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AuthtypesRole**](AuthtypesRole.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetServiceAccountRoles200Response

`func NewGetServiceAccountRoles200Response(data []AuthtypesRole, status string, ) *GetServiceAccountRoles200Response`

NewGetServiceAccountRoles200Response instantiates a new GetServiceAccountRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceAccountRoles200ResponseWithDefaults

`func NewGetServiceAccountRoles200ResponseWithDefaults() *GetServiceAccountRoles200Response`

NewGetServiceAccountRoles200ResponseWithDefaults instantiates a new GetServiceAccountRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetServiceAccountRoles200Response) GetData() []AuthtypesRole`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServiceAccountRoles200Response) GetDataOk() (*[]AuthtypesRole, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServiceAccountRoles200Response) SetData(v []AuthtypesRole)`

SetData sets Data field to given value.


### SetDataNil

`func (o *GetServiceAccountRoles200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetServiceAccountRoles200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetStatus

`func (o *GetServiceAccountRoles200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetServiceAccountRoles200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetServiceAccountRoles200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


