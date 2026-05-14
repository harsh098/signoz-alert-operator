# Querybuildertypesv5QueryRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompositeQuery** | Pointer to [**Querybuildertypesv5CompositeQuery**](Querybuildertypesv5CompositeQuery.md) |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**FormatOptions** | Pointer to [**Querybuildertypesv5FormatOptions**](Querybuildertypesv5FormatOptions.md) |  | [optional] 
**NoCache** | Pointer to **bool** |  | [optional] 
**RequestType** | Pointer to [**Querybuildertypesv5RequestType**](Querybuildertypesv5RequestType.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Variables** | Pointer to [**map[string]Querybuildertypesv5VariableItem**](Querybuildertypesv5VariableItem.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5QueryRangeRequest

`func NewQuerybuildertypesv5QueryRangeRequest() *Querybuildertypesv5QueryRangeRequest`

NewQuerybuildertypesv5QueryRangeRequest instantiates a new Querybuildertypesv5QueryRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5QueryRangeRequestWithDefaults

`func NewQuerybuildertypesv5QueryRangeRequestWithDefaults() *Querybuildertypesv5QueryRangeRequest`

NewQuerybuildertypesv5QueryRangeRequestWithDefaults instantiates a new Querybuildertypesv5QueryRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositeQuery

`func (o *Querybuildertypesv5QueryRangeRequest) GetCompositeQuery() Querybuildertypesv5CompositeQuery`

GetCompositeQuery returns the CompositeQuery field if non-nil, zero value otherwise.

### GetCompositeQueryOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetCompositeQueryOk() (*Querybuildertypesv5CompositeQuery, bool)`

GetCompositeQueryOk returns a tuple with the CompositeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeQuery

`func (o *Querybuildertypesv5QueryRangeRequest) SetCompositeQuery(v Querybuildertypesv5CompositeQuery)`

SetCompositeQuery sets CompositeQuery field to given value.

### HasCompositeQuery

`func (o *Querybuildertypesv5QueryRangeRequest) HasCompositeQuery() bool`

HasCompositeQuery returns a boolean if a field has been set.

### GetEnd

`func (o *Querybuildertypesv5QueryRangeRequest) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Querybuildertypesv5QueryRangeRequest) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Querybuildertypesv5QueryRangeRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFormatOptions

`func (o *Querybuildertypesv5QueryRangeRequest) GetFormatOptions() Querybuildertypesv5FormatOptions`

GetFormatOptions returns the FormatOptions field if non-nil, zero value otherwise.

### GetFormatOptionsOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetFormatOptionsOk() (*Querybuildertypesv5FormatOptions, bool)`

GetFormatOptionsOk returns a tuple with the FormatOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatOptions

`func (o *Querybuildertypesv5QueryRangeRequest) SetFormatOptions(v Querybuildertypesv5FormatOptions)`

SetFormatOptions sets FormatOptions field to given value.

### HasFormatOptions

`func (o *Querybuildertypesv5QueryRangeRequest) HasFormatOptions() bool`

HasFormatOptions returns a boolean if a field has been set.

### GetNoCache

`func (o *Querybuildertypesv5QueryRangeRequest) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *Querybuildertypesv5QueryRangeRequest) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *Querybuildertypesv5QueryRangeRequest) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetRequestType

`func (o *Querybuildertypesv5QueryRangeRequest) GetRequestType() Querybuildertypesv5RequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetRequestTypeOk() (*Querybuildertypesv5RequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *Querybuildertypesv5QueryRangeRequest) SetRequestType(v Querybuildertypesv5RequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *Querybuildertypesv5QueryRangeRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *Querybuildertypesv5QueryRangeRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *Querybuildertypesv5QueryRangeRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *Querybuildertypesv5QueryRangeRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetStart

`func (o *Querybuildertypesv5QueryRangeRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Querybuildertypesv5QueryRangeRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Querybuildertypesv5QueryRangeRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetVariables

`func (o *Querybuildertypesv5QueryRangeRequest) GetVariables() map[string]Querybuildertypesv5VariableItem`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Querybuildertypesv5QueryRangeRequest) GetVariablesOk() (*map[string]Querybuildertypesv5VariableItem, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Querybuildertypesv5QueryRangeRequest) SetVariables(v map[string]Querybuildertypesv5VariableItem)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Querybuildertypesv5QueryRangeRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


