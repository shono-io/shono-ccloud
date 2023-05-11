# BillingV1CostResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the resource. | [optional] 
**Environment** | Pointer to [**NullableBillingV1ResourceEnvironment**](BillingV1ResourceEnvironment.md) |  | [optional] 

## Methods

### NewBillingV1CostResource

`func NewBillingV1CostResource() *BillingV1CostResource`

NewBillingV1CostResource instantiates a new BillingV1CostResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1CostResourceWithDefaults

`func NewBillingV1CostResourceWithDefaults() *BillingV1CostResource`

NewBillingV1CostResourceWithDefaults instantiates a new BillingV1CostResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingV1CostResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingV1CostResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingV1CostResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingV1CostResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *BillingV1CostResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BillingV1CostResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BillingV1CostResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BillingV1CostResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnvironment

`func (o *BillingV1CostResource) GetEnvironment() BillingV1ResourceEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BillingV1CostResource) GetEnvironmentOk() (*BillingV1ResourceEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BillingV1CostResource) SetEnvironment(v BillingV1ResourceEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BillingV1CostResource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BillingV1CostResource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BillingV1CostResource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


