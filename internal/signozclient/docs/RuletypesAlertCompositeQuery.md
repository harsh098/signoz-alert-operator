# RuletypesAlertCompositeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PanelType** | [**RuletypesPanelType**](RuletypesPanelType.md) |  | 
**Queries** | [**[]Querybuildertypesv5QueryEnvelope**](Querybuildertypesv5QueryEnvelope.md) |  | 
**QueryType** | [**RuletypesQueryType**](RuletypesQueryType.md) |  | 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewRuletypesAlertCompositeQuery

`func NewRuletypesAlertCompositeQuery(panelType RuletypesPanelType, queries []Querybuildertypesv5QueryEnvelope, queryType RuletypesQueryType, ) *RuletypesAlertCompositeQuery`

NewRuletypesAlertCompositeQuery instantiates a new RuletypesAlertCompositeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuletypesAlertCompositeQueryWithDefaults

`func NewRuletypesAlertCompositeQueryWithDefaults() *RuletypesAlertCompositeQuery`

NewRuletypesAlertCompositeQueryWithDefaults instantiates a new RuletypesAlertCompositeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPanelType

`func (o *RuletypesAlertCompositeQuery) GetPanelType() RuletypesPanelType`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *RuletypesAlertCompositeQuery) GetPanelTypeOk() (*RuletypesPanelType, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *RuletypesAlertCompositeQuery) SetPanelType(v RuletypesPanelType)`

SetPanelType sets PanelType field to given value.


### GetQueries

`func (o *RuletypesAlertCompositeQuery) GetQueries() []Querybuildertypesv5QueryEnvelope`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *RuletypesAlertCompositeQuery) GetQueriesOk() (*[]Querybuildertypesv5QueryEnvelope, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *RuletypesAlertCompositeQuery) SetQueries(v []Querybuildertypesv5QueryEnvelope)`

SetQueries sets Queries field to given value.


### SetQueriesNil

`func (o *RuletypesAlertCompositeQuery) SetQueriesNil(b bool)`

 SetQueriesNil sets the value for Queries to be an explicit nil

### UnsetQueries
`func (o *RuletypesAlertCompositeQuery) UnsetQueries()`

UnsetQueries ensures that no value is present for Queries, not even an explicit nil
### GetQueryType

`func (o *RuletypesAlertCompositeQuery) GetQueryType() RuletypesQueryType`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *RuletypesAlertCompositeQuery) GetQueryTypeOk() (*RuletypesQueryType, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *RuletypesAlertCompositeQuery) SetQueryType(v RuletypesQueryType)`

SetQueryType sets QueryType field to given value.


### GetUnit

`func (o *RuletypesAlertCompositeQuery) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RuletypesAlertCompositeQuery) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RuletypesAlertCompositeQuery) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *RuletypesAlertCompositeQuery) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


