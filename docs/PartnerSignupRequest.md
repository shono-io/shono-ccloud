# PartnerSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | [**PartnerSignupRequestOrganization**](PartnerSignupRequestOrganization.md) |  | 
**User** | Pointer to **interface{}** |  | [optional] 
**Entitlement** | [**PartnerSignupRequestEntitlement**](PartnerSignupRequestEntitlement.md) |  | 

## Methods

### NewPartnerSignupRequest

`func NewPartnerSignupRequest(organization PartnerSignupRequestOrganization, entitlement PartnerSignupRequestEntitlement, ) *PartnerSignupRequest`

NewPartnerSignupRequest instantiates a new PartnerSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupRequestWithDefaults

`func NewPartnerSignupRequestWithDefaults() *PartnerSignupRequest`

NewPartnerSignupRequestWithDefaults instantiates a new PartnerSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *PartnerSignupRequest) GetOrganization() PartnerSignupRequestOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerSignupRequest) GetOrganizationOk() (*PartnerSignupRequestOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerSignupRequest) SetOrganization(v PartnerSignupRequestOrganization)`

SetOrganization sets Organization field to given value.


### GetUser

`func (o *PartnerSignupRequest) GetUser() interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PartnerSignupRequest) GetUserOk() (*interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PartnerSignupRequest) SetUser(v interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *PartnerSignupRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEntitlement

`func (o *PartnerSignupRequest) GetEntitlement() PartnerSignupRequestEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *PartnerSignupRequest) GetEntitlementOk() (*PartnerSignupRequestEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *PartnerSignupRequest) SetEntitlement(v PartnerSignupRequestEntitlement)`

SetEntitlement sets Entitlement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


