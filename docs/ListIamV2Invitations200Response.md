# ListIamV2Invitations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**IamV2InvitationListMetadata**](IamV2InvitationListMetadata.md) |  | 
**Data** | [**[]IamV2InvitationListDataInner**](IamV2InvitationListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListIamV2Invitations200Response

`func NewListIamV2Invitations200Response(apiVersion string, kind string, metadata IamV2InvitationListMetadata, data []IamV2InvitationListDataInner, ) *ListIamV2Invitations200Response`

NewListIamV2Invitations200Response instantiates a new ListIamV2Invitations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIamV2Invitations200ResponseWithDefaults

`func NewListIamV2Invitations200ResponseWithDefaults() *ListIamV2Invitations200Response`

NewListIamV2Invitations200ResponseWithDefaults instantiates a new ListIamV2Invitations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListIamV2Invitations200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListIamV2Invitations200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListIamV2Invitations200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListIamV2Invitations200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListIamV2Invitations200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListIamV2Invitations200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListIamV2Invitations200Response) GetMetadata() IamV2InvitationListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListIamV2Invitations200Response) GetMetadataOk() (*IamV2InvitationListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListIamV2Invitations200Response) SetMetadata(v IamV2InvitationListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListIamV2Invitations200Response) GetData() []IamV2InvitationListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIamV2Invitations200Response) GetDataOk() (*[]IamV2InvitationListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIamV2Invitations200Response) SetData(v []IamV2InvitationListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


