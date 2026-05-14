# GetAllRoutePolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AlertmanagertypesGettableRoutePolicy**](AlertmanagertypesGettableRoutePolicy.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetAllRoutePolicies200Response

`func NewGetAllRoutePolicies200Response(data []AlertmanagertypesGettableRoutePolicy, status string, ) *GetAllRoutePolicies200Response`

NewGetAllRoutePolicies200Response instantiates a new GetAllRoutePolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllRoutePolicies200ResponseWithDefaults

`func NewGetAllRoutePolicies200ResponseWithDefaults() *GetAllRoutePolicies200Response`

NewGetAllRoutePolicies200ResponseWithDefaults instantiates a new GetAllRoutePolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllRoutePolicies200Response) GetData() []AlertmanagertypesGettableRoutePolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllRoutePolicies200Response) GetDataOk() (*[]AlertmanagertypesGettableRoutePolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllRoutePolicies200Response) SetData(v []AlertmanagertypesGettableRoutePolicy)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetAllRoutePolicies200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllRoutePolicies200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllRoutePolicies200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


