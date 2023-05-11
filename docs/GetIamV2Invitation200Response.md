# GetIamV2Invitation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**Metadata** | Pointer to [**IamV2InvitationMetadata**](IamV2InvitationMetadata.md) |  | [optional] 
**Email** | **string** | The user/invitee&#39;s email address | 
**AuthType** | Pointer to **string** | The user/invitee&#39;s authentication type. Note that only the [OrganizationAdmin role](https://docs.confluent.io/cloud/current/access-management/access-control/cloud-rbac.html#organizationadmin) can invite AUTH_TYPE_LOCAL users to SSO organizations. The user&#39;s auth_type is set as AUTH_TYPE_SSO by default if the organization has SSO enabled. Otherwise, the user&#39;s auth_type is AUTH_TYPE_LOCAL by default.  | [optional] 
**Status** | Pointer to **string** | The status of invitations | [optional] [readonly] 
**AcceptedAt** | Pointer to **NullableTime** | The timestamp that the invitation was accepted | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | The timestamp that the invitation will expire | [optional] [readonly] 
**User** | Pointer to [**IamV2InvitationUser**](IamV2InvitationUser.md) |  | [optional] 
**Creator** | Pointer to [**IamV2InvitationCreator**](IamV2InvitationCreator.md) |  | [optional] 

## Methods

### NewGetIamV2Invitation200Response

`func NewGetIamV2Invitation200Response(apiVersion string, kind string, id string, email string, ) *GetIamV2Invitation200Response`

NewGetIamV2Invitation200Response instantiates a new GetIamV2Invitation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIamV2Invitation200ResponseWithDefaults

`func NewGetIamV2Invitation200ResponseWithDefaults() *GetIamV2Invitation200Response`

NewGetIamV2Invitation200ResponseWithDefaults instantiates a new GetIamV2Invitation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetIamV2Invitation200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetIamV2Invitation200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetIamV2Invitation200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *GetIamV2Invitation200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetIamV2Invitation200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetIamV2Invitation200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *GetIamV2Invitation200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIamV2Invitation200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIamV2Invitation200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *GetIamV2Invitation200Response) GetMetadata() IamV2InvitationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIamV2Invitation200Response) GetMetadataOk() (*IamV2InvitationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIamV2Invitation200Response) SetMetadata(v IamV2InvitationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIamV2Invitation200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEmail

`func (o *GetIamV2Invitation200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetIamV2Invitation200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetIamV2Invitation200Response) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAuthType

`func (o *GetIamV2Invitation200Response) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *GetIamV2Invitation200Response) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *GetIamV2Invitation200Response) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *GetIamV2Invitation200Response) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetStatus

`func (o *GetIamV2Invitation200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIamV2Invitation200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIamV2Invitation200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIamV2Invitation200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcceptedAt

`func (o *GetIamV2Invitation200Response) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *GetIamV2Invitation200Response) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *GetIamV2Invitation200Response) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *GetIamV2Invitation200Response) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *GetIamV2Invitation200Response) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *GetIamV2Invitation200Response) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetExpiresAt

`func (o *GetIamV2Invitation200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetIamV2Invitation200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetIamV2Invitation200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetIamV2Invitation200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetUser

`func (o *GetIamV2Invitation200Response) GetUser() IamV2InvitationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetIamV2Invitation200Response) GetUserOk() (*IamV2InvitationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetIamV2Invitation200Response) SetUser(v IamV2InvitationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetIamV2Invitation200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreator

`func (o *GetIamV2Invitation200Response) GetCreator() IamV2InvitationCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetIamV2Invitation200Response) GetCreatorOk() (*IamV2InvitationCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetIamV2Invitation200Response) SetCreator(v IamV2InvitationCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetIamV2Invitation200Response) HasCreator() bool`

HasCreator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


