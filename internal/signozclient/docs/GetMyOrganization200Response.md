# GetMyOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TypesOrganization**](TypesOrganization.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetMyOrganization200Response

`func NewGetMyOrganization200Response(data TypesOrganization, status string, ) *GetMyOrganization200Response`

NewGetMyOrganization200Response instantiates a new GetMyOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyOrganization200ResponseWithDefaults

`func NewGetMyOrganization200ResponseWithDefaults() *GetMyOrganization200Response`

NewGetMyOrganization200ResponseWithDefaults instantiates a new GetMyOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMyOrganization200Response) GetData() TypesOrganization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMyOrganization200Response) GetDataOk() (*TypesOrganization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMyOrganization200Response) SetData(v TypesOrganization)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetMyOrganization200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMyOrganization200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMyOrganization200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


