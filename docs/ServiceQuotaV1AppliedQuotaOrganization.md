# ServiceQuotaV1AppliedQuotaOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewServiceQuotaV1AppliedQuotaOrganization

`func NewServiceQuotaV1AppliedQuotaOrganization(id string, related string, resourceName string, ) *ServiceQuotaV1AppliedQuotaOrganization`

NewServiceQuotaV1AppliedQuotaOrganization instantiates a new ServiceQuotaV1AppliedQuotaOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQuotaV1AppliedQuotaOrganizationWithDefaults

`func NewServiceQuotaV1AppliedQuotaOrganizationWithDefaults() *ServiceQuotaV1AppliedQuotaOrganization`

NewServiceQuotaV1AppliedQuotaOrganizationWithDefaults instantiates a new ServiceQuotaV1AppliedQuotaOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceQuotaV1AppliedQuotaOrganization) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *ServiceQuotaV1AppliedQuotaOrganization) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ServiceQuotaV1AppliedQuotaOrganization) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ServiceQuotaV1AppliedQuotaOrganization) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


