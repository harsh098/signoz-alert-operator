# GetAlerts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlertmanagertypesDeprecatedGettableAlert**](AlertmanagertypesDeprecatedGettableAlert.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetAlerts200Response

`func NewGetAlerts200Response(data []AlertmanagertypesDeprecatedGettableAlert, status string, ) *GetAlerts200Response`

NewGetAlerts200Response instantiates a new GetAlerts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlerts200ResponseWithDefaults

`func NewGetAlerts200ResponseWithDefaults() *GetAlerts200Response`

NewGetAlerts200ResponseWithDefaults instantiates a new GetAlerts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAlerts200Response) GetData() []AlertmanagertypesDeprecatedGettableAlert`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAlerts200Response) GetDataOk() (*[]AlertmanagertypesDeprecatedGettableAlert, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAlerts200Response) SetData(v []AlertmanagertypesDeprecatedGettableAlert)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetAlerts200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAlerts200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAlerts200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


