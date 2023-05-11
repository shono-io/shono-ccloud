# CreateIamV2IdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2IdentityProviderMetadata**](IamV2IdentityProviderMetadata.md) |  | [optional] 
**DisplayName** | **string** | The name of the Provider. | 
**Description** | **string** | A description for your provider | 
**State** | Pointer to **string** | The current state of the provider | [optional] [readonly] 
**Issuer** | **string** | A publicly reachable issuer URI for the Identity Provider. The unique issuer URI string represents the entity for issuing tokens. | 
**JwksUri** | **string** | A publicly reachable JSON Web Key Set (JWKS) URI for the Identity Provider. A JSON Web Key Set (JWKS) provides a set of keys containing the public keys used to verify any JSON Web Token (JWT) issued by your OAuth 2.0 identity provider. | 
**Keys** | Pointer to [**[]IamV2JwksObject**](IamV2JwksObject.md) | The JWKS provided by the Provider. We only express the &#x60;kid&#x60; and &#x60;alg&#x60; for each key set | [optional] [readonly] 

## Methods

### NewCreateIamV2IdentityProviderRequest

`func NewCreateIamV2IdentityProviderRequest(displayName string, description string, issuer string, jwksUri string, ) *CreateIamV2IdentityProviderRequest`

NewCreateIamV2IdentityProviderRequest instantiates a new CreateIamV2IdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2IdentityProviderRequestWithDefaults

`func NewCreateIamV2IdentityProviderRequestWithDefaults() *CreateIamV2IdentityProviderRequest`

NewCreateIamV2IdentityProviderRequestWithDefaults instantiates a new CreateIamV2IdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2IdentityProviderRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2IdentityProviderRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2IdentityProviderRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2IdentityProviderRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2IdentityProviderRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2IdentityProviderRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2IdentityProviderRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2IdentityProviderRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2IdentityProviderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2IdentityProviderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2IdentityProviderRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2IdentityProviderRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2IdentityProviderRequest) GetMetadata() IamV2IdentityProviderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2IdentityProviderRequest) GetMetadataOk() (*IamV2IdentityProviderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2IdentityProviderRequest) SetMetadata(v IamV2IdentityProviderMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2IdentityProviderRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateIamV2IdentityProviderRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateIamV2IdentityProviderRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateIamV2IdentityProviderRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CreateIamV2IdentityProviderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIamV2IdentityProviderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIamV2IdentityProviderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetState

`func (o *CreateIamV2IdentityProviderRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateIamV2IdentityProviderRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateIamV2IdentityProviderRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateIamV2IdentityProviderRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIssuer

`func (o *CreateIamV2IdentityProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateIamV2IdentityProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateIamV2IdentityProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUri

`func (o *CreateIamV2IdentityProviderRequest) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *CreateIamV2IdentityProviderRequest) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *CreateIamV2IdentityProviderRequest) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.


### GetKeys

`func (o *CreateIamV2IdentityProviderRequest) GetKeys() []IamV2JwksObject`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CreateIamV2IdentityProviderRequest) GetKeysOk() (*[]IamV2JwksObject, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CreateIamV2IdentityProviderRequest) SetKeys(v []IamV2JwksObject)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CreateIamV2IdentityProviderRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


