# \RoleBindingsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2RoleBinding**](RoleBindingsIamV2Api.md#CreateIamV2RoleBinding) | **Post** /iam/v2/role-bindings | Create a Role Binding
[**DeleteIamV2RoleBinding**](RoleBindingsIamV2Api.md#DeleteIamV2RoleBinding) | **Delete** /iam/v2/role-bindings/{id} | Delete a Role Binding
[**GetIamV2RoleBinding**](RoleBindingsIamV2Api.md#GetIamV2RoleBinding) | **Get** /iam/v2/role-bindings/{id} | Read a Role Binding
[**ListIamV2RoleBindings**](RoleBindingsIamV2Api.md#ListIamV2RoleBindings) | **Get** /iam/v2/role-bindings | List of Role Bindings



## CreateIamV2RoleBinding

> CreateIamV2RoleBindingRequest CreateIamV2RoleBinding(ctx).CreateIamV2RoleBindingRequest(createIamV2RoleBindingRequest).Execute()

Create a Role Binding



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
    createIamV2RoleBindingRequest := *openapiclient.NewCreateIamV2RoleBindingRequest("User:u-111aaa", "CloudClusterAdmin", "crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-aaa1111/cloud-cluster=lkc-1111aaa") // CreateIamV2RoleBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsIamV2Api.CreateIamV2RoleBinding(context.Background()).CreateIamV2RoleBindingRequest(createIamV2RoleBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.CreateIamV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2RoleBinding`: CreateIamV2RoleBindingRequest
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.CreateIamV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIamV2RoleBindingRequest** | [**CreateIamV2RoleBindingRequest**](CreateIamV2RoleBindingRequest.md) |  | 

### Return type

[**CreateIamV2RoleBindingRequest**](CreateIamV2RoleBindingRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2RoleBinding

> GetIamV2RoleBinding200Response DeleteIamV2RoleBinding(ctx, id).Execute()

Delete a Role Binding



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
    id := "id_example" // string | The unique identifier for the role binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsIamV2Api.DeleteIamV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.DeleteIamV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIamV2RoleBinding`: GetIamV2RoleBinding200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.DeleteIamV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the role binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIamV2RoleBinding200Response**](GetIamV2RoleBinding200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamV2RoleBinding

> GetIamV2RoleBinding200Response GetIamV2RoleBinding(ctx, id).Execute()

Read a Role Binding



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
    id := "id_example" // string | The unique identifier for the role binding.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsIamV2Api.GetIamV2RoleBinding(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.GetIamV2RoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2RoleBinding`: GetIamV2RoleBinding200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.GetIamV2RoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the role binding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2RoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIamV2RoleBinding200Response**](GetIamV2RoleBinding200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2RoleBindings

> ListIamV2RoleBindings200Response ListIamV2RoleBindings(ctx).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()

List of Role Bindings



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
    crnPattern := "crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-aaa1111/cloud-cluster=lkc-1111aaa" // string | Filter the results by a partial search of crn_pattern.
    principal := "User:u-111aaa" // string | Filter the results by exact match for principal. (optional)
    roleName := "CloudClusterAdmin" // string | Filter the results by exact match for role_name. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsIamV2Api.ListIamV2RoleBindings(context.Background()).CrnPattern(crnPattern).Principal(principal).RoleName(roleName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsIamV2Api.ListIamV2RoleBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2RoleBindings`: ListIamV2RoleBindings200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsIamV2Api.ListIamV2RoleBindings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2RoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crnPattern** | **string** | Filter the results by a partial search of crn_pattern. | 
 **principal** | **string** | Filter the results by exact match for principal. | 
 **roleName** | **string** | Filter the results by exact match for role_name. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListIamV2RoleBindings200Response**](ListIamV2RoleBindings200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

