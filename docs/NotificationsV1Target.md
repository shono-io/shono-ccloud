# NotificationsV1Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Integration Type | 
**WebhookUrl** | **string** | MS Teams Webhook URL for the particular team channel | 
**RoleName** | **string** | name of the role | 
**User** | **string** | ID of the user | 
**Url** | **string** | URL endpoint for the webhook | 

## Methods

### NewNotificationsV1Target

`func NewNotificationsV1Target(kind string, webhookUrl string, roleName string, user string, url string, ) *NotificationsV1Target`

NewNotificationsV1Target instantiates a new NotificationsV1Target object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1TargetWithDefaults

`func NewNotificationsV1TargetWithDefaults() *NotificationsV1Target`

NewNotificationsV1TargetWithDefaults instantiates a new NotificationsV1Target object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NotificationsV1Target) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1Target) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1Target) SetKind(v string)`

SetKind sets Kind field to given value.


### GetWebhookUrl

`func (o *NotificationsV1Target) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *NotificationsV1Target) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *NotificationsV1Target) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.


### GetRoleName

`func (o *NotificationsV1Target) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *NotificationsV1Target) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *NotificationsV1Target) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetUser

`func (o *NotificationsV1Target) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationsV1Target) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationsV1Target) SetUser(v string)`

SetUser sets User field to given value.


### GetUrl

`func (o *NotificationsV1Target) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationsV1Target) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationsV1Target) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


