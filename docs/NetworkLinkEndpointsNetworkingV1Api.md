# \NetworkLinkEndpointsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1NetworkLinkEndpoint**](NetworkLinkEndpointsNetworkingV1Api.md#CreateNetworkingV1NetworkLinkEndpoint) | **Post** /networking/v1/network-link-endpoints | Create a Network Link Endpoint
[**DeleteNetworkingV1NetworkLinkEndpoint**](NetworkLinkEndpointsNetworkingV1Api.md#DeleteNetworkingV1NetworkLinkEndpoint) | **Delete** /networking/v1/network-link-endpoints/{id} | Delete a Network Link Endpoint
[**GetNetworkingV1NetworkLinkEndpoint**](NetworkLinkEndpointsNetworkingV1Api.md#GetNetworkingV1NetworkLinkEndpoint) | **Get** /networking/v1/network-link-endpoints/{id} | Read a Network Link Endpoint
[**ListNetworkingV1NetworkLinkEndpoints**](NetworkLinkEndpointsNetworkingV1Api.md#ListNetworkingV1NetworkLinkEndpoints) | **Get** /networking/v1/network-link-endpoints | List of Network Link Endpoints
[**UpdateNetworkingV1NetworkLinkEndpoint**](NetworkLinkEndpointsNetworkingV1Api.md#UpdateNetworkingV1NetworkLinkEndpoint) | **Patch** /networking/v1/network-link-endpoints/{id} | Update a Network Link Endpoint



## CreateNetworkingV1NetworkLinkEndpoint

> CreateNetworkingV1NetworkLinkEndpoint202Response CreateNetworkingV1NetworkLinkEndpoint(ctx).CreateNetworkingV1NetworkLinkEndpointRequest(createNetworkingV1NetworkLinkEndpointRequest).Execute()

Create a Network Link Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createNetworkingV1NetworkLinkEndpointRequest := *openapiclient.NewCreateNetworkingV1NetworkLinkEndpointRequest(*openapiclient.NewCreateNetworkingV1NetworkLinkEndpointRequestAllOf1Spec()) // CreateNetworkingV1NetworkLinkEndpointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkEndpointsNetworkingV1Api.CreateNetworkingV1NetworkLinkEndpoint(context.Background()).CreateNetworkingV1NetworkLinkEndpointRequest(createNetworkingV1NetworkLinkEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkEndpointsNetworkingV1Api.CreateNetworkingV1NetworkLinkEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1NetworkLinkEndpoint`: CreateNetworkingV1NetworkLinkEndpoint202Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkEndpointsNetworkingV1Api.CreateNetworkingV1NetworkLinkEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1NetworkLinkEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkingV1NetworkLinkEndpointRequest** | [**CreateNetworkingV1NetworkLinkEndpointRequest**](CreateNetworkingV1NetworkLinkEndpointRequest.md) |  | 

### Return type

[**CreateNetworkingV1NetworkLinkEndpoint202Response**](CreateNetworkingV1NetworkLinkEndpoint202Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1NetworkLinkEndpoint

> DeleteNetworkingV1NetworkLinkEndpoint(ctx, id).Environment(environment).Execute()

Delete a Network Link Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the network link endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkLinkEndpointsNetworkingV1Api.DeleteNetworkingV1NetworkLinkEndpoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkEndpointsNetworkingV1Api.DeleteNetworkingV1NetworkLinkEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1NetworkLinkEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkingV1NetworkLinkEndpoint

> GetNetworkingV1NetworkLinkEndpoint200Response GetNetworkingV1NetworkLinkEndpoint(ctx, id).Environment(environment).Execute()

Read a Network Link Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the network link endpoint.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkEndpointsNetworkingV1Api.GetNetworkingV1NetworkLinkEndpoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkEndpointsNetworkingV1Api.GetNetworkingV1NetworkLinkEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1NetworkLinkEndpoint`: GetNetworkingV1NetworkLinkEndpoint200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkEndpointsNetworkingV1Api.GetNetworkingV1NetworkLinkEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1NetworkLinkEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**GetNetworkingV1NetworkLinkEndpoint200Response**](GetNetworkingV1NetworkLinkEndpoint200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1NetworkLinkEndpoints

> ListNetworkingV1NetworkLinkEndpoints200Response ListNetworkingV1NetworkLinkEndpoints(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).SpecNetworkLinkService(specNetworkLinkService).PageSize(pageSize).PageToken(pageToken).Execute()

List of Network Link Endpoints



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environment := "env-00000" // string | Filter the results by exact match for environment.
    specDisplayName := []string{"Inner_example"} // []string | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    statusPhase := []string{"Inner_example"} // []string | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. (optional)
    specNetwork := []string{"Inner_example"} // []string | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. (optional)
    specNetworkLinkService := []string{"Inner_example"} // []string | Filter the results by exact match for spec.network_link_service. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkEndpointsNetworkingV1Api.ListNetworkingV1NetworkLinkEndpoints(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).SpecNetworkLinkService(specNetworkLinkService).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkEndpointsNetworkingV1Api.ListNetworkingV1NetworkLinkEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1NetworkLinkEndpoints`: ListNetworkingV1NetworkLinkEndpoints200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkEndpointsNetworkingV1Api.ListNetworkingV1NetworkLinkEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1NetworkLinkEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | **[]string** | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **specNetworkLinkService** | **[]string** | Filter the results by exact match for spec.network_link_service. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNetworkingV1NetworkLinkEndpoints200Response**](ListNetworkingV1NetworkLinkEndpoints200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1NetworkLinkEndpoint

> GetNetworkingV1NetworkLinkEndpoint200Response UpdateNetworkingV1NetworkLinkEndpoint(ctx, id).UpdateNetworkingV1NetworkLinkEndpointRequest(updateNetworkingV1NetworkLinkEndpointRequest).Execute()

Update a Network Link Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | The unique identifier for the network link endpoint.
    updateNetworkingV1NetworkLinkEndpointRequest := *openapiclient.NewUpdateNetworkingV1NetworkLinkEndpointRequest(*openapiclient.NewUpdateCmkV2ClusterRequestAllOfSpec(interface{}({"id":"env-00000"}))) // UpdateNetworkingV1NetworkLinkEndpointRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkEndpointsNetworkingV1Api.UpdateNetworkingV1NetworkLinkEndpoint(context.Background(), id).UpdateNetworkingV1NetworkLinkEndpointRequest(updateNetworkingV1NetworkLinkEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkEndpointsNetworkingV1Api.UpdateNetworkingV1NetworkLinkEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1NetworkLinkEndpoint`: GetNetworkingV1NetworkLinkEndpoint200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkEndpointsNetworkingV1Api.UpdateNetworkingV1NetworkLinkEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1NetworkLinkEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkingV1NetworkLinkEndpointRequest** | [**UpdateNetworkingV1NetworkLinkEndpointRequest**](UpdateNetworkingV1NetworkLinkEndpointRequest.md) |  | 

### Return type

[**GetNetworkingV1NetworkLinkEndpoint200Response**](GetNetworkingV1NetworkLinkEndpoint200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

