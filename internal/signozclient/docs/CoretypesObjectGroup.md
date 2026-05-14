# CoretypesObjectGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**CoretypesResourceRef**](CoretypesResourceRef.md) |  | 
**Selectors** | **[]string** |  | 

## Methods

### NewCoretypesObjectGroup

`func NewCoretypesObjectGroup(resource CoretypesResourceRef, selectors []string, ) *CoretypesObjectGroup`

NewCoretypesObjectGroup instantiates a new CoretypesObjectGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoretypesObjectGroupWithDefaults

`func NewCoretypesObjectGroupWithDefaults() *CoretypesObjectGroup`

NewCoretypesObjectGroupWithDefaults instantiates a new CoretypesObjectGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *CoretypesObjectGroup) GetResource() CoretypesResourceRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CoretypesObjectGroup) GetResourceOk() (*CoretypesResourceRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CoretypesObjectGroup) SetResource(v CoretypesResourceRef)`

SetResource sets Resource field to given value.


### GetSelectors

`func (o *CoretypesObjectGroup) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *CoretypesObjectGroup) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *CoretypesObjectGroup) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


