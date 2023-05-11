# CreateIamV2RoleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2RoleBindingMetadata**](IamV2RoleBindingMetadata.md) |  | [optional] 
**Principal** | **string** | The principal User or Group to bind the role to | 
**RoleName** | **string** | The name of the role to bind to the principal | 
**CrnPattern** | **string** | A CRN that specifies the scope and resource patterns necessary for the role to bind | 

## Methods

### NewCreateIamV2RoleBindingRequest

`func NewCreateIamV2RoleBindingRequest(principal string, roleName string, crnPattern string, ) *CreateIamV2RoleBindingRequest`

NewCreateIamV2RoleBindingRequest instantiates a new CreateIamV2RoleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2RoleBindingRequestWithDefaults

`func NewCreateIamV2RoleBindingRequestWithDefaults() *CreateIamV2RoleBindingRequest`

NewCreateIamV2RoleBindingRequestWithDefaults instantiates a new CreateIamV2RoleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2RoleBindingRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2RoleBindingRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2RoleBindingRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2RoleBindingRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2RoleBindingRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2RoleBindingRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2RoleBindingRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2RoleBindingRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2RoleBindingRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2RoleBindingRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2RoleBindingRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2RoleBindingRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2RoleBindingRequest) GetMetadata() IamV2RoleBindingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2RoleBindingRequest) GetMetadataOk() (*IamV2RoleBindingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2RoleBindingRequest) SetMetadata(v IamV2RoleBindingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2RoleBindingRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrincipal

`func (o *CreateIamV2RoleBindingRequest) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CreateIamV2RoleBindingRequest) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CreateIamV2RoleBindingRequest) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetRoleName

`func (o *CreateIamV2RoleBindingRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *CreateIamV2RoleBindingRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *CreateIamV2RoleBindingRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetCrnPattern

`func (o *CreateIamV2RoleBindingRequest) GetCrnPattern() string`

GetCrnPattern returns the CrnPattern field if non-nil, zero value otherwise.

### GetCrnPatternOk

`func (o *CreateIamV2RoleBindingRequest) GetCrnPatternOk() (*string, bool)`

GetCrnPatternOk returns a tuple with the CrnPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrnPattern

`func (o *CreateIamV2RoleBindingRequest) SetCrnPattern(v string)`

SetCrnPattern sets CrnPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


