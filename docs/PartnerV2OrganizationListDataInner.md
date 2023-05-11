# PartnerV2OrganizationListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**Name** | Pointer to **string** | The name of the organization | [optional] 
**SsoUrl** | Pointer to **string** | The login URL for the customer to access Confluent Cloud | [optional] [readonly] 
**SsoConfig** | Pointer to [**PartnerV2OrganizationSsoConfig**](PartnerV2OrganizationSsoConfig.md) |  | [optional] 

## Methods

### NewPartnerV2OrganizationListDataInner

`func NewPartnerV2OrganizationListDataInner(id string, metadata ObjectMeta, ) *PartnerV2OrganizationListDataInner`

NewPartnerV2OrganizationListDataInner instantiates a new PartnerV2OrganizationListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerV2OrganizationListDataInnerWithDefaults

`func NewPartnerV2OrganizationListDataInnerWithDefaults() *PartnerV2OrganizationListDataInner`

NewPartnerV2OrganizationListDataInnerWithDefaults instantiates a new PartnerV2OrganizationListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerV2OrganizationListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerV2OrganizationListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerV2OrganizationListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerV2OrganizationListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerV2OrganizationListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerV2OrganizationListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerV2OrganizationListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerV2OrganizationListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerV2OrganizationListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerV2OrganizationListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerV2OrganizationListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *PartnerV2OrganizationListDataInner) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerV2OrganizationListDataInner) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerV2OrganizationListDataInner) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *PartnerV2OrganizationListDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerV2OrganizationListDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerV2OrganizationListDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerV2OrganizationListDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSsoUrl

`func (o *PartnerV2OrganizationListDataInner) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *PartnerV2OrganizationListDataInner) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *PartnerV2OrganizationListDataInner) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *PartnerV2OrganizationListDataInner) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetSsoConfig

`func (o *PartnerV2OrganizationListDataInner) GetSsoConfig() PartnerV2OrganizationSsoConfig`

GetSsoConfig returns the SsoConfig field if non-nil, zero value otherwise.

### GetSsoConfigOk

`func (o *PartnerV2OrganizationListDataInner) GetSsoConfigOk() (*PartnerV2OrganizationSsoConfig, bool)`

GetSsoConfigOk returns a tuple with the SsoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoConfig

`func (o *PartnerV2OrganizationListDataInner) SetSsoConfig(v PartnerV2OrganizationSsoConfig)`

SetSsoConfig sets SsoConfig field to given value.

### HasSsoConfig

`func (o *PartnerV2OrganizationListDataInner) HasSsoConfig() bool`

HasSsoConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


