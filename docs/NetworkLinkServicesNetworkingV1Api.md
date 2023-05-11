# \NetworkLinkServicesNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1NetworkLinkService**](NetworkLinkServicesNetworkingV1Api.md#CreateNetworkingV1NetworkLinkService) | **Post** /networking/v1/network-link-services | Create a Network Link Service
[**DeleteNetworkingV1NetworkLinkService**](NetworkLinkServicesNetworkingV1Api.md#DeleteNetworkingV1NetworkLinkService) | **Delete** /networking/v1/network-link-services/{id} | Delete a Network Link Service
[**GetNetworkingV1NetworkLinkService**](NetworkLinkServicesNetworkingV1Api.md#GetNetworkingV1NetworkLinkService) | **Get** /networking/v1/network-link-services/{id} | Read a Network Link Service
[**ListNetworkingV1NetworkLinkServices**](NetworkLinkServicesNetworkingV1Api.md#ListNetworkingV1NetworkLinkServices) | **Get** /networking/v1/network-link-services | List of Network Link Services
[**UpdateNetworkingV1NetworkLinkService**](NetworkLinkServicesNetworkingV1Api.md#UpdateNetworkingV1NetworkLinkService) | **Patch** /networking/v1/network-link-services/{id} | Update a Network Link Service



## CreateNetworkingV1NetworkLinkService

> CreateNetworkingV1NetworkLinkService202Response CreateNetworkingV1NetworkLinkService(ctx).CreateNetworkingV1NetworkLinkServiceRequest(createNetworkingV1NetworkLinkServiceRequest).Execute()

Create a Network Link Service



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
    createNetworkingV1NetworkLinkServiceRequest := *openapiclient.NewCreateNetworkingV1NetworkLinkServiceRequest(*openapiclient.NewCreateNetworkingV1PeeringRequestAllOf1Spec()) // CreateNetworkingV1NetworkLinkServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkServicesNetworkingV1Api.CreateNetworkingV1NetworkLinkService(context.Background()).CreateNetworkingV1NetworkLinkServiceRequest(createNetworkingV1NetworkLinkServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServicesNetworkingV1Api.CreateNetworkingV1NetworkLinkService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1NetworkLinkService`: CreateNetworkingV1NetworkLinkService202Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServicesNetworkingV1Api.CreateNetworkingV1NetworkLinkService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1NetworkLinkServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkingV1NetworkLinkServiceRequest** | [**CreateNetworkingV1NetworkLinkServiceRequest**](CreateNetworkingV1NetworkLinkServiceRequest.md) |  | 

### Return type

[**CreateNetworkingV1NetworkLinkService202Response**](CreateNetworkingV1NetworkLinkService202Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1NetworkLinkService

> DeleteNetworkingV1NetworkLinkService(ctx, id).Environment(environment).Execute()

Delete a Network Link Service



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
    id := "id_example" // string | The unique identifier for the network link service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkLinkServicesNetworkingV1Api.DeleteNetworkingV1NetworkLinkService(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServicesNetworkingV1Api.DeleteNetworkingV1NetworkLinkService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1NetworkLinkServiceRequest struct via the builder pattern


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


## GetNetworkingV1NetworkLinkService

> GetNetworkingV1NetworkLinkService200Response GetNetworkingV1NetworkLinkService(ctx, id).Environment(environment).Execute()

Read a Network Link Service



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
    id := "id_example" // string | The unique identifier for the network link service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkServicesNetworkingV1Api.GetNetworkingV1NetworkLinkService(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServicesNetworkingV1Api.GetNetworkingV1NetworkLinkService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1NetworkLinkService`: GetNetworkingV1NetworkLinkService200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServicesNetworkingV1Api.GetNetworkingV1NetworkLinkService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1NetworkLinkServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**GetNetworkingV1NetworkLinkService200Response**](GetNetworkingV1NetworkLinkService200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1NetworkLinkServices

> ListNetworkingV1NetworkLinkServices200Response ListNetworkingV1NetworkLinkServices(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Network Link Services



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkServicesNetworkingV1Api.ListNetworkingV1NetworkLinkServices(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServicesNetworkingV1Api.ListNetworkingV1NetworkLinkServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1NetworkLinkServices`: ListNetworkingV1NetworkLinkServices200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServicesNetworkingV1Api.ListNetworkingV1NetworkLinkServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1NetworkLinkServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | **[]string** | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNetworkingV1NetworkLinkServices200Response**](ListNetworkingV1NetworkLinkServices200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1NetworkLinkService

> GetNetworkingV1NetworkLinkService200Response UpdateNetworkingV1NetworkLinkService(ctx, id).UpdateNetworkingV1NetworkLinkServiceRequest(updateNetworkingV1NetworkLinkServiceRequest).Execute()

Update a Network Link Service



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
    id := "id_example" // string | The unique identifier for the network link service.
    updateNetworkingV1NetworkLinkServiceRequest := *openapiclient.NewUpdateNetworkingV1NetworkLinkServiceRequest(*openapiclient.NewUpdateCmkV2ClusterRequestAllOfSpec(interface{}({"id":"env-00000"}))) // UpdateNetworkingV1NetworkLinkServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLinkServicesNetworkingV1Api.UpdateNetworkingV1NetworkLinkService(context.Background(), id).UpdateNetworkingV1NetworkLinkServiceRequest(updateNetworkingV1NetworkLinkServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLinkServicesNetworkingV1Api.UpdateNetworkingV1NetworkLinkService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1NetworkLinkService`: GetNetworkingV1NetworkLinkService200Response
    fmt.Fprintf(os.Stdout, "Response from `NetworkLinkServicesNetworkingV1Api.UpdateNetworkingV1NetworkLinkService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the network link service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1NetworkLinkServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkingV1NetworkLinkServiceRequest** | [**UpdateNetworkingV1NetworkLinkServiceRequest**](UpdateNetworkingV1NetworkLinkServiceRequest.md) |  | 

### Return type

[**GetNetworkingV1NetworkLinkService200Response**](GetNetworkingV1NetworkLinkService200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

