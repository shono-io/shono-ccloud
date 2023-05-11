# OrgV2OrganizationListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**OrgV2OrganizationMetadata**](OrgV2OrganizationMetadata.md) |  | 
**DisplayName** | **string** | A human-readable name for the Organization | 

## Methods

### NewOrgV2OrganizationListDataInner

`func NewOrgV2OrganizationListDataInner(id string, metadata OrgV2OrganizationMetadata, displayName string, ) *OrgV2OrganizationListDataInner`

NewOrgV2OrganizationListDataInner instantiates a new OrgV2OrganizationListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgV2OrganizationListDataInnerWithDefaults

`func NewOrgV2OrganizationListDataInnerWithDefaults() *OrgV2OrganizationListDataInner`

NewOrgV2OrganizationListDataInnerWithDefaults instantiates a new OrgV2OrganizationListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrgV2OrganizationListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrgV2OrganizationListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrgV2OrganizationListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OrgV2OrganizationListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *OrgV2OrganizationListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrgV2OrganizationListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrgV2OrganizationListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrgV2OrganizationListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *OrgV2OrganizationListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgV2OrganizationListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgV2OrganizationListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *OrgV2OrganizationListDataInner) GetMetadata() OrgV2OrganizationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrgV2OrganizationListDataInner) GetMetadataOk() (*OrgV2OrganizationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrgV2OrganizationListDataInner) SetMetadata(v OrgV2OrganizationMetadata)`

SetMetadata sets Metadata field to given value.


### GetDisplayName

`func (o *OrgV2OrganizationListDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OrgV2OrganizationListDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OrgV2OrganizationListDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


