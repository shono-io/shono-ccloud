# IamV2IdentityProviderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**IamV2IdentityProviderListMetadata**](IamV2IdentityProviderListMetadata.md) |  | 
**Data** | [**[]IamV2IdentityProviderListDataInner**](IamV2IdentityProviderListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewIamV2IdentityProviderList

`func NewIamV2IdentityProviderList(apiVersion string, kind string, metadata IamV2IdentityProviderListMetadata, data []IamV2IdentityProviderListDataInner, ) *IamV2IdentityProviderList`

NewIamV2IdentityProviderList instantiates a new IamV2IdentityProviderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IdentityProviderListWithDefaults

`func NewIamV2IdentityProviderListWithDefaults() *IamV2IdentityProviderList`

NewIamV2IdentityProviderListWithDefaults instantiates a new IamV2IdentityProviderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IdentityProviderList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IdentityProviderList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IdentityProviderList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *IamV2IdentityProviderList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IdentityProviderList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IdentityProviderList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *IamV2IdentityProviderList) GetMetadata() IamV2IdentityProviderListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2IdentityProviderList) GetMetadataOk() (*IamV2IdentityProviderListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2IdentityProviderList) SetMetadata(v IamV2IdentityProviderListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *IamV2IdentityProviderList) GetData() []IamV2IdentityProviderListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamV2IdentityProviderList) GetDataOk() (*[]IamV2IdentityProviderListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamV2IdentityProviderList) SetData(v []IamV2IdentityProviderListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


