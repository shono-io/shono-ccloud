# IamV2InvitationCreator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 

## Methods

### NewIamV2InvitationCreator

`func NewIamV2InvitationCreator(id string, related string, resourceName string, ) *IamV2InvitationCreator`

NewIamV2InvitationCreator instantiates a new IamV2InvitationCreator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2InvitationCreatorWithDefaults

`func NewIamV2InvitationCreatorWithDefaults() *IamV2InvitationCreator`

NewIamV2InvitationCreatorWithDefaults instantiates a new IamV2InvitationCreator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IamV2InvitationCreator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamV2InvitationCreator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamV2InvitationCreator) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *IamV2InvitationCreator) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *IamV2InvitationCreator) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *IamV2InvitationCreator) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *IamV2InvitationCreator) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *IamV2InvitationCreator) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *IamV2InvitationCreator) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


