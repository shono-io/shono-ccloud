# ValidateConnectv1ConnectorPlugin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The class name of the connector plugin. | [optional] 
**Groups** | Pointer to **[]string** | The list of groups used in configuration definitions. | [optional] 
**ErrorCount** | Pointer to **int32** | The total number of errors encountered during configuration validation. | [optional] 
**Configs** | Pointer to [**[]ValidateConnectv1ConnectorPlugin200ResponseConfigsInner**](ValidateConnectv1ConnectorPlugin200ResponseConfigsInner.md) |  | [optional] 

## Methods

### NewValidateConnectv1ConnectorPlugin200Response

`func NewValidateConnectv1ConnectorPlugin200Response() *ValidateConnectv1ConnectorPlugin200Response`

NewValidateConnectv1ConnectorPlugin200Response instantiates a new ValidateConnectv1ConnectorPlugin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateConnectv1ConnectorPlugin200ResponseWithDefaults

`func NewValidateConnectv1ConnectorPlugin200ResponseWithDefaults() *ValidateConnectv1ConnectorPlugin200Response`

NewValidateConnectv1ConnectorPlugin200ResponseWithDefaults instantiates a new ValidateConnectv1ConnectorPlugin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateConnectv1ConnectorPlugin200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValidateConnectv1ConnectorPlugin200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroups

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ValidateConnectv1ConnectorPlugin200Response) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ValidateConnectv1ConnectorPlugin200Response) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetErrorCount

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *ValidateConnectv1ConnectorPlugin200Response) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *ValidateConnectv1ConnectorPlugin200Response) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetConfigs

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetConfigs() []ValidateConnectv1ConnectorPlugin200ResponseConfigsInner`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *ValidateConnectv1ConnectorPlugin200Response) GetConfigsOk() (*[]ValidateConnectv1ConnectorPlugin200ResponseConfigsInner, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *ValidateConnectv1ConnectorPlugin200Response) SetConfigs(v []ValidateConnectv1ConnectorPlugin200ResponseConfigsInner)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *ValidateConnectv1ConnectorPlugin200Response) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


