# CreateIamV2IdentityPoolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2IdentityPoolMetadata**](IamV2IdentityPoolMetadata.md) |  | [optional] 
**DisplayName** | **string** | The name of the &#x60;IdentityPool&#x60;. | 
**Description** | **string** | A description of how this &#x60;IdentityPool&#x60; is used | 
**IdentityClaim** | **string** | The JSON Web Token (JWT) claim to extract the authenticating identity to Confluent resources from (see [Registered Claim Names](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1) for more details). This appears in the audit log records, showing, for example, that \&quot;identity Z used identity pool X to access topic A\&quot;. | 
**Filter** | **string** | A filter expression in [Supported Common Expression Language (CEL)](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#supported-common-expression-language-cel-filters) that specifies which identities can authenticate using your identity pool (see [Set identity pool filters](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#set-identity-pool-filters) for more details). | 
**Principal** | Pointer to **string** | Represents the federated identity associated with this pool. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the identity pool | [optional] [readonly] 

## Methods

### NewCreateIamV2IdentityPoolRequest

`func NewCreateIamV2IdentityPoolRequest(displayName string, description string, identityClaim string, filter string, ) *CreateIamV2IdentityPoolRequest`

NewCreateIamV2IdentityPoolRequest instantiates a new CreateIamV2IdentityPoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2IdentityPoolRequestWithDefaults

`func NewCreateIamV2IdentityPoolRequestWithDefaults() *CreateIamV2IdentityPoolRequest`

NewCreateIamV2IdentityPoolRequestWithDefaults instantiates a new CreateIamV2IdentityPoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2IdentityPoolRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2IdentityPoolRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2IdentityPoolRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2IdentityPoolRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2IdentityPoolRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2IdentityPoolRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2IdentityPoolRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2IdentityPoolRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2IdentityPoolRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2IdentityPoolRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2IdentityPoolRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2IdentityPoolRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2IdentityPoolRequest) GetMetadata() IamV2IdentityPoolMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2IdentityPoolRequest) GetMetadataOk() (*IamV2IdentityPoolMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2IdentityPoolRequest) SetMetadata(v IamV2IdentityPoolMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2IdentityPoolRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateIamV2IdentityPoolRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateIamV2IdentityPoolRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateIamV2IdentityPoolRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CreateIamV2IdentityPoolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIamV2IdentityPoolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIamV2IdentityPoolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIdentityClaim

`func (o *CreateIamV2IdentityPoolRequest) GetIdentityClaim() string`

GetIdentityClaim returns the IdentityClaim field if non-nil, zero value otherwise.

### GetIdentityClaimOk

`func (o *CreateIamV2IdentityPoolRequest) GetIdentityClaimOk() (*string, bool)`

GetIdentityClaimOk returns a tuple with the IdentityClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaim

`func (o *CreateIamV2IdentityPoolRequest) SetIdentityClaim(v string)`

SetIdentityClaim sets IdentityClaim field to given value.


### GetFilter

`func (o *CreateIamV2IdentityPoolRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateIamV2IdentityPoolRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateIamV2IdentityPoolRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetPrincipal

`func (o *CreateIamV2IdentityPoolRequest) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CreateIamV2IdentityPoolRequest) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CreateIamV2IdentityPoolRequest) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *CreateIamV2IdentityPoolRequest) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetState

`func (o *CreateIamV2IdentityPoolRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateIamV2IdentityPoolRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateIamV2IdentityPoolRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateIamV2IdentityPoolRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


