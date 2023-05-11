# GetNotificationsV1Integration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**NotificationsV1IntegrationMetadata**](NotificationsV1IntegrationMetadata.md) |  | [optional] 
**DisplayName** | **string** | A human readable name for the particular integration  | 
**Description** | Pointer to **string** | A human readable description for the particular integration  | [optional] 
**Target** | [**NotificationsV1IntegrationTarget**](NotificationsV1IntegrationTarget.md) |  | 

## Methods

### NewGetNotificationsV1Integration200Response

`func NewGetNotificationsV1Integration200Response(apiVersion string, kind string, id string, displayName string, target NotificationsV1IntegrationTarget, ) *GetNotificationsV1Integration200Response`

NewGetNotificationsV1Integration200Response instantiates a new GetNotificationsV1Integration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsV1Integration200ResponseWithDefaults

`func NewGetNotificationsV1Integration200ResponseWithDefaults() *GetNotificationsV1Integration200Response`

NewGetNotificationsV1Integration200ResponseWithDefaults instantiates a new GetNotificationsV1Integration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetNotificationsV1Integration200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetNotificationsV1Integration200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetNotificationsV1Integration200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetNotificationsV1Integration200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetNotificationsV1Integration200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetNotificationsV1Integration200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetNotificationsV1Integration200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNotificationsV1Integration200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNotificationsV1Integration200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetNotificationsV1Integration200Response) GetMetadata() NotificationsV1IntegrationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetNotificationsV1Integration200Response) GetMetadataOk() (*NotificationsV1IntegrationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetNotificationsV1Integration200Response) SetMetadata(v NotificationsV1IntegrationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetNotificationsV1Integration200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNotificationsV1Integration200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNotificationsV1Integration200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNotificationsV1Integration200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *GetNotificationsV1Integration200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNotificationsV1Integration200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNotificationsV1Integration200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNotificationsV1Integration200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTarget

`func (o *GetNotificationsV1Integration200Response) GetTarget() NotificationsV1IntegrationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetNotificationsV1Integration200Response) GetTargetOk() (*NotificationsV1IntegrationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetNotificationsV1Integration200Response) SetTarget(v NotificationsV1IntegrationTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


