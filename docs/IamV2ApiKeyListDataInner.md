# IamV2ApiKeyListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**IamV2ApiKeyMetadata**](IamV2ApiKeyMetadata.md) |  | 
**Spec** | **map[string]interface{}** |  | 

## Methods

### NewIamV2ApiKeyListDataInner

`func NewIamV2ApiKeyListDataInner(id string, metadata IamV2ApiKeyMetadata, spec map[string]interface{}, ) *IamV2ApiKeyListDataInner`

NewIamV2ApiKeyListDataInner instantiates a new IamV2ApiKeyListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ApiKeyListDataInnerWithDefaults

`func NewIamV2ApiKeyListDataInnerWithDefaults() *IamV2ApiKeyListDataInner`

NewIamV2ApiKeyListDataInnerWithDefaults instantiates a new IamV2ApiKeyListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2ApiKeyListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2ApiKeyListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2ApiKeyListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2ApiKeyListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2ApiKeyListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2ApiKeyListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2ApiKeyListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2ApiKeyListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2ApiKeyListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2ApiKeyListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2ApiKeyListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IamV2ApiKeyListDataInner) GetMetadata() IamV2ApiKeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2ApiKeyListDataInner) GetMetadataOk() (*IamV2ApiKeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2ApiKeyListDataInner) SetMetadata(v IamV2ApiKeyMetadata)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *IamV2ApiKeyListDataInner) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IamV2ApiKeyListDataInner) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IamV2ApiKeyListDataInner) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


