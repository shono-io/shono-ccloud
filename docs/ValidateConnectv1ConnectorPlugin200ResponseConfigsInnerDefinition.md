# ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the configuration | [optional] 
**Type** | Pointer to **string** | The config types | [optional] 
**Required** | Pointer to **bool** | Whether this configuration is required | [optional] 
**DefaultValue** | Pointer to **string** | Default value for this configuration | [optional] 
**Importance** | Pointer to **string** | The importance level for a configuration | [optional] 
**Documentation** | Pointer to **string** | The documentation for the configuration | [optional] 
**Group** | Pointer to **string** | The UI group to which the configuration belongs to | [optional] 
**Width** | Pointer to **string** | The width of a configuration value | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Dependents** | Pointer to **[]string** | Other configurations on which this configuration is dependent | [optional] 
**Order** | Pointer to **int32** | The order of configuration in specified group | [optional] 
**Alias** | Pointer to **string** |  | [optional] 

## Methods

### NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition

`func NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition() *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition`

NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition instantiates a new ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinitionWithDefaults

`func NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinitionWithDefaults() *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition`

NewValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinitionWithDefaults instantiates a new ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetImportance

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetImportance() string`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetImportanceOk() (*string, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportance

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetImportance(v string)`

SetImportance sets Importance field to given value.

### HasImportance

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### GetDocumentation

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetGroup

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetWidth

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetDisplayName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDependents

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDependents() []string`

GetDependents returns the Dependents field if non-nil, zero value otherwise.

### GetDependentsOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetDependentsOk() (*[]string, bool)`

GetDependentsOk returns a tuple with the Dependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependents

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetDependents(v []string)`

SetDependents sets Dependents field to given value.

### HasDependents

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasDependents() bool`

HasDependents returns a boolean if a field has been set.

### GetOrder

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAlias

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ValidateConnectv1ConnectorPlugin200ResponseConfigsInnerDefinition) HasAlias() bool`

HasAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


