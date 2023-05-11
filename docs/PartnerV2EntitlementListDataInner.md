# PartnerV2EntitlementListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**ExternalId** | **string** | The unique external ID of the entitlement (this should be unique to customer) | 
**Name** | **string** | The name of the entitlement | 
**PlanId** | **string** | The plan ID the entitlement | 
**ProductId** | **string** | The product ID of the entitlement | 
**UsageReportingId** | Pointer to **string** | The usage reporting ID of the entitlement (if usage reporting uses a different ID, otherwise, same as external_id)  | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the entitlement | [optional] 
**Organization** | Pointer to [**PartnerV2EntitlementOrganization**](PartnerV2EntitlementOrganization.md) |  | [optional] 

## Methods

### NewPartnerV2EntitlementListDataInner

`func NewPartnerV2EntitlementListDataInner(id string, metadata ObjectMeta, externalId string, name string, planId string, productId string, ) *PartnerV2EntitlementListDataInner`

NewPartnerV2EntitlementListDataInner instantiates a new PartnerV2EntitlementListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerV2EntitlementListDataInnerWithDefaults

`func NewPartnerV2EntitlementListDataInnerWithDefaults() *PartnerV2EntitlementListDataInner`

NewPartnerV2EntitlementListDataInnerWithDefaults instantiates a new PartnerV2EntitlementListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerV2EntitlementListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerV2EntitlementListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerV2EntitlementListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerV2EntitlementListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerV2EntitlementListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerV2EntitlementListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerV2EntitlementListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerV2EntitlementListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerV2EntitlementListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerV2EntitlementListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerV2EntitlementListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *PartnerV2EntitlementListDataInner) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerV2EntitlementListDataInner) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerV2EntitlementListDataInner) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetExternalId

`func (o *PartnerV2EntitlementListDataInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PartnerV2EntitlementListDataInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PartnerV2EntitlementListDataInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *PartnerV2EntitlementListDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerV2EntitlementListDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerV2EntitlementListDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *PartnerV2EntitlementListDataInner) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PartnerV2EntitlementListDataInner) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PartnerV2EntitlementListDataInner) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetProductId

`func (o *PartnerV2EntitlementListDataInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PartnerV2EntitlementListDataInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PartnerV2EntitlementListDataInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUsageReportingId

`func (o *PartnerV2EntitlementListDataInner) GetUsageReportingId() string`

GetUsageReportingId returns the UsageReportingId field if non-nil, zero value otherwise.

### GetUsageReportingIdOk

`func (o *PartnerV2EntitlementListDataInner) GetUsageReportingIdOk() (*string, bool)`

GetUsageReportingIdOk returns a tuple with the UsageReportingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportingId

`func (o *PartnerV2EntitlementListDataInner) SetUsageReportingId(v string)`

SetUsageReportingId sets UsageReportingId field to given value.

### HasUsageReportingId

`func (o *PartnerV2EntitlementListDataInner) HasUsageReportingId() bool`

HasUsageReportingId returns a boolean if a field has been set.

### GetResourceId

`func (o *PartnerV2EntitlementListDataInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PartnerV2EntitlementListDataInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PartnerV2EntitlementListDataInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PartnerV2EntitlementListDataInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetOrganization

`func (o *PartnerV2EntitlementListDataInner) GetOrganization() PartnerV2EntitlementOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerV2EntitlementListDataInner) GetOrganizationOk() (*PartnerV2EntitlementOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerV2EntitlementListDataInner) SetOrganization(v PartnerV2EntitlementOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerV2EntitlementListDataInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


