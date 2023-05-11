# \SubscriptionsNotificationsV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsV1Subscription**](SubscriptionsNotificationsV1Api.md#CreateNotificationsV1Subscription) | **Post** /notifications/v1/subscriptions | Create a Subscription
[**DeleteNotificationsV1Subscription**](SubscriptionsNotificationsV1Api.md#DeleteNotificationsV1Subscription) | **Delete** /notifications/v1/subscriptions/{id} | Delete a Subscription
[**GetNotificationsV1Subscription**](SubscriptionsNotificationsV1Api.md#GetNotificationsV1Subscription) | **Get** /notifications/v1/subscriptions/{id} | Read a Subscription
[**ListNotificationsV1Subscriptions**](SubscriptionsNotificationsV1Api.md#ListNotificationsV1Subscriptions) | **Get** /notifications/v1/subscriptions | List of Subscriptions
[**UpdateNotificationsV1Subscription**](SubscriptionsNotificationsV1Api.md#UpdateNotificationsV1Subscription) | **Patch** /notifications/v1/subscriptions/{id} | Update a Subscription



## CreateNotificationsV1Subscription

> CreateNotificationsV1SubscriptionRequest CreateNotificationsV1Subscription(ctx).CreateNotificationsV1SubscriptionRequest(createNotificationsV1SubscriptionRequest).Execute()

Create a Subscription



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
    createNotificationsV1SubscriptionRequest := *openapiclient.NewCreateNotificationsV1SubscriptionRequest(*openapiclient.NewNotificationsV1SubscriptionNotificationType("Id_example", "Related_example", "ResourceName_example"), []openapiclient.GlobalObjectReference{*openapiclient.NewGlobalObjectReference("Id_example", "Related_example", "ResourceName_example")}) // CreateNotificationsV1SubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsNotificationsV1Api.CreateNotificationsV1Subscription(context.Background()).CreateNotificationsV1SubscriptionRequest(createNotificationsV1SubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsNotificationsV1Api.CreateNotificationsV1Subscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsV1Subscription`: CreateNotificationsV1SubscriptionRequest
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsNotificationsV1Api.CreateNotificationsV1Subscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsV1SubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationsV1SubscriptionRequest** | [**CreateNotificationsV1SubscriptionRequest**](CreateNotificationsV1SubscriptionRequest.md) |  | 

### Return type

[**CreateNotificationsV1SubscriptionRequest**](CreateNotificationsV1SubscriptionRequest.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsV1Subscription

> DeleteNotificationsV1Subscription(ctx, id).Execute()

Delete a Subscription



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
    id := "id_example" // string | The unique identifier for the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionsNotificationsV1Api.DeleteNotificationsV1Subscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsNotificationsV1Api.DeleteNotificationsV1Subscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsV1SubscriptionRequest struct via the builder pattern


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


## GetNotificationsV1Subscription

> GetNotificationsV1Subscription200Response GetNotificationsV1Subscription(ctx, id).Execute()

Read a Subscription



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
    id := "id_example" // string | The unique identifier for the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsNotificationsV1Api.GetNotificationsV1Subscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsNotificationsV1Api.GetNotificationsV1Subscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsV1Subscription`: GetNotificationsV1Subscription200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsNotificationsV1Api.GetNotificationsV1Subscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsV1SubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNotificationsV1Subscription200Response**](GetNotificationsV1Subscription200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationsV1Subscriptions

> ListNotificationsV1Subscriptions200Response ListNotificationsV1Subscriptions(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Subscriptions



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
    resp, r, err := apiClient.SubscriptionsNotificationsV1Api.ListNotificationsV1Subscriptions(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsNotificationsV1Api.ListNotificationsV1Subscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsV1Subscriptions`: ListNotificationsV1Subscriptions200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsNotificationsV1Api.ListNotificationsV1Subscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsV1SubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ListNotificationsV1Subscriptions200Response**](ListNotificationsV1Subscriptions200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsV1Subscription

> GetNotificationsV1Subscription200Response UpdateNotificationsV1Subscription(ctx, id).NotificationsV1SubscriptionUpdate(notificationsV1SubscriptionUpdate).Execute()

Update a Subscription



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
    id := "id_example" // string | The unique identifier for the subscription.
    notificationsV1SubscriptionUpdate := *openapiclient.NewNotificationsV1SubscriptionUpdate() // NotificationsV1SubscriptionUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsNotificationsV1Api.UpdateNotificationsV1Subscription(context.Background(), id).NotificationsV1SubscriptionUpdate(notificationsV1SubscriptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsNotificationsV1Api.UpdateNotificationsV1Subscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsV1Subscription`: GetNotificationsV1Subscription200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsNotificationsV1Api.UpdateNotificationsV1Subscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsV1SubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsV1SubscriptionUpdate** | [**NotificationsV1SubscriptionUpdate**](NotificationsV1SubscriptionUpdate.md) |  | 

### Return type

[**GetNotificationsV1Subscription200Response**](GetNotificationsV1Subscription200Response.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

