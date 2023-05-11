# \ServiceAccountsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2ServiceAccount**](ServiceAccountsIamV2Api.md#CreateIamV2ServiceAccount) | **Post** /iam/v2/service-accounts | Create a Service Account
[**DeleteIamV2ServiceAccount**](ServiceAccountsIamV2Api.md#DeleteIamV2ServiceAccount) | **Delete** /iam/v2/service-accounts/{id} | Delete a Service Account
[**GetIamV2ServiceAccount**](ServiceAccountsIamV2Api.md#GetIamV2ServiceAccount) | **Get** /iam/v2/service-accounts/{id} | Read a Service Account
[**ListIamV2ServiceAccounts**](ServiceAccountsIamV2Api.md#ListIamV2ServiceAccounts) | **Get** /iam/v2/service-accounts | List of Service Accounts
[**UpdateIamV2ServiceAccount**](ServiceAccountsIamV2Api.md#UpdateIamV2ServiceAccount) | **Patch** /iam/v2/service-accounts/{id} | Update a Service Account



## CreateIamV2ServiceAccount

> CreateIamV2ServiceAccountRequest CreateIamV2ServiceAccount(ctx).CreateIamV2ServiceAccountRequest(createIamV2ServiceAccountRequest).Execute()

Create a Service Account



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
    createIamV2ServiceAccountRequest := *openapiclient.NewCreateIamV2ServiceAccountRequest("DeLorean_auto_repair") // CreateIamV2ServiceAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsIamV2Api.CreateIamV2ServiceAccount(context.Background()).CreateIamV2ServiceAccountRequest(createIamV2ServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsIamV2Api.CreateIamV2ServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2ServiceAccount`: CreateIamV2ServiceAccountRequest
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsIamV2Api.CreateIamV2ServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIamV2ServiceAccountRequest** | [**CreateIamV2ServiceAccountRequest**](CreateIamV2ServiceAccountRequest.md) |  | 

### Return type

[**CreateIamV2ServiceAccountRequest**](CreateIamV2ServiceAccountRequest.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2ServiceAccount

> DeleteIamV2ServiceAccount(ctx, id).Execute()

Delete a Service Account



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
    id := "id_example" // string | The unique identifier for the service account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceAccountsIamV2Api.DeleteIamV2ServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsIamV2Api.DeleteIamV2ServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIamV2ServiceAccount

> GetIamV2ServiceAccount200Response GetIamV2ServiceAccount(ctx, id).Execute()

Read a Service Account



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
    id := "id_example" // string | The unique identifier for the service account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsIamV2Api.GetIamV2ServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsIamV2Api.GetIamV2ServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2ServiceAccount`: GetIamV2ServiceAccount200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsIamV2Api.GetIamV2ServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIamV2ServiceAccount200Response**](GetIamV2ServiceAccount200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2ServiceAccounts

> ListIamV2ServiceAccounts200Response ListIamV2ServiceAccounts(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Service Accounts



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsIamV2Api.ListIamV2ServiceAccounts(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsIamV2Api.ListIamV2ServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2ServiceAccounts`: ListIamV2ServiceAccounts200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsIamV2Api.ListIamV2ServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2ServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListIamV2ServiceAccounts200Response**](ListIamV2ServiceAccounts200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2ServiceAccount

> GetIamV2ServiceAccount200Response UpdateIamV2ServiceAccount(ctx, id).IamV2ServiceAccountUpdate(iamV2ServiceAccountUpdate).Execute()

Update a Service Account



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
    id := "id_example" // string | The unique identifier for the service account.
    iamV2ServiceAccountUpdate := *openapiclient.NewIamV2ServiceAccountUpdate() // IamV2ServiceAccountUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsIamV2Api.UpdateIamV2ServiceAccount(context.Background(), id).IamV2ServiceAccountUpdate(iamV2ServiceAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsIamV2Api.UpdateIamV2ServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2ServiceAccount`: GetIamV2ServiceAccount200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsIamV2Api.UpdateIamV2ServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2ServiceAccountUpdate** | [**IamV2ServiceAccountUpdate**](IamV2ServiceAccountUpdate.md) |  | 

### Return type

[**GetIamV2ServiceAccount200Response**](GetIamV2ServiceAccount200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

