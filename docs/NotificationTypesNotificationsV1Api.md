# \NotificationTypesNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsV1NotificationType**](NotificationTypesNotificationsV1Api.md#GetNotificationsV1NotificationType) | **Get** /notifications/v1/notification-types/{id} | Read a Notification Type
[**ListNotificationsV1NotificationTypes**](NotificationTypesNotificationsV1Api.md#ListNotificationsV1NotificationTypes) | **Get** /notifications/v1/notification-types | List of Notification Types



## GetNotificationsV1NotificationType

> GetNotificationsV1NotificationType200Response GetNotificationsV1NotificationType(ctx, id).Execute()

Read a Notification Type



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
    id := "id_example" // string | The unique identifier for the notification type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1NotificationType`: GetNotificationsV1NotificationType200Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationTypesNotificationsV1Api.GetNotificationsV1NotificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the notification type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1NotificationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationsV1NotificationType200Response**](GetNotificationsV1NotificationType200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1NotificationTypes

> ListNotificationsV1NotificationTypes200Response ListNotificationsV1NotificationTypes(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Notification Types



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1NotificationTypes`: ListNotificationsV1NotificationTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `NotificationTypesNotificationsV1Api.ListNotificationsV1NotificationTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1NotificationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNotificationsV1NotificationTypes200Response**](ListNotificationsV1NotificationTypes200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

