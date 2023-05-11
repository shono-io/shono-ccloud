# GetIamV2IdentityPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**IamV2IdentityPoolMetadata**](IamV2IdentityPoolMetadata.md) |  | [optional] 
**DisplayName** | **string** | The name of the &#x60;IdentityPool&#x60;. | 
**Description** | **string** | A description of how this &#x60;IdentityPool&#x60; is used | 
**IdentityClaim** | **string** | The JSON Web Token (JWT) claim to extract the authenticating identity to Confluent resources from (see [Registered Claim Names](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1) for more details). This appears in the audit log records, showing, for example, that \&quot;identity Z used identity pool X to access topic A\&quot;. | 
**Filter** | **string** | A filter expression in [Supported Common Expression Language (CEL)](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#supported-common-expression-language-cel-filters) that specifies which identities can authenticate using your identity pool (see [Set identity pool filters](https://docs.confluent.io/cloud/current/access-management/authenticate/oauth/identity-pools.html#set-identity-pool-filters) for more details). | 
**Principal** | **string** | Represents the federated identity associated with this pool. | [readonly] 
**State** | **string** | The current state of the identity pool | [readonly] 

## Methods

### NewGetIamV2IdentityPool200Response

`func NewGetIamV2IdentityPool200Response(apiVersion string, kind string, id string, displayName string, description string, identityClaim string, filter string, principal string, state string, ) *GetIamV2IdentityPool200Response`

NewGetIamV2IdentityPool200Response instantiates a new GetIamV2IdentityPool200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIamV2IdentityPool200ResponseWithDefaults

`func NewGetIamV2IdentityPool200ResponseWithDefaults() *GetIamV2IdentityPool200Response`

NewGetIamV2IdentityPool200ResponseWithDefaults instantiates a new GetIamV2IdentityPool200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetIamV2IdentityPool200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetIamV2IdentityPool200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetIamV2IdentityPool200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetIamV2IdentityPool200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetIamV2IdentityPool200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetIamV2IdentityPool200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetIamV2IdentityPool200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIamV2IdentityPool200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIamV2IdentityPool200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetIamV2IdentityPool200Response) GetMetadata() IamV2IdentityPoolMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIamV2IdentityPool200Response) GetMetadataOk() (*IamV2IdentityPoolMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIamV2IdentityPool200Response) SetMetadata(v IamV2IdentityPoolMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIamV2IdentityPool200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetIamV2IdentityPool200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetIamV2IdentityPool200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetIamV2IdentityPool200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *GetIamV2IdentityPool200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIamV2IdentityPool200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIamV2IdentityPool200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIdentityClaim

`func (o *GetIamV2IdentityPool200Response) GetIdentityClaim() string`

GetIdentityClaim returns the IdentityClaim field if non-nil, zero value otherwise.

### GetIdentityClaimOk

`func (o *GetIamV2IdentityPool200Response) GetIdentityClaimOk() (*string, bool)`

GetIdentityClaimOk returns a tuple with the IdentityClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityClaim

`func (o *GetIamV2IdentityPool200Response) SetIdentityClaim(v string)`

SetIdentityClaim sets IdentityClaim field to given value.


### GetFilter

`func (o *GetIamV2IdentityPool200Response) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetIamV2IdentityPool200Response) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetIamV2IdentityPool200Response) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetPrincipal

`func (o *GetIamV2IdentityPool200Response) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *GetIamV2IdentityPool200Response) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *GetIamV2IdentityPool200Response) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetState

`func (o *GetIamV2IdentityPool200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetIamV2IdentityPool200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetIamV2IdentityPool200Response) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


