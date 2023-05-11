# GetNotificationsV1NotificationType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**NotificationsV1NotificationTypeMetadata**](NotificationsV1NotificationTypeMetadata.md) |  | [optional] 
**DisplayName** | **string** | Human readable display name of the notification type  | 
**Category** | **string** | Represents the group with which the notification is associated. Notifications are grouped under certain categories for better organization. - BILLING_LICENSING: All billing, payments or licensing related notifications are grouped here. - SECURITY: All Confluent Cloud and Platform security related notifications are grouped here. - SERVICE: All Confluent services (eg. Kafka, Schema Registry, Connect etc.) related notifications are   grouped here. - ACCOUNT: All Confluent account related notifications are grouped here. For example: Billing, payment or license related notifications are grouped in BILLING_LICENSING category.  | 
**Description** | **string** | Human readable description of the notification type  | 
**SubscriptionPriority** | **string** | Indicates whether the notification is auto-subscribed and if the user can opt-out. - REQUIRED: the user is auto-subscribed to this notification and can&#39;t opt-out. - RECOMMENDED: the user is auto-subscribed to this notification and can opt-out. - OPTIONAL: the user is not auto-subscribed to this notification but can explicitly subscribe to it.  | 
**IsIncludedInPlan** | **bool** | Whether this notification is available to subscribe or not as per the user&#39;s current billing plan.  | 
**Severity** | **string** | Severity indicates the impact of this notification. - CRITICAL: a high impact notification which needs immediate attention. - WARN: a warning notification which can be addressed now or later. - INFO: an informational notification.  | 

## Methods

### NewGetNotificationsV1NotificationType200Response

`func NewGetNotificationsV1NotificationType200Response(apiVersion string, kind string, id string, displayName string, category string, description string, subscriptionPriority string, isIncludedInPlan bool, severity string, ) *GetNotificationsV1NotificationType200Response`

NewGetNotificationsV1NotificationType200Response instantiates a new GetNotificationsV1NotificationType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsV1NotificationType200ResponseWithDefaults

`func NewGetNotificationsV1NotificationType200ResponseWithDefaults() *GetNotificationsV1NotificationType200Response`

NewGetNotificationsV1NotificationType200ResponseWithDefaults instantiates a new GetNotificationsV1NotificationType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetNotificationsV1NotificationType200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetNotificationsV1NotificationType200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetNotificationsV1NotificationType200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetNotificationsV1NotificationType200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetNotificationsV1NotificationType200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetNotificationsV1NotificationType200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetNotificationsV1NotificationType200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNotificationsV1NotificationType200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNotificationsV1NotificationType200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetNotificationsV1NotificationType200Response) GetMetadata() NotificationsV1NotificationTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNotificationsV1NotificationType200Response) GetMetadataOk() (*NotificationsV1NotificationTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNotificationsV1NotificationType200Response) SetMetadata(v NotificationsV1NotificationTypeMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNotificationsV1NotificationType200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNotificationsV1NotificationType200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNotificationsV1NotificationType200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNotificationsV1NotificationType200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCategory

`func (o *GetNotificationsV1NotificationType200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNotificationsV1NotificationType200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNotificationsV1NotificationType200Response) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *GetNotificationsV1NotificationType200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNotificationsV1NotificationType200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNotificationsV1NotificationType200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubscriptionPriority

`func (o *GetNotificationsV1NotificationType200Response) GetSubscriptionPriority() string`

GetSubscriptionPriority returns the SubscriptionPriority field if non-nil, zero value otherwise.

### GetSubscriptionPriorityOk

`func (o *GetNotificationsV1NotificationType200Response) GetSubscriptionPriorityOk() (*string, bool)`

GetSubscriptionPriorityOk returns a tuple with the SubscriptionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPriority

`func (o *GetNotificationsV1NotificationType200Response) SetSubscriptionPriority(v string)`

SetSubscriptionPriority sets SubscriptionPriority field to given value.


### GetIsIncludedInPlan

`func (o *GetNotificationsV1NotificationType200Response) GetIsIncludedInPlan() bool`

GetIsIncludedInPlan returns the IsIncludedInPlan field if non-nil, zero value otherwise.

### GetIsIncludedInPlanOk

`func (o *GetNotificationsV1NotificationType200Response) GetIsIncludedInPlanOk() (*bool, bool)`

GetIsIncludedInPlanOk returns a tuple with the IsIncludedInPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncludedInPlan

`func (o *GetNotificationsV1NotificationType200Response) SetIsIncludedInPlan(v bool)`

SetIsIncludedInPlan sets IsIncludedInPlan field to given value.


### GetSeverity

`func (o *GetNotificationsV1NotificationType200Response) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetNotificationsV1NotificationType200Response) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetNotificationsV1NotificationType200Response) SetSeverity(v string)`

SetSeverity sets Severity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


