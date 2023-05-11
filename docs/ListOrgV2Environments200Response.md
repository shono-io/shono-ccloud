# ListOrgV2Environments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**OrgV2EnvironmentListMetadata**](OrgV2EnvironmentListMetadata.md) |  | 
**Data** | [**[]OrgV2EnvironmentListDataInner**](OrgV2EnvironmentListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListOrgV2Environments200Response

`func NewListOrgV2Environments200Response(apiVersion string, kind string, metadata OrgV2EnvironmentListMetadata, data []OrgV2EnvironmentListDataInner, ) *ListOrgV2Environments200Response`

NewListOrgV2Environments200Response instantiates a new ListOrgV2Environments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgV2Environments200ResponseWithDefaults

`func NewListOrgV2Environments200ResponseWithDefaults() *ListOrgV2Environments200Response`

NewListOrgV2Environments200ResponseWithDefaults instantiates a new ListOrgV2Environments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListOrgV2Environments200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListOrgV2Environments200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListOrgV2Environments200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListOrgV2Environments200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListOrgV2Environments200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListOrgV2Environments200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListOrgV2Environments200Response) GetMetadata() OrgV2EnvironmentListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListOrgV2Environments200Response) GetMetadataOk() (*OrgV2EnvironmentListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListOrgV2Environments200Response) SetMetadata(v OrgV2EnvironmentListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListOrgV2Environments200Response) GetData() []OrgV2EnvironmentListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListOrgV2Environments200Response) GetDataOk() (*[]OrgV2EnvironmentListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListOrgV2Environments200Response) SetData(v []OrgV2EnvironmentListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


