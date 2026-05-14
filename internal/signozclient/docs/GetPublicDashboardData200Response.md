# GetPublicDashboardData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DashboardtypesGettablePublicDashboardData**](DashboardtypesGettablePublicDashboardData.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetPublicDashboardData200Response

`func NewGetPublicDashboardData200Response(data DashboardtypesGettablePublicDashboardData, status string, ) *GetPublicDashboardData200Response`

NewGetPublicDashboardData200Response instantiates a new GetPublicDashboardData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDashboardData200ResponseWithDefaults

`func NewGetPublicDashboardData200ResponseWithDefaults() *GetPublicDashboardData200Response`

NewGetPublicDashboardData200ResponseWithDefaults instantiates a new GetPublicDashboardData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPublicDashboardData200Response) GetData() DashboardtypesGettablePublicDashboardData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPublicDashboardData200Response) GetDataOk() (*DashboardtypesGettablePublicDashboardData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPublicDashboardData200Response) SetData(v DashboardtypesGettablePublicDashboardData)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetPublicDashboardData200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPublicDashboardData200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPublicDashboardData200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


