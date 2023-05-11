# IamV2ServiceAccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**IamV2ServiceAccountListMetadata**](IamV2ServiceAccountListMetadata.md) |  | 
**Data** | [**[]IamV2ServiceAccountListDataInner**](IamV2ServiceAccountListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewIamV2ServiceAccountList

`func NewIamV2ServiceAccountList(apiVersion string, kind string, metadata IamV2ServiceAccountListMetadata, data []IamV2ServiceAccountListDataInner, ) *IamV2ServiceAccountList`

NewIamV2ServiceAccountList instantiates a new IamV2ServiceAccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2ServiceAccountListWithDefaults

`func NewIamV2ServiceAccountListWithDefaults() *IamV2ServiceAccountList`

NewIamV2ServiceAccountListWithDefaults instantiates a new IamV2ServiceAccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2ServiceAccountList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2ServiceAccountList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2ServiceAccountList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *IamV2ServiceAccountList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2ServiceAccountList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2ServiceAccountList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *IamV2ServiceAccountList) GetMetadata() IamV2ServiceAccountListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamV2ServiceAccountList) GetMetadataOk() (*IamV2ServiceAccountListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamV2ServiceAccountList) SetMetadata(v IamV2ServiceAccountListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *IamV2ServiceAccountList) GetData() []IamV2ServiceAccountListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamV2ServiceAccountList) GetDataOk() (*[]IamV2ServiceAccountListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamV2ServiceAccountList) SetData(v []IamV2ServiceAccountListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


