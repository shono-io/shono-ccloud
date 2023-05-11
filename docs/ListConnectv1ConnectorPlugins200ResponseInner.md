# ListConnectv1ConnectorPlugins200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | **string** | The connector class name. E.g. BigQuerySink. | 
**Type** | **string** | Type of connector, sink or source. | 
**Version** | Pointer to **string** | The version string for the connector available. | [optional] 

## Methods

### NewListConnectv1ConnectorPlugins200ResponseInner

`func NewListConnectv1ConnectorPlugins200ResponseInner(class string, type_ string, ) *ListConnectv1ConnectorPlugins200ResponseInner`

NewListConnectv1ConnectorPlugins200ResponseInner instantiates a new ListConnectv1ConnectorPlugins200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectv1ConnectorPlugins200ResponseInnerWithDefaults

`func NewListConnectv1ConnectorPlugins200ResponseInnerWithDefaults() *ListConnectv1ConnectorPlugins200ResponseInner`

NewListConnectv1ConnectorPlugins200ResponseInnerWithDefaults instantiates a new ListConnectv1ConnectorPlugins200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) SetClass(v string)`

SetClass sets Class field to given value.


### GetType

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListConnectv1ConnectorPlugins200ResponseInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


