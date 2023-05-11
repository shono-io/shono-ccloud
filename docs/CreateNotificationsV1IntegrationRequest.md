# CreateNotificationsV1IntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**NotificationsV1IntegrationMetadata**](NotificationsV1IntegrationMetadata.md) |  | [optional] 
**DisplayName** | **string** | A human readable name for the particular integration  | 
**Description** | Pointer to **string** | A human readable description for the particular integration  | [optional] 
**Target** | [**NotificationsV1IntegrationTarget**](NotificationsV1IntegrationTarget.md) |  | 

## Methods

### NewCreateNotificationsV1IntegrationRequest

`func NewCreateNotificationsV1IntegrationRequest(displayName string, target NotificationsV1IntegrationTarget, ) *CreateNotificationsV1IntegrationRequest`

NewCreateNotificationsV1IntegrationRequest instantiates a new CreateNotificationsV1IntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationsV1IntegrationRequestWithDefaults

`func NewCreateNotificationsV1IntegrationRequestWithDefaults() *CreateNotificationsV1IntegrationRequest`

NewCreateNotificationsV1IntegrationRequestWithDefaults instantiates a new CreateNotificationsV1IntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateNotificationsV1IntegrationRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateNotificationsV1IntegrationRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateNotificationsV1IntegrationRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateNotificationsV1IntegrationRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateNotificationsV1IntegrationRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateNotificationsV1IntegrationRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateNotificationsV1IntegrationRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateNotificationsV1IntegrationRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateNotificationsV1IntegrationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNotificationsV1IntegrationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNotificationsV1IntegrationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNotificationsV1IntegrationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNotificationsV1IntegrationRequest) GetMetadata() NotificationsV1IntegrationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNotificationsV1IntegrationRequest) GetMetadataOk() (*NotificationsV1IntegrationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNotificationsV1IntegrationRequest) SetMetadata(v NotificationsV1IntegrationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNotificationsV1IntegrationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateNotificationsV1IntegrationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateNotificationsV1IntegrationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateNotificationsV1IntegrationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CreateNotificationsV1IntegrationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNotificationsV1IntegrationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNotificationsV1IntegrationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNotificationsV1IntegrationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTarget

`func (o *CreateNotificationsV1IntegrationRequest) GetTarget() NotificationsV1IntegrationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateNotificationsV1IntegrationRequest) GetTargetOk() (*NotificationsV1IntegrationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateNotificationsV1IntegrationRequest) SetTarget(v NotificationsV1IntegrationTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


