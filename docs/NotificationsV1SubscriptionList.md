# NotificationsV1SubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**NotificationsV1SubscriptionListMetadata**](NotificationsV1SubscriptionListMetadata.md) |  | 
**Data** | [**[]NotificationsV1SubscriptionListDataInner**](NotificationsV1SubscriptionListDataInner.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewNotificationsV1SubscriptionList

`func NewNotificationsV1SubscriptionList(apiVersion string, kind string, metadata NotificationsV1SubscriptionListMetadata, data []NotificationsV1SubscriptionListDataInner, ) *NotificationsV1SubscriptionList`

NewNotificationsV1SubscriptionList instantiates a new NotificationsV1SubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsV1SubscriptionListWithDefaults

`func NewNotificationsV1SubscriptionListWithDefaults() *NotificationsV1SubscriptionList`

NewNotificationsV1SubscriptionListWithDefaults instantiates a new NotificationsV1SubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NotificationsV1SubscriptionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NotificationsV1SubscriptionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NotificationsV1SubscriptionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *NotificationsV1SubscriptionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NotificationsV1SubscriptionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NotificationsV1SubscriptionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NotificationsV1SubscriptionList) GetMetadata() NotificationsV1SubscriptionListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NotificationsV1SubscriptionList) GetMetadataOk() (*NotificationsV1SubscriptionListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NotificationsV1SubscriptionList) SetMetadata(v NotificationsV1SubscriptionListMetadata)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *NotificationsV1SubscriptionList) GetData() []NotificationsV1SubscriptionListDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationsV1SubscriptionList) GetDataOk() (*[]NotificationsV1SubscriptionListDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationsV1SubscriptionList) SetData(v []NotificationsV1SubscriptionListDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


