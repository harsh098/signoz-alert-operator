# ListAuthDomains200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AuthtypesGettableAuthDomain**](AuthtypesGettableAuthDomain.md) |  | 
**Status** | **string** |  | 

## Methods

### NewListAuthDomains200Response

`func NewListAuthDomains200Response(data []AuthtypesGettableAuthDomain, status string, ) *ListAuthDomains200Response`

NewListAuthDomains200Response instantiates a new ListAuthDomains200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthDomains200ResponseWithDefaults

`func NewListAuthDomains200ResponseWithDefaults() *ListAuthDomains200Response`

NewListAuthDomains200ResponseWithDefaults instantiates a new ListAuthDomains200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAuthDomains200Response) GetData() []AuthtypesGettableAuthDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAuthDomains200Response) GetDataOk() (*[]AuthtypesGettableAuthDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAuthDomains200Response) SetData(v []AuthtypesGettableAuthDomain)`

SetData sets Data field to given value.


### GetStatus

`func (o *ListAuthDomains200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAuthDomains200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAuthDomains200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


