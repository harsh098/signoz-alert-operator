# AuthtypesGettableTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | **bool** |  | 
**Object** | [**CoretypesObject**](CoretypesObject.md) |  | 
**Relation** | [**AuthtypesRelation**](AuthtypesRelation.md) |  | 

## Methods

### NewAuthtypesGettableTransaction

`func NewAuthtypesGettableTransaction(authorized bool, object CoretypesObject, relation AuthtypesRelation, ) *AuthtypesGettableTransaction`

NewAuthtypesGettableTransaction instantiates a new AuthtypesGettableTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthtypesGettableTransactionWithDefaults

`func NewAuthtypesGettableTransactionWithDefaults() *AuthtypesGettableTransaction`

NewAuthtypesGettableTransactionWithDefaults instantiates a new AuthtypesGettableTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *AuthtypesGettableTransaction) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *AuthtypesGettableTransaction) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *AuthtypesGettableTransaction) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.


### GetObject

`func (o *AuthtypesGettableTransaction) GetObject() CoretypesObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AuthtypesGettableTransaction) GetObjectOk() (*CoretypesObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AuthtypesGettableTransaction) SetObject(v CoretypesObject)`

SetObject sets Object field to given value.


### GetRelation

`func (o *AuthtypesGettableTransaction) GetRelation() AuthtypesRelation`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *AuthtypesGettableTransaction) GetRelationOk() (*AuthtypesRelation, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *AuthtypesGettableTransaction) SetRelation(v AuthtypesRelation)`

SetRelation sets Relation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


