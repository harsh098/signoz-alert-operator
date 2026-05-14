# Querybuildertypesv5ColumnDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationIndex** | Pointer to **int64** |  | [optional] 
**ColumnType** | Pointer to [**Querybuildertypesv5ColumnType**](Querybuildertypesv5ColumnType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FieldContext** | Pointer to [**TelemetrytypesFieldContext**](TelemetrytypesFieldContext.md) |  | [optional] 
**FieldDataType** | Pointer to [**TelemetrytypesFieldDataType**](TelemetrytypesFieldDataType.md) |  | [optional] 
**Meta** | Pointer to [**Querybuildertypesv5AggregationBucketMeta**](Querybuildertypesv5AggregationBucketMeta.md) |  | [optional] 
**Name** | **string** |  | 
**QueryName** | Pointer to **string** |  | [optional] 
**Signal** | Pointer to [**TelemetrytypesSignal**](TelemetrytypesSignal.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewQuerybuildertypesv5ColumnDescriptor

`func NewQuerybuildertypesv5ColumnDescriptor(name string, ) *Querybuildertypesv5ColumnDescriptor`

NewQuerybuildertypesv5ColumnDescriptor instantiates a new Querybuildertypesv5ColumnDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerybuildertypesv5ColumnDescriptorWithDefaults

`func NewQuerybuildertypesv5ColumnDescriptorWithDefaults() *Querybuildertypesv5ColumnDescriptor`

NewQuerybuildertypesv5ColumnDescriptorWithDefaults instantiates a new Querybuildertypesv5ColumnDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationIndex

`func (o *Querybuildertypesv5ColumnDescriptor) GetAggregationIndex() int64`

GetAggregationIndex returns the AggregationIndex field if non-nil, zero value otherwise.

### GetAggregationIndexOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetAggregationIndexOk() (*int64, bool)`

GetAggregationIndexOk returns a tuple with the AggregationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationIndex

`func (o *Querybuildertypesv5ColumnDescriptor) SetAggregationIndex(v int64)`

SetAggregationIndex sets AggregationIndex field to given value.

### HasAggregationIndex

`func (o *Querybuildertypesv5ColumnDescriptor) HasAggregationIndex() bool`

HasAggregationIndex returns a boolean if a field has been set.

### GetColumnType

`func (o *Querybuildertypesv5ColumnDescriptor) GetColumnType() Querybuildertypesv5ColumnType`

GetColumnType returns the ColumnType field if non-nil, zero value otherwise.

### GetColumnTypeOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetColumnTypeOk() (*Querybuildertypesv5ColumnType, bool)`

GetColumnTypeOk returns a tuple with the ColumnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnType

`func (o *Querybuildertypesv5ColumnDescriptor) SetColumnType(v Querybuildertypesv5ColumnType)`

SetColumnType sets ColumnType field to given value.

### HasColumnType

`func (o *Querybuildertypesv5ColumnDescriptor) HasColumnType() bool`

HasColumnType returns a boolean if a field has been set.

### GetDescription

`func (o *Querybuildertypesv5ColumnDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Querybuildertypesv5ColumnDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Querybuildertypesv5ColumnDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldContext

`func (o *Querybuildertypesv5ColumnDescriptor) GetFieldContext() TelemetrytypesFieldContext`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetFieldContextOk() (*TelemetrytypesFieldContext, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *Querybuildertypesv5ColumnDescriptor) SetFieldContext(v TelemetrytypesFieldContext)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *Querybuildertypesv5ColumnDescriptor) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldDataType

`func (o *Querybuildertypesv5ColumnDescriptor) GetFieldDataType() TelemetrytypesFieldDataType`

GetFieldDataType returns the FieldDataType field if non-nil, zero value otherwise.

### GetFieldDataTypeOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetFieldDataTypeOk() (*TelemetrytypesFieldDataType, bool)`

GetFieldDataTypeOk returns a tuple with the FieldDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDataType

`func (o *Querybuildertypesv5ColumnDescriptor) SetFieldDataType(v TelemetrytypesFieldDataType)`

SetFieldDataType sets FieldDataType field to given value.

### HasFieldDataType

`func (o *Querybuildertypesv5ColumnDescriptor) HasFieldDataType() bool`

HasFieldDataType returns a boolean if a field has been set.

### GetMeta

`func (o *Querybuildertypesv5ColumnDescriptor) GetMeta() Querybuildertypesv5AggregationBucketMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetMetaOk() (*Querybuildertypesv5AggregationBucketMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Querybuildertypesv5ColumnDescriptor) SetMeta(v Querybuildertypesv5AggregationBucketMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Querybuildertypesv5ColumnDescriptor) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *Querybuildertypesv5ColumnDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Querybuildertypesv5ColumnDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetQueryName

`func (o *Querybuildertypesv5ColumnDescriptor) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *Querybuildertypesv5ColumnDescriptor) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *Querybuildertypesv5ColumnDescriptor) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.

### GetSignal

`func (o *Querybuildertypesv5ColumnDescriptor) GetSignal() TelemetrytypesSignal`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetSignalOk() (*TelemetrytypesSignal, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *Querybuildertypesv5ColumnDescriptor) SetSignal(v TelemetrytypesSignal)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *Querybuildertypesv5ColumnDescriptor) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetUnit

`func (o *Querybuildertypesv5ColumnDescriptor) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Querybuildertypesv5ColumnDescriptor) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Querybuildertypesv5ColumnDescriptor) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Querybuildertypesv5ColumnDescriptor) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


