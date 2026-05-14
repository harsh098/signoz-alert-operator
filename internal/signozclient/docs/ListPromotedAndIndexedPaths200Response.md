# ListPromotedAndIndexedPaths200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PromotetypesPromotePath**](PromotetypesPromotePath.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListPromotedAndIndexedPaths200Response

`func NewListPromotedAndIndexedPaths200Response(data []PromotetypesPromotePath, status string, ) *ListPromotedAndIndexedPaths200Response`

NewListPromotedAndIndexedPaths200Response instantiates a new ListPromotedAndIndexedPaths200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPromotedAndIndexedPaths200ResponseWithDefaults

`func NewListPromotedAndIndexedPaths200ResponseWithDefaults() *ListPromotedAndIndexedPaths200Response`

NewListPromotedAndIndexedPaths200ResponseWithDefaults instantiates a new ListPromotedAndIndexedPaths200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPromotedAndIndexedPaths200Response) GetData() []PromotetypesPromotePath`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPromotedAndIndexedPaths200Response) GetDataOk() (*[]PromotetypesPromotePath, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPromotedAndIndexedPaths200Response) SetData(v []PromotetypesPromotePath)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ListPromotedAndIndexedPaths200Response) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ListPromotedAndIndexedPaths200Response) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetStatus

`func (o *ListPromotedAndIndexedPaths200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPromotedAndIndexedPaths200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPromotedAndIndexedPaths200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


