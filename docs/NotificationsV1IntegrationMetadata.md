# NotificationsV1IntegrationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **interface{}** |  | 
**ResourceName** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 

## Methods

### NewNotificationsV1IntegrationMetadata

`func NewNotificationsV1IntegrationMetadata(self interface{}, ) *NotificationsV1IntegrationMetadata`

NewNotificationsV1IntegrationMetadata instantiates a new NotificationsV1IntegrationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1IntegrationMetadataWithDefaults

`func NewNotificationsV1IntegrationMetadataWithDefaults() *NotificationsV1IntegrationMetadata`

NewNotificationsV1IntegrationMetadataWithDefaults instantiates a new NotificationsV1IntegrationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NotificationsV1IntegrationMetadata) GetSelf() interface{}`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NotificationsV1IntegrationMetadata) GetSelfOk() (*interface{}, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NotificationsV1IntegrationMetadata) SetSelf(v interface{})`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *NotificationsV1IntegrationMetadata) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *NotificationsV1IntegrationMetadata) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetResourceName

`func (o *NotificationsV1IntegrationMetadata) GetResourceName() interface{}`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NotificationsV1IntegrationMetadata) GetResourceNameOk() (*interface{}, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NotificationsV1IntegrationMetadata) SetResourceName(v interface{})`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *NotificationsV1IntegrationMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *NotificationsV1IntegrationMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *NotificationsV1IntegrationMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetCreatedAt

`func (o *NotificationsV1IntegrationMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsV1IntegrationMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsV1IntegrationMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsV1IntegrationMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsV1IntegrationMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsV1IntegrationMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsV1IntegrationMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsV1IntegrationMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *NotificationsV1IntegrationMetadata) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *NotificationsV1IntegrationMetadata) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *NotificationsV1IntegrationMetadata) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *NotificationsV1IntegrationMetadata) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


