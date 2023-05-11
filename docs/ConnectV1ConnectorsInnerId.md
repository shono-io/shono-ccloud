# ConnectV1ConnectorsInnerId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | The name of the connector the task belongs to. | [optional] 
**Task** | Pointer to **int32** | Task ID within the connector. | [optional] 

## Methods

### NewConnectV1ConnectorsInnerId

`func NewConnectV1ConnectorsInnerId() *ConnectV1ConnectorsInnerId`

NewConnectV1ConnectorsInnerId instantiates a new ConnectV1ConnectorsInnerId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorsInnerIdWithDefaults

`func NewConnectV1ConnectorsInnerIdWithDefaults() *ConnectV1ConnectorsInnerId`

NewConnectV1ConnectorsInnerIdWithDefaults instantiates a new ConnectV1ConnectorsInnerId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ConnectV1ConnectorsInnerId) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ConnectV1ConnectorsInnerId) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ConnectV1ConnectorsInnerId) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ConnectV1ConnectorsInnerId) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetTask

`func (o *ConnectV1ConnectorsInnerId) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ConnectV1ConnectorsInnerId) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ConnectV1ConnectorsInnerId) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *ConnectV1ConnectorsInnerId) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


