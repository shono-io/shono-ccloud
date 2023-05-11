# GetServiceQuotaV1Scope200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**ServiceQuotaV1ScopeMetadata**](ServiceQuotaV1ScopeMetadata.md) |  | [optional] 
**Description** | **string** | the quota scope for listing quotas queries | 

## Methods

### NewGetServiceQuotaV1Scope200Response

`func NewGetServiceQuotaV1Scope200Response(apiVersion string, kind string, id string, description string, ) *GetServiceQuotaV1Scope200Response`

NewGetServiceQuotaV1Scope200Response instantiates a new GetServiceQuotaV1Scope200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceQuotaV1Scope200ResponseWithDefaults

`func NewGetServiceQuotaV1Scope200ResponseWithDefaults() *GetServiceQuotaV1Scope200Response`

NewGetServiceQuotaV1Scope200ResponseWithDefaults instantiates a new GetServiceQuotaV1Scope200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetServiceQuotaV1Scope200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetServiceQuotaV1Scope200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetServiceQuotaV1Scope200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetServiceQuotaV1Scope200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetServiceQuotaV1Scope200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetServiceQuotaV1Scope200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetServiceQuotaV1Scope200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetServiceQuotaV1Scope200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetServiceQuotaV1Scope200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetServiceQuotaV1Scope200Response) GetMetadata() ServiceQuotaV1ScopeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetServiceQuotaV1Scope200Response) GetMetadataOk() (*ServiceQuotaV1ScopeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetServiceQuotaV1Scope200Response) SetMetadata(v ServiceQuotaV1ScopeMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetServiceQuotaV1Scope200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *GetServiceQuotaV1Scope200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetServiceQuotaV1Scope200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetServiceQuotaV1Scope200Response) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


