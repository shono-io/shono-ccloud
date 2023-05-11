# IamV2IdentityPoolListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**IamV2IdentityPoolMetadata**](IamV2IdentityPoolMetadata.md) |  | 
**DisplayName** | **string** | The name of the &#x60;IdentityPool&#x60;. | 
**Description** | **string** | A description of how this &#x60;IdentityPool&#x60; is used | 
**IdentityClaim** | **string** | The JSON Web Token (JWT) claim to extract the authenticating identity to Confluent resources from (see [Registered Claim Names](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1) for more details). This appears in the audit log records, showing, for example, that \&quot;identity Z used identity pool X to access topic A\&quot;. | 
**Filter** | **string** | A filter expression in [Supported Common Expression Language (CEL)](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#supported-common-expression-language-cel-filters) that specifies which identities can authenticate using your identity pool (see [Set identity pool filters](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#set-identity-pool-filters) for more details). | 
**Principal** | **string** | Represents the federated identity associated with this pool. | [readonly] 
**State** | **string** | The current state of the identity pool | [readonly] 

## Methods

### NewIamV2IdentityPoolListDataInner

`func NewIamV2IdentityPoolListDataInner(id string, metadata IamV2IdentityPoolMetadata, displayName string, description string, identityClaim string, filter string, principal string, state string, ) *IamV2IdentityPoolListDataInner`

NewIamV2IdentityPoolListDataInner instantiates a new IamV2IdentityPoolListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityPoolListDataInnerWithDefaults

`func NewIamV2IdentityPoolListDataInnerWithDefaults() *IamV2IdentityPoolListDataInner`

NewIamV2IdentityPoolListDataInnerWithDefaults instantiates a new IamV2IdentityPoolListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityPoolListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityPoolListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityPoolListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IdentityPoolListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IdentityPoolListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityPoolListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityPoolListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IdentityPoolListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IdentityPoolListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IdentityPoolListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IdentityPoolListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IamV2IdentityPoolListDataInner) GetMetadata() IamV2IdentityPoolMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityPoolListDataInner) GetMetadataOk() (*IamV2IdentityPoolMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityPoolListDataInner) SetMetadata(v IamV2IdentityPoolMetadata)`

SetMetadata sets Metadata field to given value.


### GetDisplayName

`func (o *IamV2IdentityPoolListDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2IdentityPoolListDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2IdentityPoolListDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *IamV2IdentityPoolListDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2IdentityPoolListDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2IdentityPoolListDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIdentityClaim

`func (o *IamV2IdentityPoolListDataInner) GetIdentityClaim() string`

GetIdentityClaim returns the IdentityClaim field if non-nil, zero value otherwise.

### GetIdentityClaimOk

`func (o *IamV2IdentityPoolListDataInner) GetIdentityClaimOk() (*string, bool)`

GetIdentityClaimOk returns a tuple with the IdentityClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaim

`func (o *IamV2IdentityPoolListDataInner) SetIdentityClaim(v string)`

SetIdentityClaim sets IdentityClaim field to given value.


### GetFilter

`func (o *IamV2IdentityPoolListDataInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IamV2IdentityPoolListDataInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IamV2IdentityPoolListDataInner) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetPrincipal

`func (o *IamV2IdentityPoolListDataInner) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *IamV2IdentityPoolListDataInner) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *IamV2IdentityPoolListDataInner) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetState

`func (o *IamV2IdentityPoolListDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2IdentityPoolListDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2IdentityPoolListDataInner) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


