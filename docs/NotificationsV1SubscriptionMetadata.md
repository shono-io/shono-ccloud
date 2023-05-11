# NotificationsV1SubscriptionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **interface{}** |  | 
**ResourceName** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewNotificationsV1SubscriptionMetadata

`func NewNotificationsV1SubscriptionMetadata(self interface{}, ) *NotificationsV1SubscriptionMetadata`

NewNotificationsV1SubscriptionMetadata instantiates a new NotificationsV1SubscriptionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SubscriptionMetadataWithDefaults

`func NewNotificationsV1SubscriptionMetadataWithDefaults() *NotificationsV1SubscriptionMetadata`

NewNotificationsV1SubscriptionMetadataWithDefaults instantiates a new NotificationsV1SubscriptionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NotificationsV1SubscriptionMetadata) GetSelf() interface{}`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NotificationsV1SubscriptionMetadata) GetSelfOk() (*interface{}, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NotificationsV1SubscriptionMetadata) SetSelf(v interface{})`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *NotificationsV1SubscriptionMetadata) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *NotificationsV1SubscriptionMetadata) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetResourceName

`func (o *NotificationsV1SubscriptionMetadata) GetResourceName() interface{}`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NotificationsV1SubscriptionMetadata) GetResourceNameOk() (*interface{}, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NotificationsV1SubscriptionMetadata) SetResourceName(v interface{})`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *NotificationsV1SubscriptionMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *NotificationsV1SubscriptionMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *NotificationsV1SubscriptionMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetCreatedAt

`func (o *NotificationsV1SubscriptionMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsV1SubscriptionMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsV1SubscriptionMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsV1SubscriptionMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsV1SubscriptionMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsV1SubscriptionMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsV1SubscriptionMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsV1SubscriptionMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NotificationsV1SubscriptionMetadata) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NotificationsV1SubscriptionMetadata) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NotificationsV1SubscriptionMetadata) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NotificationsV1SubscriptionMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


