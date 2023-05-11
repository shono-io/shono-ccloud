# CreateIamV2ServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2ServiceAccountMetadata**](IamV2ServiceAccountMetadata.md) |  | [optional] 
**DisplayName** | **string** | A human-readable name for the Service Account | 
**Description** | Pointer to **string** | A free-form description of the Service Account | [optional] 

## Methods

### NewCreateIamV2ServiceAccountRequest

`func NewCreateIamV2ServiceAccountRequest(displayName string, ) *CreateIamV2ServiceAccountRequest`

NewCreateIamV2ServiceAccountRequest instantiates a new CreateIamV2ServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2ServiceAccountRequestWithDefaults

`func NewCreateIamV2ServiceAccountRequestWithDefaults() *CreateIamV2ServiceAccountRequest`

NewCreateIamV2ServiceAccountRequestWithDefaults instantiates a new CreateIamV2ServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2ServiceAccountRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2ServiceAccountRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2ServiceAccountRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2ServiceAccountRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2ServiceAccountRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2ServiceAccountRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2ServiceAccountRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2ServiceAccountRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2ServiceAccountRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2ServiceAccountRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2ServiceAccountRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2ServiceAccountRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2ServiceAccountRequest) GetMetadata() IamV2ServiceAccountMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2ServiceAccountRequest) GetMetadataOk() (*IamV2ServiceAccountMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2ServiceAccountRequest) SetMetadata(v IamV2ServiceAccountMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2ServiceAccountRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateIamV2ServiceAccountRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateIamV2ServiceAccountRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateIamV2ServiceAccountRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CreateIamV2ServiceAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIamV2ServiceAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIamV2ServiceAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIamV2ServiceAccountRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


