# ListServiceAccountKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServiceaccounttypesGettableFactorAPIKey**](ServiceaccounttypesGettableFactorAPIKey.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListServiceAccountKeys200Response

`func NewListServiceAccountKeys200Response(data []ServiceaccounttypesGettableFactorAPIKey, status string, ) *ListServiceAccountKeys200Response`

NewListServiceAccountKeys200Response instantiates a new ListServiceAccountKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceAccountKeys200ResponseWithDefaults

`func NewListServiceAccountKeys200ResponseWithDefaults() *ListServiceAccountKeys200Response`

NewListServiceAccountKeys200ResponseWithDefaults instantiates a new ListServiceAccountKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListServiceAccountKeys200Response) GetData() []ServiceaccounttypesGettableFactorAPIKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListServiceAccountKeys200Response) GetDataOk() (*[]ServiceaccounttypesGettableFactorAPIKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListServiceAccountKeys200Response) SetData(v []ServiceaccounttypesGettableFactorAPIKey)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListServiceAccountKeys200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListServiceAccountKeys200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListServiceAccountKeys200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


