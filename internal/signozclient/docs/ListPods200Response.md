# ListPods200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**InframonitoringtypesPods**](InframonitoringtypesPods.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListPods200Response

`func NewListPods200Response(data InframonitoringtypesPods, status string, ) *ListPods200Response`

NewListPods200Response instantiates a new ListPods200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPods200ResponseWithDefaults

`func NewListPods200ResponseWithDefaults() *ListPods200Response`

NewListPods200ResponseWithDefaults instantiates a new ListPods200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPods200Response) GetData() InframonitoringtypesPods`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPods200Response) GetDataOk() (*InframonitoringtypesPods, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPods200Response) SetData(v InframonitoringtypesPods)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListPods200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPods200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPods200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


