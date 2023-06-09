# \EntityV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTags**](EntityV1Api.md#CreateTags) | **Post** /catalog/v1/entity/tags | Bulk Create Tags
[**DeleteTag**](EntityV1Api.md#DeleteTag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a Tag for an Entity
[**GetByUniqueAttributes**](EntityV1Api.md#GetByUniqueAttributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Read an Entity
[**GetTags**](EntityV1Api.md#GetTags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Read Tags for an Entity
[**UpdateTags**](EntityV1Api.md#UpdateTags) | **Put** /catalog/v1/entity/tags | Bulk Update Tags



## CreateTags

> []TagResponse CreateTags(ctx).Tag(tag).Execute()

Bulk Create Tags



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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntityV1Api.CreateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.CreateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.CreateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, typeName, qualifiedName, tagName).Execute()

Delete a Tag for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    tagName := "tagName_example" // string | The name of the tag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EntityV1Api.DeleteTag(context.Background(), typeName, qualifiedName, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 
**tagName** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByUniqueAttributes

> EntityWithExtInfo GetByUniqueAttributes(ctx, typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).Execute()

Read an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    minExtInfo := true // bool | Whether to populate on header and schema attributes (optional) (default to false)
    ignoreRelationships := true // bool | Whether to ignore relationships (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntityV1Api.GetByUniqueAttributes(context.Background(), typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.GetByUniqueAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByUniqueAttributes`: EntityWithExtInfo
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.GetByUniqueAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByUniqueAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minExtInfo** | **bool** | Whether to populate on header and schema attributes | [default to false]
 **ignoreRelationships** | **bool** | Whether to ignore relationships | [default to false]

### Return type

[**EntityWithExtInfo**](EntityWithExtInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []TagResponse GetTags(ctx, typeName, qualifiedName).Execute()

Read Tags for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntityV1Api.GetTags(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTags

> []TagResponse UpdateTags(ctx).Tag(tag).Execute()

Bulk Update Tags



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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntityV1Api.UpdateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.UpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.UpdateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

