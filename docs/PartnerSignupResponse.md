# PartnerSignupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** | The ID of the organization | 
**SsoUrl** | **string** | The login URL for the customer to access Confluent Cloud | 
**DisplayMessage** | Pointer to **string** | The display message contains useful information which is shown on the Marketplace UI to the customers. | [optional] 

## Methods

### NewPartnerSignupResponse

`func NewPartnerSignupResponse(organizationId string, ssoUrl string, ) *PartnerSignupResponse`

NewPartnerSignupResponse instantiates a new PartnerSignupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSignupResponseWithDefaults

`func NewPartnerSignupResponseWithDefaults() *PartnerSignupResponse`

NewPartnerSignupResponseWithDefaults instantiates a new PartnerSignupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *PartnerSignupResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PartnerSignupResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PartnerSignupResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetSsoUrl

`func (o *PartnerSignupResponse) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *PartnerSignupResponse) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *PartnerSignupResponse) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### GetDisplayMessage

`func (o *PartnerSignupResponse) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *PartnerSignupResponse) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *PartnerSignupResponse) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *PartnerSignupResponse) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


