# \CollectionsApi

All URIs are relative to *https://api.use1a1.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionsApi.md#CreateCollection) | **Post** /v1/orgs/self/ws/{workspace}/collections | Create Collection
[**DeleteCollection**](CollectionsApi.md#DeleteCollection) | **Delete** /v1/orgs/self/ws/{workspace}/collections/{collection} | Delete Collection
[**GetCollection**](CollectionsApi.md#GetCollection) | **Get** /v1/orgs/self/ws/{workspace}/collections/{collection} | Retrieve Collection
[**ListCollections**](CollectionsApi.md#ListCollections) | **Get** /v1/orgs/self/collections | List Collections
[**UpdateCollection**](CollectionsApi.md#UpdateCollection) | **Put** /v1/orgs/self/ws/{workspace}/collections/{collection} | Update Collection
[**WorkspaceCollections**](CollectionsApi.md#WorkspaceCollections) | **Get** /v1/orgs/self/ws/{workspace}/collections | List Collections in Workspace



## CreateCollection

> CreateCollectionResponse CreateCollection(ctx, workspace).Body(body).Execute()

Create Collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    body := *openapiclient.NewCreateCollectionRequest() // CreateCollectionRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CreateCollection(context.Background(), workspace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCollection`: CreateCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CreateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateCollectionRequest**](CreateCollectionRequest.md) | JSON object | 

### Return type

[**CreateCollectionResponse**](CreateCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollectionResponse DeleteCollection(ctx, workspace, collection).Execute()

Delete Collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.DeleteCollection(context.Background(), workspace, collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.DeleteCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCollection`: DeleteCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.DeleteCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteCollectionResponse**](DeleteCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> GetCollectionResponse GetCollection(ctx, workspace, collection).Execute()

Retrieve Collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.GetCollection(context.Background(), workspace, collection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.GetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollection`: GetCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCollectionResponse**](GetCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> ListCollectionsResponse ListCollections(ctx).Execute()

List Collections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.ListCollections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollections`: ListCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


### Return type

[**ListCollectionsResponse**](ListCollectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> GetCollectionResponse UpdateCollection(ctx, workspace, collection).Body(body).Execute()

Update Collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")
    collection := "collection_example" // string | name of the collection
    body := *openapiclient.NewUpdateCollectionRequest() // UpdateCollectionRequest | JSON object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.UpdateCollection(context.Background(), workspace, collection).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.UpdateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCollection`: GetCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]
**collection** | **string** | name of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCollectionRequest**](UpdateCollectionRequest.md) | JSON object | 

### Return type

[**GetCollectionResponse**](GetCollectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspaceCollections

> ListCollectionsResponse WorkspaceCollections(ctx, workspace).Execute()

List Collections in Workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.WorkspaceCollections(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.WorkspaceCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkspaceCollections`: ListCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.WorkspaceCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCollectionsResponse**](ListCollectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

