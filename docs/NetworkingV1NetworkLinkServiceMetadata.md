# NetworkingV1NetworkLinkServiceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **interface{}** |  | 
**ResourceName** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkLinkServiceMetadata

`func NewNetworkingV1NetworkLinkServiceMetadata(self interface{}, ) *NetworkingV1NetworkLinkServiceMetadata`

NewNetworkingV1NetworkLinkServiceMetadata instantiates a new NetworkingV1NetworkLinkServiceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceMetadataWithDefaults

`func NewNetworkingV1NetworkLinkServiceMetadataWithDefaults() *NetworkingV1NetworkLinkServiceMetadata`

NewNetworkingV1NetworkLinkServiceMetadataWithDefaults instantiates a new NetworkingV1NetworkLinkServiceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetSelf() interface{}`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetSelfOk() (*interface{}, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetSelf(v interface{})`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *NetworkingV1NetworkLinkServiceMetadata) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetResourceName

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetResourceName() interface{}`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetResourceNameOk() (*interface{}, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetResourceName(v interface{})`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *NetworkingV1NetworkLinkServiceMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *NetworkingV1NetworkLinkServiceMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetCreatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NetworkingV1NetworkLinkServiceMetadata) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NetworkingV1NetworkLinkServiceMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


