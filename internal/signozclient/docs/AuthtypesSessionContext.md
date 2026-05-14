# AuthtypesSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** |  | [optional] 
**Orgs** | Pointer to [**[]AuthtypesOrgSessionContext**](AuthtypesOrgSessionContext.md) |  | [optional] 

## Methods

### NewAuthtypesSessionContext

`func NewAuthtypesSessionContext() *AuthtypesSessionContext`

NewAuthtypesSessionContext instantiates a new AuthtypesSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesSessionContextWithDefaults

`func NewAuthtypesSessionContextWithDefaults() *AuthtypesSessionContext`

NewAuthtypesSessionContextWithDefaults instantiates a new AuthtypesSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *AuthtypesSessionContext) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *AuthtypesSessionContext) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *AuthtypesSessionContext) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *AuthtypesSessionContext) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetOrgs

`func (o *AuthtypesSessionContext) GetOrgs() []AuthtypesOrgSessionContext`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *AuthtypesSessionContext) GetOrgsOk() (*[]AuthtypesOrgSessionContext, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *AuthtypesSessionContext) SetOrgs(v []AuthtypesOrgSessionContext)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *AuthtypesSessionContext) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### SetOrgsNil

`func (o *AuthtypesSessionContext) SetOrgsNil(b bool)`

 SetOrgsNil sets the value for Orgs to be an explicit nil

### UnsetOrgs
`func (o *AuthtypesSessionContext) UnsetOrgs()`

UnsetOrgs ensures that no value is present for Orgs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


