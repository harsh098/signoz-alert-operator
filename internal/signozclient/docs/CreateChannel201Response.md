# CreateChannel201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AlertmanagertypesChannel**](AlertmanagertypesChannel.md) |  | 
**Status** | **string** |  | 

## Methods

### NewCreateChannel201Response

`func NewCreateChannel201Response(data AlertmanagertypesChannel, status string, ) *CreateChannel201Response`

NewCreateChannel201Response instantiates a new CreateChannel201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChannel201ResponseWithDefaults

`func NewCreateChannel201ResponseWithDefaults() *CreateChannel201Response`

NewCreateChannel201ResponseWithDefaults instantiates a new CreateChannel201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateChannel201Response) GetData() AlertmanagertypesChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateChannel201Response) GetDataOk() (*AlertmanagertypesChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateChannel201Response) SetData(v AlertmanagertypesChannel)`

SetData sets Data field to given value.


### GetStatus

`func (o *CreateChannel201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateChannel201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateChannel201Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


