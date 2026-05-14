# FactoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Healthy** | Pointer to **bool** |  | [optional] 
**Services** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewFactoryResponse

`func NewFactoryResponse() *FactoryResponse`

NewFactoryResponse instantiates a new FactoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactoryResponseWithDefaults

`func NewFactoryResponseWithDefaults() *FactoryResponse`

NewFactoryResponseWithDefaults instantiates a new FactoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthy

`func (o *FactoryResponse) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *FactoryResponse) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *FactoryResponse) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *FactoryResponse) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetServices

`func (o *FactoryResponse) GetServices() map[string][]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *FactoryResponse) GetServicesOk() (*map[string][]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *FactoryResponse) SetServices(v map[string][]string)`

SetServices sets Services field to given value.

### HasServices

`func (o *FactoryResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *FactoryResponse) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *FactoryResponse) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


