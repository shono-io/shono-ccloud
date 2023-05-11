# CreateIamV2InvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**IamV2InvitationMetadata**](IamV2InvitationMetadata.md) |  | [optional] 
**Email** | **string** | The user/invitee&#39;s email address | 
**AuthType** | Pointer to **string** | The user/invitee&#39;s authentication type. Note that only the [OrganizationAdmin role](https://docs.confluent.io/cloud/current/access-management/access-control/cloud-rbac.html#organizationadmin) can invite AUTH_TYPE_LOCAL users to SSO organizations. The user&#39;s auth_type is set as AUTH_TYPE_SSO by default if the organization has SSO enabled. Otherwise, the user&#39;s auth_type is AUTH_TYPE_LOCAL by default.  | [optional] 
**Status** | Pointer to **string** | The status of invitations | [optional] [readonly] 
**AcceptedAt** | Pointer to **NullableTime** | The timestamp that the invitation was accepted | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The timestamp that the invitation will expire | [optional] [readonly] 
**User** | Pointer to [**IamV2InvitationUser**](IamV2InvitationUser.md) |  | [optional] 
**Creator** | Pointer to [**IamV2InvitationCreator**](IamV2InvitationCreator.md) |  | [optional] 

## Methods

### NewCreateIamV2InvitationRequest

`func NewCreateIamV2InvitationRequest(email string, ) *CreateIamV2InvitationRequest`

NewCreateIamV2InvitationRequest instantiates a new CreateIamV2InvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamV2InvitationRequestWithDefaults

`func NewCreateIamV2InvitationRequestWithDefaults() *CreateIamV2InvitationRequest`

NewCreateIamV2InvitationRequestWithDefaults instantiates a new CreateIamV2InvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CreateIamV2InvitationRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CreateIamV2InvitationRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CreateIamV2InvitationRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CreateIamV2InvitationRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CreateIamV2InvitationRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIamV2InvitationRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIamV2InvitationRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateIamV2InvitationRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CreateIamV2InvitationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIamV2InvitationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIamV2InvitationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIamV2InvitationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIamV2InvitationRequest) GetMetadata() IamV2InvitationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIamV2InvitationRequest) GetMetadataOk() (*IamV2InvitationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIamV2InvitationRequest) SetMetadata(v IamV2InvitationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIamV2InvitationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEmail

`func (o *CreateIamV2InvitationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateIamV2InvitationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateIamV2InvitationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAuthType

`func (o *CreateIamV2InvitationRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateIamV2InvitationRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateIamV2InvitationRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *CreateIamV2InvitationRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateIamV2InvitationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateIamV2InvitationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateIamV2InvitationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateIamV2InvitationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *CreateIamV2InvitationRequest) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *CreateIamV2InvitationRequest) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *CreateIamV2InvitationRequest) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *CreateIamV2InvitationRequest) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *CreateIamV2InvitationRequest) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *CreateIamV2InvitationRequest) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetExpiresAt

`func (o *CreateIamV2InvitationRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateIamV2InvitationRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateIamV2InvitationRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateIamV2InvitationRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUser

`func (o *CreateIamV2InvitationRequest) GetUser() IamV2InvitationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateIamV2InvitationRequest) GetUserOk() (*IamV2InvitationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateIamV2InvitationRequest) SetUser(v IamV2InvitationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateIamV2InvitationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreator

`func (o *CreateIamV2InvitationRequest) GetCreator() IamV2InvitationCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateIamV2InvitationRequest) GetCreatorOk() (*IamV2InvitationCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateIamV2InvitationRequest) SetCreator(v IamV2InvitationCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CreateIamV2InvitationRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


