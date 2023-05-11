# \TransitGatewayAttachmentsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#CreateNetworkingV1TransitGatewayAttachment) | **Post** /networking/v1/transit-gateway-attachments | Create a Transit Gateway Attachment
[**DeleteNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#DeleteNetworkingV1TransitGatewayAttachment) | **Delete** /networking/v1/transit-gateway-attachments/{id} | Delete a Transit Gateway Attachment
[**GetNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#GetNetworkingV1TransitGatewayAttachment) | **Get** /networking/v1/transit-gateway-attachments/{id} | Read a Transit Gateway Attachment
[**ListNetworkingV1TransitGatewayAttachments**](TransitGatewayAttachmentsNetworkingV1Api.md#ListNetworkingV1TransitGatewayAttachments) | **Get** /networking/v1/transit-gateway-attachments | List of Transit Gateway Attachments
[**UpdateNetworkingV1TransitGatewayAttachment**](TransitGatewayAttachmentsNetworkingV1Api.md#UpdateNetworkingV1TransitGatewayAttachment) | **Patch** /networking/v1/transit-gateway-attachments/{id} | Update a Transit Gateway Attachment



## CreateNetworkingV1TransitGatewayAttachment

> CreateNetworkingV1TransitGatewayAttachment202Response CreateNetworkingV1TransitGatewayAttachment(ctx).CreateNetworkingV1TransitGatewayAttachmentRequest(createNetworkingV1TransitGatewayAttachmentRequest).Execute()

Create a Transit Gateway Attachment



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
    createNetworkingV1TransitGatewayAttachmentRequest := *openapiclient.NewCreateNetworkingV1TransitGatewayAttachmentRequest(*openapiclient.NewCreateNetworkingV1PeeringRequestAllOf1Spec()) // CreateNetworkingV1TransitGatewayAttachmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment(context.Background()).CreateNetworkingV1TransitGatewayAttachmentRequest(createNetworkingV1TransitGatewayAttachmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1TransitGatewayAttachment`: CreateNetworkingV1TransitGatewayAttachment202Response
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.CreateNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkingV1TransitGatewayAttachmentRequest** | [**CreateNetworkingV1TransitGatewayAttachmentRequest**](CreateNetworkingV1TransitGatewayAttachmentRequest.md) |  | 

### Return type

[**CreateNetworkingV1TransitGatewayAttachment202Response**](CreateNetworkingV1TransitGatewayAttachment202Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1TransitGatewayAttachment

> DeleteNetworkingV1TransitGatewayAttachment(ctx, id).Environment(environment).Execute()

Delete a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransitGatewayAttachmentsNetworkingV1Api.DeleteNetworkingV1TransitGatewayAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.DeleteNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


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


## GetNetworkingV1TransitGatewayAttachment

> GetNetworkingV1TransitGatewayAttachment200Response GetNetworkingV1TransitGatewayAttachment(ctx, id).Environment(environment).Execute()

Read a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1TransitGatewayAttachment`: GetNetworkingV1TransitGatewayAttachment200Response
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.GetNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**GetNetworkingV1TransitGatewayAttachment200Response**](GetNetworkingV1TransitGatewayAttachment200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1TransitGatewayAttachments

> ListNetworkingV1TransitGatewayAttachments200Response ListNetworkingV1TransitGatewayAttachments(ctx).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()

List of Transit Gateway Attachments



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
    resp, r, err := apiClient.TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).StatusPhase(statusPhase).SpecNetwork(specNetwork).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1TransitGatewayAttachments`: ListNetworkingV1TransitGatewayAttachments200Response
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.ListNetworkingV1TransitGatewayAttachments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1TransitGatewayAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **statusPhase** | **[]string** | Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values. | 
 **specNetwork** | **[]string** | Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNetworkingV1TransitGatewayAttachments200Response**](ListNetworkingV1TransitGatewayAttachments200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1TransitGatewayAttachment

> GetNetworkingV1TransitGatewayAttachment200Response UpdateNetworkingV1TransitGatewayAttachment(ctx, id).UpdateNetworkingV1TransitGatewayAttachmentRequest(updateNetworkingV1TransitGatewayAttachmentRequest).Execute()

Update a Transit Gateway Attachment



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
    id := "id_example" // string | The unique identifier for the transit gateway attachment.
    updateNetworkingV1TransitGatewayAttachmentRequest := *openapiclient.NewUpdateNetworkingV1TransitGatewayAttachmentRequest(*openapiclient.NewUpdateCmkV2ClusterRequestAllOfSpec(interface{}({"id":"env-00000"}))) // UpdateNetworkingV1TransitGatewayAttachmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment(context.Background(), id).UpdateNetworkingV1TransitGatewayAttachmentRequest(updateNetworkingV1TransitGatewayAttachmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1TransitGatewayAttachment`: GetNetworkingV1TransitGatewayAttachment200Response
    fmt.Fprintf(os.Stdout, "Response from `TransitGatewayAttachmentsNetworkingV1Api.UpdateNetworkingV1TransitGatewayAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the transit gateway attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1TransitGatewayAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkingV1TransitGatewayAttachmentRequest** | [**UpdateNetworkingV1TransitGatewayAttachmentRequest**](UpdateNetworkingV1TransitGatewayAttachmentRequest.md) |  | 

### Return type

[**GetNetworkingV1TransitGatewayAttachment200Response**](GetNetworkingV1TransitGatewayAttachment200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

