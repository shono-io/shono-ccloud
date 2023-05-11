# IamV2UserListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**IamV2UserMetadata**](IamV2UserMetadata.md) |  | 
**Email** | **string** | The user&#39;s email address | 
**FullName** | Pointer to **string** | The user&#39;s full name | [optional] 
**AuthType** | Pointer to **string** | The user&#39;s authentication method | [optional] [readonly] 

## Methods

### NewIamV2UserListDataInner

`func NewIamV2UserListDataInner(id string, metadata IamV2UserMetadata, email string, ) *IamV2UserListDataInner`

NewIamV2UserListDataInner instantiates a new IamV2UserListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2UserListDataInnerWithDefaults

`func NewIamV2UserListDataInnerWithDefaults() *IamV2UserListDataInner`

NewIamV2UserListDataInnerWithDefaults instantiates a new IamV2UserListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2UserListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2UserListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2UserListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2UserListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2UserListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2UserListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2UserListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2UserListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *IamV2UserListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2UserListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2UserListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *IamV2UserListDataInner) GetMetadata() IamV2UserMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2UserListDataInner) GetMetadataOk() (*IamV2UserMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2UserListDataInner) SetMetadata(v IamV2UserMetadata)`

SetMetadata sets Metadata field to given value.


### GetEmail

`func (o *IamV2UserListDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamV2UserListDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamV2UserListDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *IamV2UserListDataInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *IamV2UserListDataInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *IamV2UserListDataInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *IamV2UserListDataInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetAuthType

`func (o *IamV2UserListDataInner) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *IamV2UserListDataInner) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *IamV2UserListDataInner) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *IamV2UserListDataInner) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


