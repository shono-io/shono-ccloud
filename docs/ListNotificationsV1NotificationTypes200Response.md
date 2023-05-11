# ListNotificationsV1NotificationTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**NotificationsV1NotificationTypeListMetadata**](NotificationsV1NotificationTypeListMetadata.md) |  | 
**Data** | [**[]NotificationsV1NotificationTypeListDataInner**](NotificationsV1NotificationTypeListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewListNotificationsV1NotificationTypes200Response

`func NewListNotificationsV1NotificationTypes200Response(apiVersion string, kind string, metadata NotificationsV1NotificationTypeListMetadata, data []NotificationsV1NotificationTypeListDataInner, ) *ListNotificationsV1NotificationTypes200Response`

NewListNotificationsV1NotificationTypes200Response instantiates a new ListNotificationsV1NotificationTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationsV1NotificationTypes200ResponseWithDefaults

`func NewListNotificationsV1NotificationTypes200ResponseWithDefaults() *ListNotificationsV1NotificationTypes200Response`

NewListNotificationsV1NotificationTypes200ResponseWithDefaults instantiates a new ListNotificationsV1NotificationTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListNotificationsV1NotificationTypes200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListNotificationsV1NotificationTypes200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListNotificationsV1NotificationTypes200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *ListNotificationsV1NotificationTypes200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListNotificationsV1NotificationTypes200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListNotificationsV1NotificationTypes200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListNotificationsV1NotificationTypes200Response) GetMetadata() NotificationsV1NotificationTypeListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListNotificationsV1NotificationTypes200Response) GetMetadataOk() (*NotificationsV1NotificationTypeListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListNotificationsV1NotificationTypes200Response) SetMetadata(v NotificationsV1NotificationTypeListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *ListNotificationsV1NotificationTypes200Response) GetData() []NotificationsV1NotificationTypeListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNotificationsV1NotificationTypes200Response) GetDataOk() (*[]NotificationsV1NotificationTypeListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNotificationsV1NotificationTypes200Response) SetData(v []NotificationsV1NotificationTypeListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


