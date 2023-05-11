# NotificationsV1IntegrationListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**NotificationsV1IntegrationMetadata**](NotificationsV1IntegrationMetadata.md) |  | 
**DisplayName** | **string** | A human readable name for the particular integration  | 
**Description** | Pointer to **string** | A human readable description for the particular integration  | [optional] 
**Target** | [**NotificationsV1IntegrationTarget**](NotificationsV1IntegrationTarget.md) |  | 

## Methods

### NewNotificationsV1IntegrationListDataInner

`func NewNotificationsV1IntegrationListDataInner(id string, metadata NotificationsV1IntegrationMetadata, displayName string, target NotificationsV1IntegrationTarget, ) *NotificationsV1IntegrationListDataInner`

NewNotificationsV1IntegrationListDataInner instantiates a new NotificationsV1IntegrationListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1IntegrationListDataInnerWithDefaults

`func NewNotificationsV1IntegrationListDataInnerWithDefaults() *NotificationsV1IntegrationListDataInner`

NewNotificationsV1IntegrationListDataInnerWithDefaults instantiates a new NotificationsV1IntegrationListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1IntegrationListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1IntegrationListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1IntegrationListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NotificationsV1IntegrationListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NotificationsV1IntegrationListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1IntegrationListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1IntegrationListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NotificationsV1IntegrationListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *NotificationsV1IntegrationListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsV1IntegrationListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsV1IntegrationListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *NotificationsV1IntegrationListDataInner) GetMetadata() NotificationsV1IntegrationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1IntegrationListDataInner) GetMetadataOk() (*NotificationsV1IntegrationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1IntegrationListDataInner) SetMetadata(v NotificationsV1IntegrationMetadata)`

SetMetadata sets Metadata field to given value.


### GetDisplayName

`func (o *NotificationsV1IntegrationListDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotificationsV1IntegrationListDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NotificationsV1IntegrationListDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *NotificationsV1IntegrationListDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationsV1IntegrationListDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationsV1IntegrationListDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationsV1IntegrationListDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTarget

`func (o *NotificationsV1IntegrationListDataInner) GetTarget() NotificationsV1IntegrationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NotificationsV1IntegrationListDataInner) GetTargetOk() (*NotificationsV1IntegrationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NotificationsV1IntegrationListDataInner) SetTarget(v NotificationsV1IntegrationTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


