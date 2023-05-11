# GetIamV2RoleBinding200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**IamV2RoleBindingMetadata**](IamV2RoleBindingMetadata.md) |  | [optional] 
**Principal** | **string** | The principal User or Group to bind the role to | 
**RoleName** | **string** | The name of the role to bind to the principal | 
**CrnPattern** | **string** | A CRN that specifies the scope and resource patterns necessary for the role to bind | 

## Methods

### NewGetIamV2RoleBinding200Response

`func NewGetIamV2RoleBinding200Response(apiVersion string, kind string, id string, principal string, roleName string, crnPattern string, ) *GetIamV2RoleBinding200Response`

NewGetIamV2RoleBinding200Response instantiates a new GetIamV2RoleBinding200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIamV2RoleBinding200ResponseWithDefaults

`func NewGetIamV2RoleBinding200ResponseWithDefaults() *GetIamV2RoleBinding200Response`

NewGetIamV2RoleBinding200ResponseWithDefaults instantiates a new GetIamV2RoleBinding200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetIamV2RoleBinding200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetIamV2RoleBinding200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetIamV2RoleBinding200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetIamV2RoleBinding200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetIamV2RoleBinding200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetIamV2RoleBinding200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetIamV2RoleBinding200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIamV2RoleBinding200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIamV2RoleBinding200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetIamV2RoleBinding200Response) GetMetadata() IamV2RoleBindingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIamV2RoleBinding200Response) GetMetadataOk() (*IamV2RoleBindingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIamV2RoleBinding200Response) SetMetadata(v IamV2RoleBindingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIamV2RoleBinding200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPrincipal

`func (o *GetIamV2RoleBinding200Response) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *GetIamV2RoleBinding200Response) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *GetIamV2RoleBinding200Response) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetRoleName

`func (o *GetIamV2RoleBinding200Response) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *GetIamV2RoleBinding200Response) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *GetIamV2RoleBinding200Response) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetCrnPattern

`func (o *GetIamV2RoleBinding200Response) GetCrnPattern() string`

GetCrnPattern returns the CrnPattern field if non-nil, zero value otherwise.

### GetCrnPatternOk

`func (o *GetIamV2RoleBinding200Response) GetCrnPatternOk() (*string, bool)`

GetCrnPatternOk returns a tuple with the CrnPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrnPattern

`func (o *GetIamV2RoleBinding200Response) SetCrnPattern(v string)`

SetCrnPattern sets CrnPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


