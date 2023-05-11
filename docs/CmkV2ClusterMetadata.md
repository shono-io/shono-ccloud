# CmkV2ClusterMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **interface{}** |  | 
**ResourceName** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewCmkV2ClusterMetadata

`func NewCmkV2ClusterMetadata(self interface{}, ) *CmkV2ClusterMetadata`

NewCmkV2ClusterMetadata instantiates a new CmkV2ClusterMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2ClusterMetadataWithDefaults

`func NewCmkV2ClusterMetadataWithDefaults() *CmkV2ClusterMetadata`

NewCmkV2ClusterMetadataWithDefaults instantiates a new CmkV2ClusterMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *CmkV2ClusterMetadata) GetSelf() interface{}`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CmkV2ClusterMetadata) GetSelfOk() (*interface{}, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CmkV2ClusterMetadata) SetSelf(v interface{})`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *CmkV2ClusterMetadata) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *CmkV2ClusterMetadata) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetResourceName

`func (o *CmkV2ClusterMetadata) GetResourceName() interface{}`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CmkV2ClusterMetadata) GetResourceNameOk() (*interface{}, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CmkV2ClusterMetadata) SetResourceName(v interface{})`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *CmkV2ClusterMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *CmkV2ClusterMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *CmkV2ClusterMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetCreatedAt

`func (o *CmkV2ClusterMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CmkV2ClusterMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CmkV2ClusterMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CmkV2ClusterMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CmkV2ClusterMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CmkV2ClusterMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CmkV2ClusterMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CmkV2ClusterMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CmkV2ClusterMetadata) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CmkV2ClusterMetadata) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CmkV2ClusterMetadata) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CmkV2ClusterMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


