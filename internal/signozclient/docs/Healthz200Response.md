# Healthz200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**FactoryResponse**](FactoryResponse.md) |  | 
**Status** | **string** |  | 

## Methods

### NewHealthz200Response

`func NewHealthz200Response(data FactoryResponse, status string, ) *Healthz200Response`

NewHealthz200Response instantiates a new Healthz200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthz200ResponseWithDefaults

`func NewHealthz200ResponseWithDefaults() *Healthz200Response`

NewHealthz200ResponseWithDefaults instantiates a new Healthz200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Healthz200Response) GetData() FactoryResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Healthz200Response) GetDataOk() (*FactoryResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Healthz200Response) SetData(v FactoryResponse)`

SetData sets Data field to given value.


### GetStatus

`func (o *Healthz200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Healthz200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Healthz200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


