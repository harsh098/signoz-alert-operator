# ListServiceAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServiceaccounttypesServiceAccount**](ServiceaccounttypesServiceAccount.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListServiceAccounts200Response

`func NewListServiceAccounts200Response(data []ServiceaccounttypesServiceAccount, status string, ) *ListServiceAccounts200Response`

NewListServiceAccounts200Response instantiates a new ListServiceAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceAccounts200ResponseWithDefaults

`func NewListServiceAccounts200ResponseWithDefaults() *ListServiceAccounts200Response`

NewListServiceAccounts200ResponseWithDefaults instantiates a new ListServiceAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListServiceAccounts200Response) GetData() []ServiceaccounttypesServiceAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListServiceAccounts200Response) GetDataOk() (*[]ServiceaccounttypesServiceAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListServiceAccounts200Response) SetData(v []ServiceaccounttypesServiceAccount)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListServiceAccounts200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListServiceAccounts200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListServiceAccounts200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


