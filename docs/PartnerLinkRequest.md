# PartnerLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The linking token that was generated. | 
**Organization** | [**PartnerLinkRequestOrganization**](PartnerLinkRequestOrganization.md) |  | 
**Entitlement** | [**PartnerSignupRequestEntitlement**](PartnerSignupRequestEntitlement.md) |  | 

## Methods

### NewPartnerLinkRequest

`func NewPartnerLinkRequest(token string, organization PartnerLinkRequestOrganization, entitlement PartnerSignupRequestEntitlement, ) *PartnerLinkRequest`

NewPartnerLinkRequest instantiates a new PartnerLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerLinkRequestWithDefaults

`func NewPartnerLinkRequestWithDefaults() *PartnerLinkRequest`

NewPartnerLinkRequestWithDefaults instantiates a new PartnerLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PartnerLinkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PartnerLinkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PartnerLinkRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetOrganization

`func (o *PartnerLinkRequest) GetOrganization() PartnerLinkRequestOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerLinkRequest) GetOrganizationOk() (*PartnerLinkRequestOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerLinkRequest) SetOrganization(v PartnerLinkRequestOrganization)`

SetOrganization sets Organization field to given value.


### GetEntitlement

`func (o *PartnerLinkRequest) GetEntitlement() PartnerSignupRequestEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *PartnerLinkRequest) GetEntitlementOk() (*PartnerSignupRequestEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *PartnerLinkRequest) SetEntitlement(v PartnerSignupRequestEntitlement)`

SetEntitlement sets Entitlement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


