# GetGlobalConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GlobaltypesConfig**](GlobaltypesConfig.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetGlobalConfig200Response

`func NewGetGlobalConfig200Response(data GlobaltypesConfig, status string, ) *GetGlobalConfig200Response`

NewGetGlobalConfig200Response instantiates a new GetGlobalConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalConfig200ResponseWithDefaults

`func NewGetGlobalConfig200ResponseWithDefaults() *GetGlobalConfig200Response`

NewGetGlobalConfig200ResponseWithDefaults instantiates a new GetGlobalConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetGlobalConfig200Response) GetData() GlobaltypesConfig`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetGlobalConfig200Response) GetDataOk() (*GlobaltypesConfig, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetGlobalConfig200Response) SetData(v GlobaltypesConfig)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetGlobalConfig200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGlobalConfig200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGlobalConfig200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


