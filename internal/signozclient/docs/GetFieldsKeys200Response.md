# GetFieldsKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TelemetrytypesGettableFieldKeys**](TelemetrytypesGettableFieldKeys.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetFieldsKeys200Response

`func NewGetFieldsKeys200Response(data TelemetrytypesGettableFieldKeys, status string, ) *GetFieldsKeys200Response`

NewGetFieldsKeys200Response instantiates a new GetFieldsKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFieldsKeys200ResponseWithDefaults

`func NewGetFieldsKeys200ResponseWithDefaults() *GetFieldsKeys200Response`

NewGetFieldsKeys200ResponseWithDefaults instantiates a new GetFieldsKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFieldsKeys200Response) GetData() TelemetrytypesGettableFieldKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFieldsKeys200Response) GetDataOk() (*TelemetrytypesGettableFieldKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFieldsKeys200Response) SetData(v TelemetrytypesGettableFieldKeys)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetFieldsKeys200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFieldsKeys200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFieldsKeys200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


