# Querybuildertypesv5QueryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** |  | [optional] 
**Aggregations** | Pointer to [**[]Querybuildertypesv5AggregationBucket**](Querybuildertypesv5AggregationBucket.md) |  | [optional] 
**QueryName** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]Querybuildertypesv5ColumnDescriptor**](Querybuildertypesv5ColumnDescriptor.md) |  | [optional] 
**Data** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]Querybuildertypesv5RawRow**](Querybuildertypesv5RawRow.md) |  | [optional] 

## Methods

### NewQuerybuildertypesv5QueryData

`func NewQuerybuildertypesv5QueryData() *Querybuildertypesv5QueryData`

NewQuerybuildertypesv5QueryData instantiates a new Querybuildertypesv5QueryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5QueryDataWithDefaults

`func NewQuerybuildertypesv5QueryDataWithDefaults() *Querybuildertypesv5QueryData`

NewQuerybuildertypesv5QueryDataWithDefaults instantiates a new Querybuildertypesv5QueryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *Querybuildertypesv5QueryData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Querybuildertypesv5QueryData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Querybuildertypesv5QueryData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *Querybuildertypesv5QueryData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *Querybuildertypesv5QueryData) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *Querybuildertypesv5QueryData) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetAggregations

`func (o *Querybuildertypesv5QueryData) GetAggregations() []Querybuildertypesv5AggregationBucket`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *Querybuildertypesv5QueryData) GetAggregationsOk() (*[]Querybuildertypesv5AggregationBucket, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *Querybuildertypesv5QueryData) SetAggregations(v []Querybuildertypesv5AggregationBucket)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *Querybuildertypesv5QueryData) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### SetAggregationsNil

`func (o *Querybuildertypesv5QueryData) SetAggregationsNil(b bool)`

 SetAggregationsNil sets the value for Aggregations to be an explicit nil

### UnsetAggregations
`func (o *Querybuildertypesv5QueryData) UnsetAggregations()`

UnsetAggregations ensures that no value is present for Aggregations, not even an explicit nil
### GetQueryName

`func (o *Querybuildertypesv5QueryData) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *Querybuildertypesv5QueryData) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *Querybuildertypesv5QueryData) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *Querybuildertypesv5QueryData) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.

### GetColumns

`func (o *Querybuildertypesv5QueryData) GetColumns() []Querybuildertypesv5ColumnDescriptor`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Querybuildertypesv5QueryData) GetColumnsOk() (*[]Querybuildertypesv5ColumnDescriptor, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Querybuildertypesv5QueryData) SetColumns(v []Querybuildertypesv5ColumnDescriptor)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Querybuildertypesv5QueryData) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *Querybuildertypesv5QueryData) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *Querybuildertypesv5QueryData) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetData

`func (o *Querybuildertypesv5QueryData) GetData() [][]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Querybuildertypesv5QueryData) GetDataOk() (*[][]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Querybuildertypesv5QueryData) SetData(v [][]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Querybuildertypesv5QueryData) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Querybuildertypesv5QueryData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Querybuildertypesv5QueryData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNextCursor

`func (o *Querybuildertypesv5QueryData) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *Querybuildertypesv5QueryData) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *Querybuildertypesv5QueryData) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *Querybuildertypesv5QueryData) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetRows

`func (o *Querybuildertypesv5QueryData) GetRows() []Querybuildertypesv5RawRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *Querybuildertypesv5QueryData) GetRowsOk() (*[]Querybuildertypesv5RawRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *Querybuildertypesv5QueryData) SetRows(v []Querybuildertypesv5RawRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *Querybuildertypesv5QueryData) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRowsNil

`func (o *Querybuildertypesv5QueryData) SetRowsNil(b bool)`

 SetRowsNil sets the value for Rows to be an explicit nil

### UnsetRows
`func (o *Querybuildertypesv5QueryData) UnsetRows()`

UnsetRows ensures that no value is present for Rows, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


