# RefreshIamV2JsonWebKeySet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Spec** | **map[string]interface{}** |  | 
**Status** | Pointer to [**IamV2JwksStatus**](IamV2JwksStatus.md) |  | [optional] 

## Methods

### NewRefreshIamV2JsonWebKeySet200Response

`func NewRefreshIamV2JsonWebKeySet200Response(apiVersion string, kind string, spec map[string]interface{}, ) *RefreshIamV2JsonWebKeySet200Response`

NewRefreshIamV2JsonWebKeySet200Response instantiates a new RefreshIamV2JsonWebKeySet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshIamV2JsonWebKeySet200ResponseWithDefaults

`func NewRefreshIamV2JsonWebKeySet200ResponseWithDefaults() *RefreshIamV2JsonWebKeySet200Response`

NewRefreshIamV2JsonWebKeySet200ResponseWithDefaults instantiates a new RefreshIamV2JsonWebKeySet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RefreshIamV2JsonWebKeySet200Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RefreshIamV2JsonWebKeySet200Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RefreshIamV2JsonWebKeySet200Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *RefreshIamV2JsonWebKeySet200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RefreshIamV2JsonWebKeySet200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RefreshIamV2JsonWebKeySet200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSpec

`func (o *RefreshIamV2JsonWebKeySet200Response) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RefreshIamV2JsonWebKeySet200Response) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RefreshIamV2JsonWebKeySet200Response) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *RefreshIamV2JsonWebKeySet200Response) GetStatus() IamV2JwksStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefreshIamV2JsonWebKeySet200Response) GetStatusOk() (*IamV2JwksStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefreshIamV2JsonWebKeySet200Response) SetStatus(v IamV2JwksStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RefreshIamV2JsonWebKeySet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


