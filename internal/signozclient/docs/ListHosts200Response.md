# ListHosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**InframonitoringtypesHosts**](InframonitoringtypesHosts.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListHosts200Response

`func NewListHosts200Response(data InframonitoringtypesHosts, status string, ) *ListHosts200Response`

NewListHosts200Response instantiates a new ListHosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseWithDefaults

`func NewListHosts200ResponseWithDefaults() *ListHosts200Response`

NewListHosts200ResponseWithDefaults instantiates a new ListHosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListHosts200Response) GetData() InframonitoringtypesHosts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListHosts200Response) GetDataOk() (*InframonitoringtypesHosts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListHosts200Response) SetData(v InframonitoringtypesHosts)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListHosts200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHosts200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHosts200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


