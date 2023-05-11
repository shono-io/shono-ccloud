# CreateOrgV2EnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**OrgV2EnvironmentMetadata**](OrgV2EnvironmentMetadata.md) |  | [optional] 
**DisplayName** | **string** | A human-readable name for the Environment | 

## Methods

### NewCreateOrgV2EnvironmentRequest

`func NewCreateOrgV2EnvironmentRequest(displayName string, ) *CreateOrgV2EnvironmentRequest`

NewCreateOrgV2EnvironmentRequest instantiates a new CreateOrgV2EnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgV2EnvironmentRequestWithDefaults

`func NewCreateOrgV2EnvironmentRequestWithDefaults() *CreateOrgV2EnvironmentRequest`

NewCreateOrgV2EnvironmentRequestWithDefaults instantiates a new CreateOrgV2EnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateOrgV2EnvironmentRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateOrgV2EnvironmentRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateOrgV2EnvironmentRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateOrgV2EnvironmentRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateOrgV2EnvironmentRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateOrgV2EnvironmentRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateOrgV2EnvironmentRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateOrgV2EnvironmentRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateOrgV2EnvironmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrgV2EnvironmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrgV2EnvironmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrgV2EnvironmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateOrgV2EnvironmentRequest) GetMetadata() OrgV2EnvironmentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateOrgV2EnvironmentRequest) GetMetadataOk() (*OrgV2EnvironmentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateOrgV2EnvironmentRequest) SetMetadata(v OrgV2EnvironmentMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateOrgV2EnvironmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateOrgV2EnvironmentRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateOrgV2EnvironmentRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateOrgV2EnvironmentRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


