# GetFeatures200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FeaturetypesGettableFeature**](FeaturetypesGettableFeature.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetFeatures200Response

`func NewGetFeatures200Response(data []FeaturetypesGettableFeature, status string, ) *GetFeatures200Response`

NewGetFeatures200Response instantiates a new GetFeatures200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeatures200ResponseWithDefaults

`func NewGetFeatures200ResponseWithDefaults() *GetFeatures200Response`

NewGetFeatures200ResponseWithDefaults instantiates a new GetFeatures200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFeatures200Response) GetData() []FeaturetypesGettableFeature`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFeatures200Response) GetDataOk() (*[]FeaturetypesGettableFeature, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFeatures200Response) SetData(v []FeaturetypesGettableFeature)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetFeatures200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFeatures200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFeatures200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


