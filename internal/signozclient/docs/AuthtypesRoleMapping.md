# AuthtypesRoleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRole** | Pointer to **string** |  | [optional] 
**GroupMappings** | Pointer to **map[string]string** |  | [optional] 
**UseRoleAttribute** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthtypesRoleMapping

`func NewAuthtypesRoleMapping() *AuthtypesRoleMapping`

NewAuthtypesRoleMapping instantiates a new AuthtypesRoleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesRoleMappingWithDefaults

`func NewAuthtypesRoleMappingWithDefaults() *AuthtypesRoleMapping`

NewAuthtypesRoleMappingWithDefaults instantiates a new AuthtypesRoleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRole

`func (o *AuthtypesRoleMapping) GetDefaultRole() string`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *AuthtypesRoleMapping) GetDefaultRoleOk() (*string, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *AuthtypesRoleMapping) SetDefaultRole(v string)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *AuthtypesRoleMapping) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetGroupMappings

`func (o *AuthtypesRoleMapping) GetGroupMappings() map[string]string`

GetGroupMappings returns the GroupMappings field if non-nil, zero value otherwise.

### GetGroupMappingsOk

`func (o *AuthtypesRoleMapping) GetGroupMappingsOk() (*map[string]string, bool)`

GetGroupMappingsOk returns a tuple with the GroupMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMappings

`func (o *AuthtypesRoleMapping) SetGroupMappings(v map[string]string)`

SetGroupMappings sets GroupMappings field to given value.

### HasGroupMappings

`func (o *AuthtypesRoleMapping) HasGroupMappings() bool`

HasGroupMappings returns a boolean if a field has been set.

### SetGroupMappingsNil

`func (o *AuthtypesRoleMapping) SetGroupMappingsNil(b bool)`

 SetGroupMappingsNil sets the value for GroupMappings to be an explicit nil

### UnsetGroupMappings
`func (o *AuthtypesRoleMapping) UnsetGroupMappings()`

UnsetGroupMappings ensures that no value is present for GroupMappings, not even an explicit nil
### GetUseRoleAttribute

`func (o *AuthtypesRoleMapping) GetUseRoleAttribute() bool`

GetUseRoleAttribute returns the UseRoleAttribute field if non-nil, zero value otherwise.

### GetUseRoleAttributeOk

`func (o *AuthtypesRoleMapping) GetUseRoleAttributeOk() (*bool, bool)`

GetUseRoleAttributeOk returns a tuple with the UseRoleAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoleAttribute

`func (o *AuthtypesRoleMapping) SetUseRoleAttribute(v bool)`

SetUseRoleAttribute sets UseRoleAttribute field to given value.

### HasUseRoleAttribute

`func (o *AuthtypesRoleMapping) HasUseRoleAttribute() bool`

HasUseRoleAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


