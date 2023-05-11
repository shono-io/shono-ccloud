# GetNotificationsV1Subscription200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**NotificationsV1SubscriptionMetadata**](NotificationsV1SubscriptionMetadata.md) |  | [optional] 
**CurrentState** | Pointer to **string** | Denotes the state of the subscription. When the subscription is ENABLED, the user will receive notification on the configured Integrations. If the subscription is DISABLED, the user will not recieve any notification for the configured notification type. Note that, you cannot disable a subscription for &#x60;REQUIRED&#x60; notification type.  | [optional] 
**NotificationType** | [**NotificationsV1SubscriptionNotificationType**](NotificationsV1SubscriptionNotificationType.md) |  | 
**Integrations** | [**[]GlobalObjectReference**](GlobalObjectReference.md) | Integrations to which notifications are to be sent. | 

## Methods

### NewGetNotificationsV1Subscription200Response

`func NewGetNotificationsV1Subscription200Response(apiVersion string, kind string, id string, notificationType NotificationsV1SubscriptionNotificationType, integrations []GlobalObjectReference, ) *GetNotificationsV1Subscription200Response`

NewGetNotificationsV1Subscription200Response instantiates a new GetNotificationsV1Subscription200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsV1Subscription200ResponseWithDefaults

`func NewGetNotificationsV1Subscription200ResponseWithDefaults() *GetNotificationsV1Subscription200Response`

NewGetNotificationsV1Subscription200ResponseWithDefaults instantiates a new GetNotificationsV1Subscription200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetNotificationsV1Subscription200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetNotificationsV1Subscription200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetNotificationsV1Subscription200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetNotificationsV1Subscription200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetNotificationsV1Subscription200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetNotificationsV1Subscription200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetNotificationsV1Subscription200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNotificationsV1Subscription200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNotificationsV1Subscription200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetNotificationsV1Subscription200Response) GetMetadata() NotificationsV1SubscriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNotificationsV1Subscription200Response) GetMetadataOk() (*NotificationsV1SubscriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNotificationsV1Subscription200Response) SetMetadata(v NotificationsV1SubscriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNotificationsV1Subscription200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCurrentState

`func (o *GetNotificationsV1Subscription200Response) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *GetNotificationsV1Subscription200Response) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *GetNotificationsV1Subscription200Response) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *GetNotificationsV1Subscription200Response) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetNotificationType

`func (o *GetNotificationsV1Subscription200Response) GetNotificationType() NotificationsV1SubscriptionNotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *GetNotificationsV1Subscription200Response) GetNotificationTypeOk() (*NotificationsV1SubscriptionNotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *GetNotificationsV1Subscription200Response) SetNotificationType(v NotificationsV1SubscriptionNotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetIntegrations

`func (o *GetNotificationsV1Subscription200Response) GetIntegrations() []GlobalObjectReference`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *GetNotificationsV1Subscription200Response) GetIntegrationsOk() (*[]GlobalObjectReference, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *GetNotificationsV1Subscription200Response) SetIntegrations(v []GlobalObjectReference)`

SetIntegrations sets Integrations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


