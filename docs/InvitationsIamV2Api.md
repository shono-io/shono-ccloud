# \InvitationsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2Invitation**](InvitationsIamV2Api.md#CreateIamV2Invitation) | **Post** /iam/v2/invitations | Create an Invitation
[**DeleteIamV2Invitation**](InvitationsIamV2Api.md#DeleteIamV2Invitation) | **Delete** /iam/v2/invitations/{id} | Delete an Invitation
[**GetIamV2Invitation**](InvitationsIamV2Api.md#GetIamV2Invitation) | **Get** /iam/v2/invitations/{id} | Read an Invitation
[**ListIamV2Invitations**](InvitationsIamV2Api.md#ListIamV2Invitations) | **Get** /iam/v2/invitations | List of Invitations



## CreateIamV2Invitation

> CreateIamV2InvitationRequest CreateIamV2Invitation(ctx).CreateIamV2InvitationRequest(createIamV2InvitationRequest).Execute()

Create an Invitation



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
    createIamV2InvitationRequest := *openapiclient.NewCreateIamV2InvitationRequest("johndoe@confluent.io") // CreateIamV2InvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsIamV2Api.CreateIamV2Invitation(context.Background()).CreateIamV2InvitationRequest(createIamV2InvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.CreateIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2Invitation`: CreateIamV2InvitationRequest
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.CreateIamV2Invitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2InvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIamV2InvitationRequest** | [**CreateIamV2InvitationRequest**](CreateIamV2InvitationRequest.md) |  | 

### Return type

[**CreateIamV2InvitationRequest**](CreateIamV2InvitationRequest.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2Invitation

> DeleteIamV2Invitation(ctx, id).Execute()

Delete an Invitation



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
    id := "id_example" // string | The unique identifier for the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InvitationsIamV2Api.DeleteIamV2Invitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.DeleteIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2InvitationRequest struct via the builder pattern


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


## GetIamV2Invitation

> GetIamV2Invitation200Response GetIamV2Invitation(ctx, id).Execute()

Read an Invitation



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
    id := "id_example" // string | The unique identifier for the invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsIamV2Api.GetIamV2Invitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.GetIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2Invitation`: GetIamV2Invitation200Response
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.GetIamV2Invitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2InvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIamV2Invitation200Response**](GetIamV2Invitation200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2Invitations

> ListIamV2Invitations200Response ListIamV2Invitations(ctx).Email(email).Status(status).User(user).Creator(creator).PageSize(pageSize).PageToken(pageToken).Execute()

List of Invitations



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
    email := "johndoe@confluent.io" // string | Filter the results by exact match for email. (optional)
    status := "INVITE_STATUS_SENT" // string | Filter the results by exact match for status. (optional)
    user := "u-j93dy8" // string | Filter the results by exact match for user. (optional)
    creator := "u-m2r9o7" // string | Filter the results by exact match for creator. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InvitationsIamV2Api.ListIamV2Invitations(context.Background()).Email(email).Status(status).User(user).Creator(creator).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.ListIamV2Invitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2Invitations`: ListIamV2Invitations200Response
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.ListIamV2Invitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2InvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter the results by exact match for email. | 
 **status** | **string** | Filter the results by exact match for status. | 
 **user** | **string** | Filter the results by exact match for user. | 
 **creator** | **string** | Filter the results by exact match for creator. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListIamV2Invitations200Response**](ListIamV2Invitations200Response.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

