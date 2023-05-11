# \APIKeysIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2ApiKey**](APIKeysIamV2Api.md#CreateIamV2ApiKey) | **Post** /iam/v2/api-keys | Create an API Key
[**DeleteIamV2ApiKey**](APIKeysIamV2Api.md#DeleteIamV2ApiKey) | **Delete** /iam/v2/api-keys/{id} | Delete an API Key
[**GetIamV2ApiKey**](APIKeysIamV2Api.md#GetIamV2ApiKey) | **Get** /iam/v2/api-keys/{id} | Read an API Key
[**ListIamV2ApiKeys**](APIKeysIamV2Api.md#ListIamV2ApiKeys) | **Get** /iam/v2/api-keys | List of API Keys
[**UpdateIamV2ApiKey**](APIKeysIamV2Api.md#UpdateIamV2ApiKey) | **Patch** /iam/v2/api-keys/{id} | Update an API Key



## CreateIamV2ApiKey

> CreateIamV2ApiKey202Response CreateIamV2ApiKey(ctx).CreateIamV2ApiKeyRequest(createIamV2ApiKeyRequest).Execute()

Create an API Key



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
    createIamV2ApiKeyRequest := *openapiclient.NewCreateIamV2ApiKeyRequest(*openapiclient.NewCreateIamV2ApiKeyRequestAllOf1Spec()) // CreateIamV2ApiKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysIamV2Api.CreateIamV2ApiKey(context.Background()).CreateIamV2ApiKeyRequest(createIamV2ApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysIamV2Api.CreateIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2ApiKey`: CreateIamV2ApiKey202Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysIamV2Api.CreateIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIamV2ApiKeyRequest** | [**CreateIamV2ApiKeyRequest**](CreateIamV2ApiKeyRequest.md) |  | 

### Return type

[**CreateIamV2ApiKey202Response**](CreateIamV2ApiKey202Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2ApiKey

> DeleteIamV2ApiKey(ctx, id).Execute()

Delete an API Key



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
    id := "id_example" // string | The unique identifier for the API key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.APIKeysIamV2Api.DeleteIamV2ApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysIamV2Api.DeleteIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamV2ApiKey

> GetIamV2ApiKey200Response GetIamV2ApiKey(ctx, id).Execute()

Read an API Key



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
    id := "id_example" // string | The unique identifier for the API key.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysIamV2Api.GetIamV2ApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysIamV2Api.GetIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2ApiKey`: GetIamV2ApiKey200Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysIamV2Api.GetIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIamV2ApiKey200Response**](GetIamV2ApiKey200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2ApiKeys

> ListIamV2ApiKeys200Response ListIamV2ApiKeys(ctx).SpecOwner(specOwner).SpecResource(specResource).PageSize(pageSize).PageToken(pageToken).Execute()

List of API Keys



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
    specOwner := "specOwner_example" // string | Filter the results by exact match for spec.owner. (optional)
    specResource := "specResource_example" // string | Filter the results by exact match for spec.resource. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysIamV2Api.ListIamV2ApiKeys(context.Background()).SpecOwner(specOwner).SpecResource(specResource).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysIamV2Api.ListIamV2ApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2ApiKeys`: ListIamV2ApiKeys200Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysIamV2Api.ListIamV2ApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2ApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specOwner** | **string** | Filter the results by exact match for spec.owner. | 
 **specResource** | **string** | Filter the results by exact match for spec.resource. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListIamV2ApiKeys200Response**](ListIamV2ApiKeys200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2ApiKey

> GetIamV2ApiKey200Response UpdateIamV2ApiKey(ctx, id).IamV2ApiKeyUpdate(iamV2ApiKeyUpdate).Execute()

Update an API Key



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
    id := "id_example" // string | The unique identifier for the API key.
    iamV2ApiKeyUpdate := *openapiclient.NewIamV2ApiKeyUpdate() // IamV2ApiKeyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIKeysIamV2Api.UpdateIamV2ApiKey(context.Background(), id).IamV2ApiKeyUpdate(iamV2ApiKeyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysIamV2Api.UpdateIamV2ApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2ApiKey`: GetIamV2ApiKey200Response
    fmt.Fprintf(os.Stdout, "Response from `APIKeysIamV2Api.UpdateIamV2ApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2ApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2ApiKeyUpdate** | [**IamV2ApiKeyUpdate**](IamV2ApiKeyUpdate.md) |  | 

### Return type

[**GetIamV2ApiKey200Response**](GetIamV2ApiKey200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

