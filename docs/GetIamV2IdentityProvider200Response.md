# GetIamV2IdentityProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**IamV2IdentityProviderMetadata**](IamV2IdentityProviderMetadata.md) |  | [optional] 
**DisplayName** | **string** | The name of the Provider. | 
**Description** | **string** | A description for your provider | 
**State** | **string** | The current state of the provider | [readonly] 
**Issuer** | **string** | A publicly reachable issuer URI for the Identity Provider. The unique issuer URI string represents the entity for issuing tokens. | 
**JwksUri** | **string** | A publicly reachable JSON Web Key Set (JWKS) URI for the Identity Provider. A JSON Web Key Set (JWKS) provides a set of keys containing the public keys used to verify any JSON Web Token (JWT) issued by your OAuth 2.0 identity provider. | 
**Keys** | Pointer to [**[]IamV2JwksObject**](IamV2JwksObject.md) | The JWKS provided by the Provider. We only express the &#x60;kid&#x60; and &#x60;alg&#x60; for each key set | [optional] [readonly] 

## Methods

### NewGetIamV2IdentityProvider200Response

`func NewGetIamV2IdentityProvider200Response(apiVersion string, kind string, id string, displayName string, description string, state string, issuer string, jwksUri string, ) *GetIamV2IdentityProvider200Response`

NewGetIamV2IdentityProvider200Response instantiates a new GetIamV2IdentityProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIamV2IdentityProvider200ResponseWithDefaults

`func NewGetIamV2IdentityProvider200ResponseWithDefaults() *GetIamV2IdentityProvider200Response`

NewGetIamV2IdentityProvider200ResponseWithDefaults instantiates a new GetIamV2IdentityProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetIamV2IdentityProvider200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetIamV2IdentityProvider200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetIamV2IdentityProvider200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetIamV2IdentityProvider200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetIamV2IdentityProvider200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetIamV2IdentityProvider200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetIamV2IdentityProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIamV2IdentityProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIamV2IdentityProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetIamV2IdentityProvider200Response) GetMetadata() IamV2IdentityProviderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIamV2IdentityProvider200Response) GetMetadataOk() (*IamV2IdentityProviderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIamV2IdentityProvider200Response) SetMetadata(v IamV2IdentityProviderMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIamV2IdentityProvider200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetIamV2IdentityProvider200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetIamV2IdentityProvider200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetIamV2IdentityProvider200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *GetIamV2IdentityProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIamV2IdentityProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIamV2IdentityProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetState

`func (o *GetIamV2IdentityProvider200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetIamV2IdentityProvider200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetIamV2IdentityProvider200Response) SetState(v string)`

SetState sets State field to given value.


### GetIssuer

`func (o *GetIamV2IdentityProvider200Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetIamV2IdentityProvider200Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetIamV2IdentityProvider200Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUri

`func (o *GetIamV2IdentityProvider200Response) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *GetIamV2IdentityProvider200Response) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *GetIamV2IdentityProvider200Response) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.


### GetKeys

`func (o *GetIamV2IdentityProvider200Response) GetKeys() []IamV2JwksObject`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *GetIamV2IdentityProvider200Response) GetKeysOk() (*[]IamV2JwksObject, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *GetIamV2IdentityProvider200Response) SetKeys(v []IamV2JwksObject)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *GetIamV2IdentityProvider200Response) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


