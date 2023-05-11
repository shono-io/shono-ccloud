# NetworkingV1NetworkLinkEndpointMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **interface{}** |  | 
**ResourceName** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkLinkEndpointMetadata

`func NewNetworkingV1NetworkLinkEndpointMetadata(self interface{}, ) *NetworkingV1NetworkLinkEndpointMetadata`

NewNetworkingV1NetworkLinkEndpointMetadata instantiates a new NetworkingV1NetworkLinkEndpointMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkEndpointMetadataWithDefaults

`func NewNetworkingV1NetworkLinkEndpointMetadataWithDefaults() *NetworkingV1NetworkLinkEndpointMetadata`

NewNetworkingV1NetworkLinkEndpointMetadataWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetSelf() interface{}`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetSelfOk() (*interface{}, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetSelf(v interface{})`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *NetworkingV1NetworkLinkEndpointMetadata) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetResourceName

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetResourceName() interface{}`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetResourceNameOk() (*interface{}, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetResourceName(v interface{})`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *NetworkingV1NetworkLinkEndpointMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *NetworkingV1NetworkLinkEndpointMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetCreatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NetworkingV1NetworkLinkEndpointMetadata) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NetworkingV1NetworkLinkEndpointMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


