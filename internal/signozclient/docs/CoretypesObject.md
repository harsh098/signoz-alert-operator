# CoretypesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**CoretypesResourceRef**](CoretypesResourceRef.md) |  | 
**Selector** | **string** |  | 

## Methods

### NewCoretypesObject

`func NewCoretypesObject(resource CoretypesResourceRef, selector string, ) *CoretypesObject`

NewCoretypesObject instantiates a new CoretypesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoretypesObjectWithDefaults

`func NewCoretypesObjectWithDefaults() *CoretypesObject`

NewCoretypesObjectWithDefaults instantiates a new CoretypesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *CoretypesObject) GetResource() CoretypesResourceRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CoretypesObject) GetResourceOk() (*CoretypesResourceRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CoretypesObject) SetResource(v CoretypesResourceRef)`

SetResource sets Resource field to given value.


### GetSelector

`func (o *CoretypesObject) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *CoretypesObject) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *CoretypesObject) SetSelector(v string)`

SetSelector sets Selector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


