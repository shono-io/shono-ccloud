# PartnerSignupRequestEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 
**Id** | **string** | ID of the referred resource | 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ExternalId** | **string** | The unique external ID of the entitlement (this should be unique to customer) | 
**Name** | **string** | The name of the entitlement | 
**PlanId** | **string** | The plan ID the entitlement | 
**ProductId** | **string** | The product ID of the entitlement | 
**UsageReportingId** | Pointer to **string** | The usage reporting ID of the entitlement (if usage reporting uses a different ID, otherwise, same as external_id)  | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the entitlement | [optional] 
**Organization** | Pointer to [**PartnerV2EntitlementOrganization**](PartnerV2EntitlementOrganization.md) |  | [optional] 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewPartnerSignupRequestEntitlement

`func NewPartnerSignupRequestEntitlement(id string, externalId string, name string, planId string, productId string, related string, resourceName string, ) *PartnerSignupRequestEntitlement`

NewPartnerSignupRequestEntitlement instantiates a new PartnerSignupRequestEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestEntitlementWithDefaults

`func NewPartnerSignupRequestEntitlementWithDefaults() *PartnerSignupRequestEntitlement`

NewPartnerSignupRequestEntitlementWithDefaults instantiates a new PartnerSignupRequestEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerSignupRequestEntitlement) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerSignupRequestEntitlement) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerSignupRequestEntitlement) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerSignupRequestEntitlement) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerSignupRequestEntitlement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerSignupRequestEntitlement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerSignupRequestEntitlement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerSignupRequestEntitlement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerSignupRequestEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerSignupRequestEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerSignupRequestEntitlement) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *PartnerSignupRequestEntitlement) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerSignupRequestEntitlement) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerSignupRequestEntitlement) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartnerSignupRequestEntitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *PartnerSignupRequestEntitlement) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PartnerSignupRequestEntitlement) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PartnerSignupRequestEntitlement) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *PartnerSignupRequestEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerSignupRequestEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerSignupRequestEntitlement) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *PartnerSignupRequestEntitlement) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PartnerSignupRequestEntitlement) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PartnerSignupRequestEntitlement) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetProductId

`func (o *PartnerSignupRequestEntitlement) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PartnerSignupRequestEntitlement) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PartnerSignupRequestEntitlement) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetUsageReportingId

`func (o *PartnerSignupRequestEntitlement) GetUsageReportingId() string`

GetUsageReportingId returns the UsageReportingId field if non-nil, zero value otherwise.

### GetUsageReportingIdOk

`func (o *PartnerSignupRequestEntitlement) GetUsageReportingIdOk() (*string, bool)`

GetUsageReportingIdOk returns a tuple with the UsageReportingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportingId

`func (o *PartnerSignupRequestEntitlement) SetUsageReportingId(v string)`

SetUsageReportingId sets UsageReportingId field to given value.

### HasUsageReportingId

`func (o *PartnerSignupRequestEntitlement) HasUsageReportingId() bool`

HasUsageReportingId returns a boolean if a field has been set.

### GetResourceId

`func (o *PartnerSignupRequestEntitlement) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PartnerSignupRequestEntitlement) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PartnerSignupRequestEntitlement) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PartnerSignupRequestEntitlement) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetOrganization

`func (o *PartnerSignupRequestEntitlement) GetOrganization() PartnerV2EntitlementOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerSignupRequestEntitlement) GetOrganizationOk() (*PartnerV2EntitlementOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerSignupRequestEntitlement) SetOrganization(v PartnerV2EntitlementOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerSignupRequestEntitlement) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEnvironment

`func (o *PartnerSignupRequestEntitlement) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PartnerSignupRequestEntitlement) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PartnerSignupRequestEntitlement) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PartnerSignupRequestEntitlement) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *PartnerSignupRequestEntitlement) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *PartnerSignupRequestEntitlement) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *PartnerSignupRequestEntitlement) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *PartnerSignupRequestEntitlement) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PartnerSignupRequestEntitlement) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PartnerSignupRequestEntitlement) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


