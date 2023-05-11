# IamV2ApiKeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | The API key secret. Only provided in &#x60;create&#x60; responses, not in &#x60;get&#x60; or &#x60;list&#x60;. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human readable name for the API key | [optional] 
**Description** | Pointer to **string** | A human readable description for the API key | [optional] 
**Owner** | Pointer to [**IamV2ApiKeySpecOwner**](IamV2ApiKeySpecOwner.md) |  | [optional] 
**Resource** | Pointer to [**NullableIamV2ApiKeySpecResource**](IamV2ApiKeySpecResource.md) |  | [optional] 

## Methods

### NewIamV2ApiKeySpec

`func NewIamV2ApiKeySpec() *IamV2ApiKeySpec`

NewIamV2ApiKeySpec instantiates a new IamV2ApiKeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ApiKeySpecWithDefaults

`func NewIamV2ApiKeySpecWithDefaults() *IamV2ApiKeySpec`

NewIamV2ApiKeySpecWithDefaults instantiates a new IamV2ApiKeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *IamV2ApiKeySpec) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IamV2ApiKeySpec) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IamV2ApiKeySpec) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IamV2ApiKeySpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamV2ApiKeySpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamV2ApiKeySpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamV2ApiKeySpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamV2ApiKeySpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *IamV2ApiKeySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamV2ApiKeySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamV2ApiKeySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamV2ApiKeySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IamV2ApiKeySpec) GetOwner() IamV2ApiKeySpecOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamV2ApiKeySpec) GetOwnerOk() (*IamV2ApiKeySpecOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamV2ApiKeySpec) SetOwner(v IamV2ApiKeySpecOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamV2ApiKeySpec) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResource

`func (o *IamV2ApiKeySpec) GetResource() IamV2ApiKeySpecResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IamV2ApiKeySpec) GetResourceOk() (*IamV2ApiKeySpecResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IamV2ApiKeySpec) SetResource(v IamV2ApiKeySpecResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IamV2ApiKeySpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *IamV2ApiKeySpec) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *IamV2ApiKeySpec) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


