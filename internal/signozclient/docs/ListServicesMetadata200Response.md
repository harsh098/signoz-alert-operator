# ListServicesMetadata200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CloudintegrationtypesGettableServicesMetadata**](CloudintegrationtypesGettableServicesMetadata.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListServicesMetadata200Response

`func NewListServicesMetadata200Response(data CloudintegrationtypesGettableServicesMetadata, status string, ) *ListServicesMetadata200Response`

NewListServicesMetadata200Response instantiates a new ListServicesMetadata200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicesMetadata200ResponseWithDefaults

`func NewListServicesMetadata200ResponseWithDefaults() *ListServicesMetadata200Response`

NewListServicesMetadata200ResponseWithDefaults instantiates a new ListServicesMetadata200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListServicesMetadata200Response) GetData() CloudintegrationtypesGettableServicesMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListServicesMetadata200Response) GetDataOk() (*CloudintegrationtypesGettableServicesMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListServicesMetadata200Response) SetData(v CloudintegrationtypesGettableServicesMetadata)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListServicesMetadata200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListServicesMetadata200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListServicesMetadata200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


