# CreateRoutePolicy201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AlertmanagertypesGettableRoutePolicy**](AlertmanagertypesGettableRoutePolicy.md) |  | 
**Status** | **string** |  | 

## Methods

### NewCreateRoutePolicy201Response

`func NewCreateRoutePolicy201Response(data AlertmanagertypesGettableRoutePolicy, status string, ) *CreateRoutePolicy201Response`

NewCreateRoutePolicy201Response instantiates a new CreateRoutePolicy201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoutePolicy201ResponseWithDefaults

`func NewCreateRoutePolicy201ResponseWithDefaults() *CreateRoutePolicy201Response`

NewCreateRoutePolicy201ResponseWithDefaults instantiates a new CreateRoutePolicy201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateRoutePolicy201Response) GetData() AlertmanagertypesGettableRoutePolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRoutePolicy201Response) GetDataOk() (*AlertmanagertypesGettableRoutePolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRoutePolicy201Response) SetData(v AlertmanagertypesGettableRoutePolicy)`

SetData sets Data field to given value.


### GetStatus

`func (o *CreateRoutePolicy201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRoutePolicy201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRoutePolicy201Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


