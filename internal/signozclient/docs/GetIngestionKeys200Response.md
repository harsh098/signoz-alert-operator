# GetIngestionKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GatewaytypesGettableIngestionKeys**](GatewaytypesGettableIngestionKeys.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetIngestionKeys200Response

`func NewGetIngestionKeys200Response(data GatewaytypesGettableIngestionKeys, status string, ) *GetIngestionKeys200Response`

NewGetIngestionKeys200Response instantiates a new GetIngestionKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIngestionKeys200ResponseWithDefaults

`func NewGetIngestionKeys200ResponseWithDefaults() *GetIngestionKeys200Response`

NewGetIngestionKeys200ResponseWithDefaults instantiates a new GetIngestionKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIngestionKeys200Response) GetData() GatewaytypesGettableIngestionKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIngestionKeys200Response) GetDataOk() (*GatewaytypesGettableIngestionKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIngestionKeys200Response) SetData(v GatewaytypesGettableIngestionKeys)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetIngestionKeys200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIngestionKeys200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIngestionKeys200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


