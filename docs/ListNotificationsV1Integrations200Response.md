# ListNotificationsV1Integrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**NotificationsV1IntegrationListMetadata**](NotificationsV1IntegrationListMetadata.md) |  | 
**Data** | [**[]NotificationsV1IntegrationListDataInner**](NotificationsV1IntegrationListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListNotificationsV1Integrations200Response

`func NewListNotificationsV1Integrations200Response(apiVersion string, kind string, metadata NotificationsV1IntegrationListMetadata, data []NotificationsV1IntegrationListDataInner, ) *ListNotificationsV1Integrations200Response`

NewListNotificationsV1Integrations200Response instantiates a new ListNotificationsV1Integrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationsV1Integrations200ResponseWithDefaults

`func NewListNotificationsV1Integrations200ResponseWithDefaults() *ListNotificationsV1Integrations200Response`

NewListNotificationsV1Integrations200ResponseWithDefaults instantiates a new ListNotificationsV1Integrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListNotificationsV1Integrations200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListNotificationsV1Integrations200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListNotificationsV1Integrations200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListNotificationsV1Integrations200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListNotificationsV1Integrations200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListNotificationsV1Integrations200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListNotificationsV1Integrations200Response) GetMetadata() NotificationsV1IntegrationListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListNotificationsV1Integrations200Response) GetMetadataOk() (*NotificationsV1IntegrationListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListNotificationsV1Integrations200Response) SetMetadata(v NotificationsV1IntegrationListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListNotificationsV1Integrations200Response) GetData() []NotificationsV1IntegrationListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNotificationsV1Integrations200Response) GetDataOk() (*[]NotificationsV1IntegrationListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNotificationsV1Integrations200Response) SetData(v []NotificationsV1IntegrationListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


