# PartnerSignupRequestEntitlementOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 

## Methods

### NewPartnerSignupRequestEntitlementOneOf1

`func NewPartnerSignupRequestEntitlementOneOf1(id string, related string, resourceName string, ) *PartnerSignupRequestEntitlementOneOf1`

NewPartnerSignupRequestEntitlementOneOf1 instantiates a new PartnerSignupRequestEntitlementOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestEntitlementOneOf1WithDefaults

`func NewPartnerSignupRequestEntitlementOneOf1WithDefaults() *PartnerSignupRequestEntitlementOneOf1`

NewPartnerSignupRequestEntitlementOneOf1WithDefaults instantiates a new PartnerSignupRequestEntitlementOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnerSignupRequestEntitlementOneOf1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerSignupRequestEntitlementOneOf1) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *PartnerSignupRequestEntitlementOneOf1) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PartnerSignupRequestEntitlementOneOf1) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PartnerSignupRequestEntitlementOneOf1) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *PartnerSignupRequestEntitlementOneOf1) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *PartnerSignupRequestEntitlementOneOf1) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *PartnerSignupRequestEntitlementOneOf1) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *PartnerSignupRequestEntitlementOneOf1) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf1) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf1) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerSignupRequestEntitlementOneOf1) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerSignupRequestEntitlementOneOf1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerSignupRequestEntitlementOneOf1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerSignupRequestEntitlementOneOf1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerSignupRequestEntitlementOneOf1) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


