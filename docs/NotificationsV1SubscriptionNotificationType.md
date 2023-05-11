# NotificationsV1SubscriptionNotificationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewNotificationsV1SubscriptionNotificationType

`func NewNotificationsV1SubscriptionNotificationType(id string, related string, resourceName string, ) *NotificationsV1SubscriptionNotificationType`

NewNotificationsV1SubscriptionNotificationType instantiates a new NotificationsV1SubscriptionNotificationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SubscriptionNotificationTypeWithDefaults

`func NewNotificationsV1SubscriptionNotificationTypeWithDefaults() *NotificationsV1SubscriptionNotificationType`

NewNotificationsV1SubscriptionNotificationTypeWithDefaults instantiates a new NotificationsV1SubscriptionNotificationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationsV1SubscriptionNotificationType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1SubscriptionNotificationType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1SubscriptionNotificationType) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *NotificationsV1SubscriptionNotificationType) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *NotificationsV1SubscriptionNotificationType) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *NotificationsV1SubscriptionNotificationType) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *NotificationsV1SubscriptionNotificationType) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NotificationsV1SubscriptionNotificationType) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NotificationsV1SubscriptionNotificationType) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


