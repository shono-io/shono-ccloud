# IamV2IdentityProviderListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**IamV2IdentityProviderMetadata**](IamV2IdentityProviderMetadata.md) |  | 
**DisplayName** | **string** | The name of the Provider. | 
**Description** | **string** | A description for your provider | 
**State** | **string** | The current state of the provider | [readonly] 
**Issuer** | **string** | A publicly reachable issuer URI for the Identity Provider. The unique issuer URI string represents the entity for issuing tokens. | 
**JwksUri** | **string** | A publicly reachable JSON Web Key Set (JWKS) URI for the Identity Provider. A JSON Web Key Set (JWKS) provides a set of keys containing the public keys used to verify any JSON Web Token (JWT) issued by your OAuth 2.0 identity provider. | 
**Keys** | Pointer to [**[]IamV2JwksObject**](IamV2JwksObject.md) | The JWKS provided by the Provider. We only express the &#x60;kid&#x60; and &#x60;alg&#x60; for each key set | [optional] [readonly] 

## Methods

### NewIamV2IdentityProviderListDataInner

`func NewIamV2IdentityProviderListDataInner(id string, metadata IamV2IdentityProviderMetadata, displayName string, description string, state string, issuer string, jwksUri string, ) *IamV2IdentityProviderListDataInner`

NewIamV2IdentityProviderListDataInner instantiates a new IamV2IdentityProviderListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityProviderListDataInnerWithDefaults

`func NewIamV2IdentityProviderListDataInnerWithDefaults() *IamV2IdentityProviderListDataInner`

NewIamV2IdentityProviderListDataInnerWithDefaults instantiates a new IamV2IdentityProviderListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityProviderListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityProviderListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityProviderListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IdentityProviderListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IdentityProviderListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityProviderListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityProviderListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IdentityProviderListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2IdentityProviderListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2IdentityProviderListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2IdentityProviderListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IamV2IdentityProviderListDataInner) GetMetadata() IamV2IdentityProviderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityProviderListDataInner) GetMetadataOk() (*IamV2IdentityProviderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityProviderListDataInner) SetMetadata(v IamV2IdentityProviderMetadata)`

SetMetadata sets Metadata field to given value.


### GetDisplayName

`func (o *IamV2IdentityProviderListDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2IdentityProviderListDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2IdentityProviderListDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *IamV2IdentityProviderListDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2IdentityProviderListDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2IdentityProviderListDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetState

`func (o *IamV2IdentityProviderListDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IamV2IdentityProviderListDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IamV2IdentityProviderListDataInner) SetState(v string)`

SetState sets State field to given value.


### GetIssuer

`func (o *IamV2IdentityProviderListDataInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IamV2IdentityProviderListDataInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IamV2IdentityProviderListDataInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUri

`func (o *IamV2IdentityProviderListDataInner) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IamV2IdentityProviderListDataInner) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IamV2IdentityProviderListDataInner) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.


### GetKeys

`func (o *IamV2IdentityProviderListDataInner) GetKeys() []IamV2JwksObject`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IamV2IdentityProviderListDataInner) GetKeysOk() (*[]IamV2JwksObject, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IamV2IdentityProviderListDataInner) SetKeys(v []IamV2JwksObject)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IamV2IdentityProviderListDataInner) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


