# ListIamV2RoleBindings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**IamV2RoleBindingListMetadata**](IamV2RoleBindingListMetadata.md) |  | 
**Data** | [**[]IamV2RoleBindingListDataInner**](IamV2RoleBindingListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListIamV2RoleBindings200Response

`func NewListIamV2RoleBindings200Response(apiVersion string, kind string, metadata IamV2RoleBindingListMetadata, data []IamV2RoleBindingListDataInner, ) *ListIamV2RoleBindings200Response`

NewListIamV2RoleBindings200Response instantiates a new ListIamV2RoleBindings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIamV2RoleBindings200ResponseWithDefaults

`func NewListIamV2RoleBindings200ResponseWithDefaults() *ListIamV2RoleBindings200Response`

NewListIamV2RoleBindings200ResponseWithDefaults instantiates a new ListIamV2RoleBindings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListIamV2RoleBindings200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListIamV2RoleBindings200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListIamV2RoleBindings200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListIamV2RoleBindings200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListIamV2RoleBindings200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListIamV2RoleBindings200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListIamV2RoleBindings200Response) GetMetadata() IamV2RoleBindingListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListIamV2RoleBindings200Response) GetMetadataOk() (*IamV2RoleBindingListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListIamV2RoleBindings200Response) SetMetadata(v IamV2RoleBindingListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListIamV2RoleBindings200Response) GetData() []IamV2RoleBindingListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIamV2RoleBindings200Response) GetDataOk() (*[]IamV2RoleBindingListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIamV2RoleBindings200Response) SetData(v []IamV2RoleBindingListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


