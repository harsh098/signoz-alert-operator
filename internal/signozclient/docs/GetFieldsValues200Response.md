# GetFieldsValues200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TelemetrytypesGettableFieldValues**](TelemetrytypesGettableFieldValues.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetFieldsValues200Response

`func NewGetFieldsValues200Response(data TelemetrytypesGettableFieldValues, status string, ) *GetFieldsValues200Response`

NewGetFieldsValues200Response instantiates a new GetFieldsValues200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFieldsValues200ResponseWithDefaults

`func NewGetFieldsValues200ResponseWithDefaults() *GetFieldsValues200Response`

NewGetFieldsValues200ResponseWithDefaults instantiates a new GetFieldsValues200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFieldsValues200Response) GetData() TelemetrytypesGettableFieldValues`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFieldsValues200Response) GetDataOk() (*TelemetrytypesGettableFieldValues, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFieldsValues200Response) SetData(v TelemetrytypesGettableFieldValues)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetFieldsValues200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFieldsValues200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFieldsValues200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


