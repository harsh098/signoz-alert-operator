# CreateServiceAccountKey201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ServiceaccounttypesGettableFactorAPIKeyWithKey**](ServiceaccounttypesGettableFactorAPIKeyWithKey.md) |  | 
**Status** | **string** |  | 

## Methods

### NewCreateServiceAccountKey201Response

`func NewCreateServiceAccountKey201Response(data ServiceaccounttypesGettableFactorAPIKeyWithKey, status string, ) *CreateServiceAccountKey201Response`

NewCreateServiceAccountKey201Response instantiates a new CreateServiceAccountKey201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountKey201ResponseWithDefaults

`func NewCreateServiceAccountKey201ResponseWithDefaults() *CreateServiceAccountKey201Response`

NewCreateServiceAccountKey201ResponseWithDefaults instantiates a new CreateServiceAccountKey201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateServiceAccountKey201Response) GetData() ServiceaccounttypesGettableFactorAPIKeyWithKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateServiceAccountKey201Response) GetDataOk() (*ServiceaccounttypesGettableFactorAPIKeyWithKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateServiceAccountKey201Response) SetData(v ServiceaccounttypesGettableFactorAPIKeyWithKey)`

SetData sets Data field to given value.


### GetStatus

`func (o *CreateServiceAccountKey201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateServiceAccountKey201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateServiceAccountKey201Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


