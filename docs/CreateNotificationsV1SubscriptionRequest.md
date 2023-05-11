# CreateNotificationsV1SubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**NotificationsV1SubscriptionMetadata**](NotificationsV1SubscriptionMetadata.md) |  | [optional] 
**CurrentState** | Pointer to **string** | Denotes the state of the subscription. When the subscription is ENABLED, the user will receive notification on the configured Integrations. If the subscription is DISABLED, the user will not recieve any notification for the configured notification type. Note that, you cannot disable a subscription for &#x60;REQUIRED&#x60; notification type.  | [optional] 
**NotificationType** | [**NotificationsV1SubscriptionNotificationType**](NotificationsV1SubscriptionNotificationType.md) |  | 
**Integrations** | [**[]GlobalObjectReference**](GlobalObjectReference.md) | Integrations to which notifications are to be sent. | 

## Methods

### NewCreateNotificationsV1SubscriptionRequest

`func NewCreateNotificationsV1SubscriptionRequest(notificationType NotificationsV1SubscriptionNotificationType, integrations []GlobalObjectReference, ) *CreateNotificationsV1SubscriptionRequest`

NewCreateNotificationsV1SubscriptionRequest instantiates a new CreateNotificationsV1SubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationsV1SubscriptionRequestWithDefaults

`func NewCreateNotificationsV1SubscriptionRequestWithDefaults() *CreateNotificationsV1SubscriptionRequest`

NewCreateNotificationsV1SubscriptionRequestWithDefaults instantiates a new CreateNotificationsV1SubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateNotificationsV1SubscriptionRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateNotificationsV1SubscriptionRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateNotificationsV1SubscriptionRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateNotificationsV1SubscriptionRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNotificationsV1SubscriptionRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateNotificationsV1SubscriptionRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateNotificationsV1SubscriptionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNotificationsV1SubscriptionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNotificationsV1SubscriptionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNotificationsV1SubscriptionRequest) GetMetadata() NotificationsV1SubscriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetMetadataOk() (*NotificationsV1SubscriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNotificationsV1SubscriptionRequest) SetMetadata(v NotificationsV1SubscriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNotificationsV1SubscriptionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCurrentState

`func (o *CreateNotificationsV1SubscriptionRequest) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *CreateNotificationsV1SubscriptionRequest) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *CreateNotificationsV1SubscriptionRequest) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetNotificationType

`func (o *CreateNotificationsV1SubscriptionRequest) GetNotificationType() NotificationsV1SubscriptionNotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetNotificationTypeOk() (*NotificationsV1SubscriptionNotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CreateNotificationsV1SubscriptionRequest) SetNotificationType(v NotificationsV1SubscriptionNotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetIntegrations

`func (o *CreateNotificationsV1SubscriptionRequest) GetIntegrations() []GlobalObjectReference`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *CreateNotificationsV1SubscriptionRequest) GetIntegrationsOk() (*[]GlobalObjectReference, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *CreateNotificationsV1SubscriptionRequest) SetIntegrations(v []GlobalObjectReference)`

SetIntegrations sets Integrations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


