# PartnerSignupRequestOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | **string** | The name of the organization | 
**SsoUrl** | Pointer to **string** | The login URL for the customer to access Confluent Cloud | [optional] [readonly] 
**SsoConfig** | [**PartnerV2OrganizationSsoConfig**](PartnerV2OrganizationSsoConfig.md) |  | 

## Methods

### NewPartnerSignupRequestOrganization

`func NewPartnerSignupRequestOrganization(name string, ssoConfig PartnerV2OrganizationSsoConfig, ) *PartnerSignupRequestOrganization`

NewPartnerSignupRequestOrganization instantiates a new PartnerSignupRequestOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestOrganizationWithDefaults

`func NewPartnerSignupRequestOrganizationWithDefaults() *PartnerSignupRequestOrganization`

NewPartnerSignupRequestOrganizationWithDefaults instantiates a new PartnerSignupRequestOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerSignupRequestOrganization) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerSignupRequestOrganization) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerSignupRequestOrganization) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerSignupRequestOrganization) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerSignupRequestOrganization) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerSignupRequestOrganization) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerSignupRequestOrganization) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerSignupRequestOrganization) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerSignupRequestOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerSignupRequestOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerSignupRequestOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartnerSignupRequestOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PartnerSignupRequestOrganization) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerSignupRequestOrganization) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerSignupRequestOrganization) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartnerSignupRequestOrganization) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *PartnerSignupRequestOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerSignupRequestOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerSignupRequestOrganization) SetName(v string)`

SetName sets Name field to given value.


### GetSsoUrl

`func (o *PartnerSignupRequestOrganization) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *PartnerSignupRequestOrganization) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *PartnerSignupRequestOrganization) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *PartnerSignupRequestOrganization) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetSsoConfig

`func (o *PartnerSignupRequestOrganization) GetSsoConfig() PartnerV2OrganizationSsoConfig`

GetSsoConfig returns the SsoConfig field if non-nil, zero value otherwise.

### GetSsoConfigOk

`func (o *PartnerSignupRequestOrganization) GetSsoConfigOk() (*PartnerV2OrganizationSsoConfig, bool)`

GetSsoConfigOk returns a tuple with the SsoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoConfig

`func (o *PartnerSignupRequestOrganization) SetSsoConfig(v PartnerV2OrganizationSsoConfig)`

SetSsoConfig sets SsoConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


