# GetOrgPreference200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**PreferencetypesPreference**](PreferencetypesPreference.md) |  | 
**Status** | **string** |  | 

## Methods

### NewGetOrgPreference200Response

`func NewGetOrgPreference200Response(data PreferencetypesPreference, status string, ) *GetOrgPreference200Response`

NewGetOrgPreference200Response instantiates a new GetOrgPreference200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgPreference200ResponseWithDefaults

`func NewGetOrgPreference200ResponseWithDefaults() *GetOrgPreference200Response`

NewGetOrgPreference200ResponseWithDefaults instantiates a new GetOrgPreference200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrgPreference200Response) GetData() PreferencetypesPreference`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrgPreference200Response) GetDataOk() (*PreferencetypesPreference, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrgPreference200Response) SetData(v PreferencetypesPreference)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetOrgPreference200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrgPreference200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrgPreference200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


