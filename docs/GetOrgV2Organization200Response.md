# GetOrgV2Organization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**OrgV2OrganizationMetadata**](OrgV2OrganizationMetadata.md) |  | [optional] 
**DisplayName** | **string** | A human-readable name for the Organization | 

## Methods

### NewGetOrgV2Organization200Response

`func NewGetOrgV2Organization200Response(apiVersion string, kind string, id string, displayName string, ) *GetOrgV2Organization200Response`

NewGetOrgV2Organization200Response instantiates a new GetOrgV2Organization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgV2Organization200ResponseWithDefaults

`func NewGetOrgV2Organization200ResponseWithDefaults() *GetOrgV2Organization200Response`

NewGetOrgV2Organization200ResponseWithDefaults instantiates a new GetOrgV2Organization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetOrgV2Organization200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetOrgV2Organization200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetOrgV2Organization200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetOrgV2Organization200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetOrgV2Organization200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetOrgV2Organization200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetOrgV2Organization200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrgV2Organization200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrgV2Organization200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetOrgV2Organization200Response) GetMetadata() OrgV2OrganizationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetOrgV2Organization200Response) GetMetadataOk() (*OrgV2OrganizationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetOrgV2Organization200Response) SetMetadata(v OrgV2OrganizationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetOrgV2Organization200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetOrgV2Organization200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetOrgV2Organization200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetOrgV2Organization200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


