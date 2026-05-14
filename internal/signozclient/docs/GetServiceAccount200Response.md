# GetServiceAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ServiceaccounttypesServiceAccountWithRoles**](ServiceaccounttypesServiceAccountWithRoles.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetServiceAccount200Response

`func NewGetServiceAccount200Response(data ServiceaccounttypesServiceAccountWithRoles, status string, ) *GetServiceAccount200Response`

NewGetServiceAccount200Response instantiates a new GetServiceAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceAccount200ResponseWithDefaults

`func NewGetServiceAccount200ResponseWithDefaults() *GetServiceAccount200Response`

NewGetServiceAccount200ResponseWithDefaults instantiates a new GetServiceAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetServiceAccount200Response) GetData() ServiceaccounttypesServiceAccountWithRoles`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetServiceAccount200Response) GetDataOk() (*ServiceaccounttypesServiceAccountWithRoles, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetServiceAccount200Response) SetData(v ServiceaccounttypesServiceAccountWithRoles)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetServiceAccount200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetServiceAccount200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetServiceAccount200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


