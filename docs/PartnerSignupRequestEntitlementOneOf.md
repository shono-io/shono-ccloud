# PartnerSignupRequestEntitlementOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ExternalId** | **string** | The unique external ID of the entitlement (this should be unique to customer) | 
**Name** | **string** | The name of the entitlement | 
**PlanId** | **string** | The plan ID the entitlement | 
**ProductId** | **string** | The product ID of the entitlement | 
**UsageReportingId** | Pointer to **string** | The usage reporting ID of the entitlement (if usage reporting uses a different ID, otherwise, same as external_id)  | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the entitlement | [optional] 
**Organization** | Pointer to [**PartnerV2EntitlementOrganization**](PartnerV2EntitlementOrganization.md) |  | [optional] 

## Methods

### NewPartnerSignupRequestEntitlementOneOf

`func NewPartnerSignupRequestEntitlementOneOf(externalId string, name string, planId string, productId string, ) *PartnerSignupRequestEntitlementOneOf`

NewPartnerSignupRequestEntitlementOneOf instantiates a new PartnerSignupRequestEntitlementOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestEntitlementOneOfWithDefaults

`func NewPartnerSignupRequestEntitlementOneOfWithDefaults() *PartnerSignupRequestEntitlementOneOf`

NewPartnerSignupRequestEntitlementOneOfWithDefaults instantiates a new PartnerSignupRequestEntitlementOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerSignupRequestEntitlementOneOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerSignupRequestEntitlementOneOf) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerSignupRequestEntitlementOneOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerSignupRequestEntitlementOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerSignupRequestEntitlementOneOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartnerSignupRequestEntitlementOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PartnerSignupRequestEntitlementOneOf) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerSignupRequestEntitlementOneOf) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartnerSignupRequestEntitlementOneOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *PartnerSignupRequestEntitlementOneOf) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PartnerSignupRequestEntitlementOneOf) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *PartnerSignupRequestEntitlementOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerSignupRequestEntitlementOneOf) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *PartnerSignupRequestEntitlementOneOf) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PartnerSignupRequestEntitlementOneOf) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetProductId

`func (o *PartnerSignupRequestEntitlementOneOf) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PartnerSignupRequestEntitlementOneOf) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUsageReportingId

`func (o *PartnerSignupRequestEntitlementOneOf) GetUsageReportingId() string`

GetUsageReportingId returns the UsageReportingId field if non-nil, zero value otherwise.

### GetUsageReportingIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetUsageReportingIdOk() (*string, bool)`

GetUsageReportingIdOk returns a tuple with the UsageReportingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportingId

`func (o *PartnerSignupRequestEntitlementOneOf) SetUsageReportingId(v string)`

SetUsageReportingId sets UsageReportingId field to given value.

### HasUsageReportingId

`func (o *PartnerSignupRequestEntitlementOneOf) HasUsageReportingId() bool`

HasUsageReportingId returns a boolean if a field has been set.

### GetResourceId

`func (o *PartnerSignupRequestEntitlementOneOf) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PartnerSignupRequestEntitlementOneOf) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PartnerSignupRequestEntitlementOneOf) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetOrganization

`func (o *PartnerSignupRequestEntitlementOneOf) GetOrganization() PartnerV2EntitlementOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerSignupRequestEntitlementOneOf) GetOrganizationOk() (*PartnerV2EntitlementOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerSignupRequestEntitlementOneOf) SetOrganization(v PartnerV2EntitlementOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerSignupRequestEntitlementOneOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


