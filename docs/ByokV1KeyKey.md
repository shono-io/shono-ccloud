# ByokV1KeyKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyArn** | **string** | The Amazon Resource Name (ARN) of an AWS KMS key.  | 
**Roles** | Pointer to **[]string** | The Amazon Resource Names (ARNs) of IAM Roles created for this key-environment combination.  | [optional] [readonly] 
**Kind** | **string** | BYOK kind type.  | 
**ApplicationId** | Pointer to **string** | The Application ID created for this key-environment combination.  | [optional] [readonly] 
**KeyId** | **string** | The unique Key Object Identifier URL without version of an Azure Key Vault key.  | 
**KeyVaultId** | **string** | Key Vault ID containing the key  | 
**TenantId** | **string** | Tenant ID (uuid) hosting the Key Vault containing the key  | 

## Methods

### NewByokV1KeyKey

`func NewByokV1KeyKey(keyArn string, kind string, keyId string, keyVaultId string, tenantId string, ) *ByokV1KeyKey`

NewByokV1KeyKey instantiates a new ByokV1KeyKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeyKeyWithDefaults

`func NewByokV1KeyKeyWithDefaults() *ByokV1KeyKey`

NewByokV1KeyKeyWithDefaults instantiates a new ByokV1KeyKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyArn

`func (o *ByokV1KeyKey) GetKeyArn() string`

GetKeyArn returns the KeyArn field if non-nil, zero value otherwise.

### GetKeyArnOk

`func (o *ByokV1KeyKey) GetKeyArnOk() (*string, bool)`

GetKeyArnOk returns a tuple with the KeyArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyArn

`func (o *ByokV1KeyKey) SetKeyArn(v string)`

SetKeyArn sets KeyArn field to given value.


### GetRoles

`func (o *ByokV1KeyKey) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ByokV1KeyKey) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ByokV1KeyKey) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ByokV1KeyKey) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetKind

`func (o *ByokV1KeyKey) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ByokV1KeyKey) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ByokV1KeyKey) SetKind(v string)`

SetKind sets Kind field to given value.


### GetApplicationId

`func (o *ByokV1KeyKey) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ByokV1KeyKey) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ByokV1KeyKey) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ByokV1KeyKey) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetKeyId

`func (o *ByokV1KeyKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ByokV1KeyKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ByokV1KeyKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetKeyVaultId

`func (o *ByokV1KeyKey) GetKeyVaultId() string`

GetKeyVaultId returns the KeyVaultId field if non-nil, zero value otherwise.

### GetKeyVaultIdOk

`func (o *ByokV1KeyKey) GetKeyVaultIdOk() (*string, bool)`

GetKeyVaultIdOk returns a tuple with the KeyVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyVaultId

`func (o *ByokV1KeyKey) SetKeyVaultId(v string)`

SetKeyVaultId sets KeyVaultId field to given value.


### GetTenantId

`func (o *ByokV1KeyKey) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ByokV1KeyKey) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ByokV1KeyKey) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


