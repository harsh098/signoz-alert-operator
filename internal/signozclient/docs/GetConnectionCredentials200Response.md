# GetConnectionCredentials200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CloudintegrationtypesCredentials**](CloudintegrationtypesCredentials.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetConnectionCredentials200Response

`func NewGetConnectionCredentials200Response(data CloudintegrationtypesCredentials, status string, ) *GetConnectionCredentials200Response`

NewGetConnectionCredentials200Response instantiates a new GetConnectionCredentials200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionCredentials200ResponseWithDefaults

`func NewGetConnectionCredentials200ResponseWithDefaults() *GetConnectionCredentials200Response`

NewGetConnectionCredentials200ResponseWithDefaults instantiates a new GetConnectionCredentials200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetConnectionCredentials200Response) GetData() CloudintegrationtypesCredentials`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetConnectionCredentials200Response) GetDataOk() (*CloudintegrationtypesCredentials, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetConnectionCredentials200Response) SetData(v CloudintegrationtypesCredentials)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetConnectionCredentials200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetConnectionCredentials200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetConnectionCredentials200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


